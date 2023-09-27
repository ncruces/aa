package aa

import "testing"

// https://web.archive.org/web/20181104022612/eternallyconfuzzled.com/tuts/datastructures/jsw_tut_andersson.aspx

func TestKeyValueLevelLeftRight(t *testing.T) {
	var aat *Tree[int, string]

	if key := aat.Key(); key != 0 {
		t.Error(key)
	}
	if value := aat.Value(); value != "" {
		t.Error(value)
	}
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

func TestPutGetAll(t *testing.T) {
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

	var slice []int
	aat.Put(0, "zero").All()(func(t *Tree[int, string]) bool {
		slice = append(slice, t.key)
		return true
	})

	for i, j := range slice {
		if i != j {
			t.Errorf("%d ≠ %d", i, j)
		}
	}
}

func TestAddContainsDelete(t *testing.T) {
	var aat *Tree[int, string]

	aat = aat.Add(1).Add(3).Add(5)
	if ok := aat.Contains(3); !ok {
		t.Error()
	}

	aat = aat.Delete(3)
	if ok := aat.Contains(3); ok {
		t.Error()
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

func TestTree_skew(t *testing.T) {
	//          d,2               b,2
	//         /   \             /   \
	//      b,2     e,1  -->  a,1     d,2
	//     /   \                     /   \
	//  a,1     c,1               c,1     e,1

	in := &Tree[string, struct{}]{
		key: "d", level: 2,
		left: &Tree[string, struct{}]{
			key: "b", level: 2,
			left: &Tree[string, struct{}]{
				key: "a", level: 1},
			right: &Tree[string, struct{}]{
				key: "c", level: 1}},
		right: &Tree[string, struct{}]{
			key: "e", level: 1}}

	out := in.skew()
	if n := out; n.key != "b" || n.level != 2 {
		t.Fatalf("%s,%d", n.key, n.level)
	}
	if n := out.left; n.key != "a" || n.level != 1 {
		t.Fatalf("%s,%d", n.key, n.level)
	}
	if n := out.right; n.key != "d" || n.level != 2 {
		t.Fatalf("%s,%d", n.key, n.level)
	}
	if n := out.right.left; n.key != "c" || n.level != 1 {
		t.Fatalf("%s,%d", n.key, n.level)
	}
	if n := out.right.right; n.key != "e" || n.level != 1 {
		t.Fatalf("%s,%d", n.key, n.level)
	}
	if a, b := in.left.left, out.left; a != b {
		t.Fatalf("%p ≠ %p", a, b)
	}
	if a, b := in.left.right, out.right.left; a != b {
		t.Fatalf("%p ≠ %p", a, b)
	}
	if a, b := in.right, out.right.right; a != b {
		t.Fatalf("%p ≠ %p", a, b)
	}
}

func TestTree_split(t *testing.T) {
	//      b,2                     d,3
	//     /   \                   /   \
	//  a,1     d,2     -->     b,2     e,2
	//         /   \           /   \
	//      c,1     e,2     a,1     c,1

	in := &Tree[string, struct{}]{
		key: "b", level: 2,
		left: &Tree[string, struct{}]{
			key: "a", level: 1,
		},
		right: &Tree[string, struct{}]{
			key: "d", level: 2,
			left: &Tree[string, struct{}]{
				key: "c", level: 1},
			right: &Tree[string, struct{}]{
				key: "e", level: 2}}}

	out := in.split()
	if n := out; n.key != "d" || n.level != 3 {
		t.Fatalf("%s,%d", n.key, n.level)
	}
	if n := out.left; n.key != "b" || n.level != 2 {
		t.Fatalf("%s,%d", n.key, n.level)
	}
	if n := out.left.left; n.key != "a" || n.level != 1 {
		t.Fatalf("%s,%d", n.key, n.level)
	}
	if n := out.left.right; n.key != "c" || n.level != 1 {
		t.Fatalf("%s,%d", n.key, n.level)
	}
	if n := out.right; n.key != "e" || n.level != 2 {
		t.Fatalf("%s,%d", n.key, n.level)
	}
	if a, b := in.left, out.left.left; a != b {
		t.Fatalf("%p ≠ %p", a, b)
	}
	if a, b := in.right.left, out.left.right; a != b {
		t.Fatalf("%p ≠ %p", a, b)
	}
	if a, b := in.right.right, out.right; a != b {
		t.Fatalf("%p ≠ %p", a, b)
	}
}

func TestTree_ins_rebalance(t *testing.T) {
	//      5,1        4,1                5,2
	//     /   \    ->    \        ->    /   \
	//  4,1     6,1        5,1        4,1     6,1
	//                        \
	//                         6,1

	in := &Tree[int, struct{}]{
		key: 5, level: 1,
		left: &Tree[int, struct{}]{
			key: 4, level: 1},
		right: &Tree[int, struct{}]{
			key: 6, level: 1}}

	out := in.ins_rebalance()
	if n := out; n.key != 5 || n.level != 2 {
		t.Fatalf("%d,%d", n.key, n.level)
	}
	if a, b := in.left, out.left; a != b {
		t.Fatalf("%p ≠ %p", a, b)
	}
	if a, b := in.right, out.right; a != b {
		t.Fatalf("%p ≠ %p", a, b)
	}
}

func TestTree_del_rebalance(t *testing.T) {
	//  2,2                    3,2
	//     \                  /   \
	//      5,2            2,1     5,2
	//     /   \        ->        /   \
	//  3,1     6,1            4,1     6,1
	//     \       \                      \
	//      4,1     7,1                    7,1

	in := &Tree[int, struct{}]{
		key: 2, level: 2,
		right: &Tree[int, struct{}]{
			key: 5, level: 2,
			left: &Tree[int, struct{}]{
				key: 3, level: 1,
				right: &Tree[int, struct{}]{
					key: 4, level: 1}},
			right: &Tree[int, struct{}]{
				key: 6, level: 1,
				right: &Tree[int, struct{}]{
					key: 7, level: 1}}}}

	out := in.del_rebalance()
	if n := out; n.key != 3 || n.level != 2 {
		t.Fatalf("%d,%d", n.key, n.level)
	}
	if n := out.left; n.key != 2 || n.level != 1 {
		t.Fatalf("%d,%d", n.key, n.level)
	}
	if n := out.right; n.key != 5 || n.level != 2 {
		t.Fatalf("%d,%d", n.key, n.level)
	}
	if n := out.right.left; n.key != 4 || n.level != 1 {
		t.Fatalf("%d,%d", n.key, n.level)
	}
	if n := out.right.right; n.key != 6 || n.level != 1 {
		t.Fatalf("%d,%d", n.key, n.level)
	}
	if n := out.right.right.right; n.key != 7 || n.level != 1 {
		t.Fatalf("%d,%d", n.key, n.level)
	}
}

func FuzzTree(f *testing.F) {
	f.Fuzz(func(t *testing.T, cmds []byte) {
		var aat *Tree[byte, byte]

		for i, cmd := range cmds {
			switch i % 3 {
			case 0:
				aat = aat.Add(cmd)
				if !aat.Contains(cmd) {
					t.Fail()
				}
				aat.check()

			case 1:
				aat = aat.Delete(cmd)
				if aat.Contains(cmd) {
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
