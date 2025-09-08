package aa

import (
	"math/rand"
	"strings"
	"testing"
)

// https://web.archive.org/web/20181104022612/eternallyconfuzzled.com/tuts/datastructures/jsw_tut_andersson.aspx

func TestKeyValueLevelLeftRight(t *testing.T) {
	var aat *Tree[int, string]

	if lvl := aat.Level(); lvl != 0 {
		t.Error(lvl)
	}
	if left := aat.Left(); left != nil {
		t.Error(left)
	}
	if right := aat.Right(); right != nil {
		t.Error(right)
	}

	aat = aat.Put(1, "one")

	if key := aat.Key(); key != 1 {
		t.Error(key)
	}
	if value := aat.Value(); value != "one" {
		t.Error(value)
	}
	if lvl := aat.Level(); lvl != 1 {
		t.Error(lvl)
	}
	if left := aat.Left(); left != nil {
		t.Error(left)
	}
	if right := aat.Right(); right != nil {
		t.Error(right)
	}
}

func TestPutGet(t *testing.T) {
	var aat *Tree[int, string]
	aat = aat.Put(1, "one").Put(3, "three").Put(5, "five")
	aat = aat.Put(4, "four").Put(2, "two")
	aat = aat.Put(3, "THREE")

	if s, ok := aat.Get(0); ok {
		t.Error(s, ok)
	}
	if s, ok := aat.Get(1); !ok || s != "one" {
		t.Error(s, ok)
	}
	if s, ok := aat.Get(3); !ok || s != "THREE" {
		t.Error(s, ok)
	}
	if s, ok := aat.Get(5); !ok || s != "five" {
		t.Error(s, ok)
	}

	if n := aat.Len(); n != 5 {
		t.Error(n)
	}
}

func TestAddHasDelete(t *testing.T) {
	var aat *Tree[int, struct{}]

	aat = aat.Add(1).Add(3).Add(5)
	if ok := aat.Has(3); !ok {
		t.Error()
	}

	aat = aat.Delete(3)
	if ok := aat.Has(3); ok {
		t.Error()
	}

	aat = aat.Delete(5).Delete(1)
	if aat != nil {
		t.Error()
	}
}

func TestAddFloorCeil(t *testing.T) {
	var aat *Tree[int, struct{}]
	aat = aat.Add(1).Add(3).Add(5)

	if n := aat.Floor(0); n != nil {
		t.Error()
	}
	if n := aat.Floor(3); n.Key() != 3 {
		t.Error()
	}
	if n := aat.Floor(6); n.Key() != 5 {
		t.Error()
	}

	if n := aat.Ceil(0); n.Key() != 1 {
		t.Error()
	}
	if n := aat.Ceil(3); n.Key() != 3 {
		t.Error()
	}
	if n := aat.Ceil(6); n != nil {
		t.Error()
	}
}

func TestAddPatchGet(t *testing.T) {
	var aat *Tree[int, string]

	upper := func(n *Tree[int, string]) (string, bool) {
		if n != nil {
			return strings.ToUpper(n.Value()), true
		}
		return "", false
	}

	aat = aat.Put(1, "one").Put(3, "three").Put(5, "five")
	aat = aat.Patch(0, upper).Patch(3, upper)

	if s, ok := aat.Get(0); ok {
		t.Error(s, ok)
	}
	if s, ok := aat.Get(3); !ok || s != "THREE" {
		t.Error(s, ok)
	}
}

func TestTree_Add_inc(t *testing.T) {
	var aat *Tree[int, struct{}]
	//             3,3
	//         /         \
	//      1,2           5,2
	//     /   \         /   \
	//  0,1     2,1   4,1     6,1
	aat = aat.Add(0).Add(1).Add(2).Add(3).Add(4).Add(5).Add(6)

	if n := aat; n.key != 3 || n.Level() != 3 {
		t.Fatalf("%d,%d", n.key, n.Level())
	}
	if n := aat.left; n.key != 1 || n.Level() != 2 {
		t.Fatalf("%d,%d", n.key, n.Level())
	}
	if n := aat.right; n.key != 5 || n.Level() != 2 {
		t.Fatalf("%d,%d", n.key, n.Level())
	}
	if n := aat.left.left; n.key != 0 || n.Level() != 1 {
		t.Fatalf("%d,%d", n.key, n.Level())
	}
	if n := aat.left.right; n.key != 2 || n.Level() != 1 {
		t.Fatalf("%d,%d", n.key, n.Level())
	}
	if n := aat.right.left; n.key != 4 || n.Level() != 1 {
		t.Fatalf("%d,%d", n.key, n.Level())
	}
	if n := aat.right.right; n.key != 6 || n.Level() != 1 {
		t.Fatalf("%d,%d", n.key, n.Level())
	}
}

func TestTree_Add_dec(t *testing.T) {
	var aat *Tree[int, struct{}]
	//      3,2
	//     /   \
	//  2,1     5,2
	//         /   \
	//      4,1     6,1
	aat = aat.Add(6).Add(5).Add(4).Add(3).Add(2)

	if n := aat; n.key != 3 || n.Level() != 2 {
		t.Fatalf("%d,%d", n.key, n.Level())
	}
	if n := aat.left; n.key != 2 || n.Level() != 1 {
		t.Fatalf("%d,%d", n.key, n.Level())
	}
	if n := aat.right; n.key != 5 || n.Level() != 2 {
		t.Fatalf("%d,%d", n.key, n.Level())
	}
	if n := aat.right.left; n.key != 4 || n.Level() != 1 {
		t.Fatalf("%d,%d", n.key, n.Level())
	}
	if n := aat.right.right; n.key != 6 || n.Level() != 1 {
		t.Fatalf("%d,%d", n.key, n.Level())
	}
}

func TestTree_Add_existing(t *testing.T) {
	var aat *Tree[int, bool]
	aat = aat.Add(1).Add(3).Add(5)

	if a, b := aat, aat.Add(1); a != b {
		t.Fatalf("%p ≠ %p", a, b)
	}
	if a, b := aat, aat.Add(5); a != b {
		t.Fatalf("%p ≠ %p", a, b)
	}
}

func TestTree_Delete(t *testing.T) {
	var aat *Tree[int, struct{}]
	//             3,3
	//         /         \
	//      1,2           5,2
	//     /   \         /   \
	//  0,1     2,1   4,1     6,1
	aat = aat.Add(0).Add(1).Add(2).Add(3).Add(4).Add(5).Add(6)
	//      4,2
	//     /   \
	//  2,1     5,1
	//             \
	//              6,1
	aat = aat.Delete(0).Delete(3).Delete(1)

	if n := aat; n.key != 4 || n.Level() != 2 {
		t.Fatalf("%d,%d", n.key, n.Level())
	}
	if n := aat.left; n.key != 2 || n.Level() != 1 {
		t.Fatalf("%d,%d", n.key, n.Level())
	}
	if n := aat.right; n.key != 5 || n.Level() != 1 {
		t.Fatalf("%d,%d", n.key, n.Level())
	}
	if n := aat.right.right; n.key != 6 || n.Level() != 1 {
		t.Fatalf("%d,%d", n.key, n.Level())
	}

	if a, b := aat, aat.Delete(-1); a != b {
		t.Fatalf("%p ≠ %p", a, b)
	}
	if a, b := aat, aat.Delete(10); a != b {
		t.Fatalf("%p ≠ %p", a, b)
	}
}

func TestTree_DeleteMin(t *testing.T) {
	var aat *Tree[int, struct{}]
	aat = aat.Add(0).Add(1).Add(2).Add(3).Add(4).Add(5).Add(6)

	for i := range 7 {
		n := aat.Min()
		if n.Key() != i {
			t.Fatalf("%d", n.Key())
		}
		aat, n = aat.DeleteMin()
		if n.Key() != i {
			t.Fatalf("%d", n.Key())
		}
		aat.check()
	}

	if aat != nil {
		t.Error(aat)
	}
	if aat.Min() != nil {
		t.Error(aat)
	}
	aat.DeleteMin()
}

func TestTree_DeleteMax(t *testing.T) {
	var aat *Tree[int, struct{}]
	aat = aat.Add(0).Add(1).Add(2).Add(3).Add(4).Add(5).Add(6)

	for i := 6; i >= 0; i-- {
		n := aat.Max()
		if n.Key() != i {
			t.Fatalf("%d", n.Key())
		}
		aat, n = aat.DeleteMax()
		if n.Key() != i {
			t.Fatalf("%d", n.Key())
		}
		aat.check()
	}

	if aat != nil {
		t.Error(aat)
	}
	if aat.Max() != nil {
		t.Error(aat)
	}
	aat.DeleteMax()
}

func TestTree_Delete_missing(t *testing.T) {
	var aat *Tree[int, struct{}]
	aat = aat.Add(1).Add(3).Add(5)

	if a, b := aat, aat.Delete(2); a != b {
		t.Fatalf("%p ≠ %p", a, b)
	}
	if a, b := aat, aat.Delete(4); a != b {
		t.Fatalf("%p ≠ %p", a, b)
	}
}

func FuzzTree(f *testing.F) {
	f.Fuzz(func(t *testing.T, cmds []byte) {
		var aat *Tree[byte, byte]

		for i, cmd := range cmds {
			switch i % 3 {
			case 0:
				aat = aat.Add(cmd)
				if !aat.Has(cmd) {
					t.Fail()
				}
				aat.check()

			case 1:
				aat = aat.Delete(cmd)
				if aat.Has(cmd) {
					t.Fail()
				}
				aat.check()

			case 2:
				aat = aat.Put(cmd, cmd)
				if v, ok := aat.Get(cmd); !ok || v != cmd {
					t.Fail()
				}
				aat.check()
			}
		}
	})
}

func BenchmarkAdd(b *testing.B) {
	var aat *Tree[int, string]

	r := rand.New(rand.NewSource(42))
	n := max(16, b.N)

	for range n {
		aat = aat.Add(r.Intn(n))
	}
	for range n {
		aat = aat.Delete(r.Intn(n))
	}
}
