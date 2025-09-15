package aa

import (
	"math/rand"
	"strings"
	"testing"
)

// https://web.archive.org/web/20181104022612/eternallyconfuzzled.com/tuts/datastructures/jsw_tut_andersson.aspx

func TestKeyValueLeftRight(t *testing.T) {
	var tt *Tree[int, string]

	if lvl := tt.Level(); lvl != 0 {
		t.Error(lvl)
	}
	if left := tt.Left(); left != nil {
		t.Error(left)
	}
	if right := tt.Right(); right != nil {
		t.Error(right)
	}

	tt = tt.Put(1, "one")

	if key := tt.Key(); key != 1 {
		t.Error(key)
	}
	if value := tt.Value(); value != "one" {
		t.Error(value)
	}
	if lvl := tt.Level(); lvl != 1 {
		t.Error(lvl)
	}
	if left := tt.Left(); left != nil {
		t.Error(left)
	}
	if right := tt.Right(); right != nil {
		t.Error(right)
	}
}

func TestPutGet(t *testing.T) {
	var tt *Tree[int, string]
	tt = tt.Put(1, "one").Put(3, "three").Put(5, "five")
	tt = tt.Put(4, "four").Put(2, "two")
	tt = tt.Put(3, "THREE")
	tt.check()

	if s, ok := tt.Get(0); ok {
		t.Error(s, ok)
	}
	if s, ok := tt.Get(1); !ok || s != "one" {
		t.Error(s, ok)
	}
	if s, ok := tt.Get(3); !ok || s != "THREE" {
		t.Error(s, ok)
	}
	if s, ok := tt.Get(5); !ok || s != "five" {
		t.Error(s, ok)
	}
}

func TestAddHasDelete(t *testing.T) {
	var tt *Tree[int, struct{}]

	tt = tt.Add(1).Add(3).Add(5)
	if ok := tt.Has(3); !ok {
		t.Error()
	}

	tt = tt.Delete(3)
	if ok := tt.Has(3); ok {
		t.Error()
	}

	tt = tt.Delete(5).Delete(1)
	if tt != nil {
		t.Error()
	}
}

func TestAddFloorCeil(t *testing.T) {
	var tt *Tree[int, struct{}]
	tt = tt.Add(1).Add(3).Add(5)
	tt.check()

	if n := tt.Floor(0); n != nil {
		t.Error()
	}
	if n := tt.Floor(3); n.Key() != 3 {
		t.Error()
	}
	if n := tt.Floor(6); n.Key() != 5 {
		t.Error()
	}

	if n := tt.Ceil(0); n.Key() != 1 {
		t.Error()
	}
	if n := tt.Ceil(3); n.Key() != 3 {
		t.Error()
	}
	if n := tt.Ceil(6); n != nil {
		t.Error()
	}
}

func TestAddPatchGet(t *testing.T) {
	var tt *Tree[int, string]

	upper := func(n *Tree[int, string]) (string, bool) {
		if n != nil {
			return strings.ToUpper(n.Value()), true
		}
		return "", false
	}

	tt = tt.Put(1, "one").Put(3, "three").Put(5, "five")
	tt = tt.Patch(0, upper).Patch(3, upper)
	tt.check()

	if s, ok := tt.Get(0); ok {
		t.Error(s, ok)
	}
	if s, ok := tt.Get(3); !ok || s != "THREE" {
		t.Error(s, ok)
	}
}

func TestTree_Add_inc(t *testing.T) {
	var tt *Tree[int, struct{}]
	//             3,3
	//         /         \
	//      1,2           5,2
	//     /   \         /   \
	//  0,1     2,1   4,1     6,1
	tt = tt.Add(0).Add(1).Add(2).Add(3).Add(4).Add(5).Add(6)
	tt.check()

	if n := tt; n.key != 3 || n.Level() != 3 {
		t.Fatalf("%d,%d", n.key, n.Level())
	}
	if n := tt.left; n.key != 1 || n.Level() != 2 {
		t.Fatalf("%d,%d", n.key, n.Level())
	}
	if n := tt.right; n.key != 5 || n.Level() != 2 {
		t.Fatalf("%d,%d", n.key, n.Level())
	}
	if n := tt.left.left; n.key != 0 || n.Level() != 1 {
		t.Fatalf("%d,%d", n.key, n.Level())
	}
	if n := tt.left.right; n.key != 2 || n.Level() != 1 {
		t.Fatalf("%d,%d", n.key, n.Level())
	}
	if n := tt.right.left; n.key != 4 || n.Level() != 1 {
		t.Fatalf("%d,%d", n.key, n.Level())
	}
	if n := tt.right.right; n.key != 6 || n.Level() != 1 {
		t.Fatalf("%d,%d", n.key, n.Level())
	}
}

func TestTree_Add_dec(t *testing.T) {
	var tt *Tree[int, struct{}]
	//      3,2
	//     /   \
	//  2,1     5,2
	//         /   \
	//      4,1     6,1
	tt = tt.Add(6).Add(5).Add(4).Add(3).Add(2)
	tt.check()

	if n := tt; n.key != 3 || n.Level() != 2 {
		t.Fatalf("%d,%d", n.key, n.Level())
	}
	if n := tt.left; n.key != 2 || n.Level() != 1 {
		t.Fatalf("%d,%d", n.key, n.Level())
	}
	if n := tt.right; n.key != 5 || n.Level() != 2 {
		t.Fatalf("%d,%d", n.key, n.Level())
	}
	if n := tt.right.left; n.key != 4 || n.Level() != 1 {
		t.Fatalf("%d,%d", n.key, n.Level())
	}
	if n := tt.right.right; n.key != 6 || n.Level() != 1 {
		t.Fatalf("%d,%d", n.key, n.Level())
	}
}

func TestTree_Add_existing(t *testing.T) {
	var tt *Tree[int, bool]
	tt = tt.Add(1).Add(3).Add(5)

	if a, b := tt, tt.Add(1); a != b {
		t.Fatalf("%p ≠ %p", a, b)
	}
	if a, b := tt, tt.Add(5); a != b {
		t.Fatalf("%p ≠ %p", a, b)
	}
}

func TestTree_Delete(t *testing.T) {
	var tt *Tree[int, struct{}]
	//             3,3
	//         /         \
	//      1,2           5,2
	//     /   \         /   \
	//  0,1     2,1   4,1     6,1
	tt = tt.Add(0).Add(1).Add(2).Add(3).Add(4).Add(5).Add(6)
	//      4,2
	//     /   \
	//  2,1     5,1
	//             \
	//              6,1
	tt = tt.Delete(0).Delete(3).Delete(1)
	tt.check()

	if n := tt; n.key != 4 || n.Level() != 2 {
		t.Fatalf("%d,%d", n.key, n.Level())
	}
	if n := tt.left; n.key != 2 || n.Level() != 1 {
		t.Fatalf("%d,%d", n.key, n.Level())
	}
	if n := tt.right; n.key != 5 || n.Level() != 1 {
		t.Fatalf("%d,%d", n.key, n.Level())
	}
	if n := tt.right.right; n.key != 6 || n.Level() != 1 {
		t.Fatalf("%d,%d", n.key, n.Level())
	}
}

func TestTree_DeleteMin(t *testing.T) {
	var tt *Tree[int, struct{}]
	tt = tt.Add(0).Add(1).Add(2).Add(3).Add(4).Add(5).Add(6)

	for i := range 7 {
		n := tt.Min()
		if n.Key() != i {
			t.Fatalf("%d", n.Key())
		}
		tt, n = tt.DeleteMin()
		if n.Key() != i {
			t.Fatalf("%d", n.Key())
		}
		tt.check()
	}

	if tt != nil {
		t.Error(tt)
	}
	if tt.Min() != nil {
		t.Error(tt)
	}
	tt.DeleteMin()
}

func TestTree_DeleteMax(t *testing.T) {
	var tt *Tree[int, struct{}]
	tt = tt.Add(0).Add(1).Add(2).Add(3).Add(4).Add(5).Add(6)

	for i := 6; i >= 0; i-- {
		n := tt.Max()
		if n.Key() != i {
			t.Fatalf("%d", n.Key())
		}
		tt, n = tt.DeleteMax()
		if n.Key() != i {
			t.Fatalf("%d", n.Key())
		}
		tt.check()
	}

	if tt != nil {
		t.Error(tt)
	}
	if tt.Max() != nil {
		t.Error(tt)
	}
	tt.DeleteMax()
}

func TestTree_Delete_missing(t *testing.T) {
	var tt *Tree[int, struct{}]
	tt = tt.Add(1).Add(3).Add(5)

	if a, b := tt, tt.Delete(0); a != b {
		t.Fatalf("%p ≠ %p", a, b)
	}
	if a, b := tt, tt.Delete(2); a != b {
		t.Fatalf("%p ≠ %p", a, b)
	}
	if a, b := tt, tt.Delete(4); a != b {
		t.Fatalf("%p ≠ %p", a, b)
	}
	if a, b := tt, tt.Delete(6); a != b {
		t.Fatalf("%p ≠ %p", a, b)
	}
}

func TestTree_Delete_pred(t *testing.T) {
	var tt *Tree[int, struct{}]
	tt = tt.Add(1).Add(3).Add(5)

	if a, b := tt, tt.Delete(3, func(node *Tree[int, struct{}]) bool {
		return false
	}); a != b {
		t.Fatalf("%p ≠ %p", a, b)
	}
}

func FuzzTree(f *testing.F) {
	f.Fuzz(func(t *testing.T, cmds []byte) {
		var tt *Tree[byte, byte]

		for i, cmd := range cmds {
			switch i % 3 {
			case 0:
				tt = tt.Add(cmd)
				if !tt.Has(cmd) {
					t.Fail()
				}
				tt.check()

			case 1:
				tt = tt.Delete(cmd)
				if tt.Has(cmd) {
					t.Fail()
				}
				tt.check()

			case 2:
				tt = tt.Put(cmd, cmd)
				if v, ok := tt.Get(cmd); !ok || v != cmd {
					t.Fail()
				}
				tt.check()
			}
		}
	})
}

func TestAddDelete(t *testing.T) {
	var tt *Tree[int, string]

	r := rand.New(rand.NewSource(42))

	for range 1000 {
		n := r.Intn(1000)
		tt = tt.Add(n)
		if !tt.Has(n) {
			t.Fail()
		}
		tt.check()
	}
	for range 1000 {
		n := r.Intn(1000)
		tt = tt.Delete(n)
		if tt.Has(n) {
			t.Fail()
		}
		tt.check()
	}
}

func BenchmarkAddGetDelete(b *testing.B) {
	var tt *Tree[int, string]

	r := rand.New(rand.NewSource(42))
	n := max(16, 16*b.N)

	for range n {
		tt = tt.Add(r.Intn(n))
	}
	for range 8 * n {
		tt.Get(r.Intn(n))
	}
	for range n {
		tt = tt.Delete(r.Intn(n))
	}
}
