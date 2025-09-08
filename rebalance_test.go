package aa

import "testing"

// https://web.archive.org/web/20181104022612/eternallyconfuzzled.com/tuts/datastructures/jsw_tut_andersson.aspx

func TestTree_skew(t *testing.T) {
	//          d,2               b,2
	//         /   \             /   \
	//      b,2     e,1  -->  a,1     d,2
	//     /   \                     /   \
	//  a,1     c,1               c,1     e,1

	in :=
		node("d", 2,
			node("b", 2,
				node("a", 1),
				node("c", 1)),
			node("e", 1))

	tmp := *in
	out := tmp.skew()
	if n := out; n.key != "b" || n.Level() != 2 {
		t.Fatalf("%s,%d", n.key, n.Level())
	}
	if n := out.left; n.key != "a" || n.Level() != 1 {
		t.Fatalf("%s,%d", n.key, n.Level())
	}
	if n := out.right; n.key != "d" || n.Level() != 2 {
		t.Fatalf("%s,%d", n.key, n.Level())
	}
	if n := out.right.left; n.key != "c" || n.Level() != 1 {
		t.Fatalf("%s,%d", n.key, n.Level())
	}
	if n := out.right.right; n.key != "e" || n.Level() != 1 {
		t.Fatalf("%s,%d", n.key, n.Level())
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

	in :=
		node("b", 2,
			node("a", 1),
			node("d", 2,
				node("c", 1),
				node("e", 2)))

	tmp := *in
	out := tmp.split()
	if n := out; n.key != "d" || n.Level() != 3 {
		t.Fatalf("%s,%d", n.key, n.Level())
	}
	if n := out.left; n.key != "b" || n.Level() != 2 {
		t.Fatalf("%s,%d", n.key, n.Level())
	}
	if n := out.left.left; n.key != "a" || n.Level() != 1 {
		t.Fatalf("%s,%d", n.key, n.Level())
	}
	if n := out.left.right; n.key != "c" || n.Level() != 1 {
		t.Fatalf("%s,%d", n.key, n.Level())
	}
	if n := out.right; n.key != "e" || n.Level() != 2 {
		t.Fatalf("%s,%d", n.key, n.Level())
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

	in :=
		node(5, 1,
			node(4, 1),
			node(6, 1))

	tmp := *in
	out := tmp.ins_rebalance()
	if n := out; n.key != 5 || n.Level() != 2 {
		t.Fatalf("%d,%d", n.key, n.Level())
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

	in :=
		node(2, 2,
			nil,
			node(5, 2,
				node(3, 1,
					nil,
					node(4, 1)),
				node(6, 1,
					nil,
					node(7, 1))))

	tmp := *in
	out := tmp.del_rebalance()
	if n := out; n.key != 3 || n.Level() != 2 {
		t.Fatalf("%d,%d", n.key, n.Level())
	}
	if n := out.left; n.key != 2 || n.Level() != 1 {
		t.Fatalf("%d,%d", n.key, n.Level())
	}
	if n := out.right; n.key != 5 || n.Level() != 2 {
		t.Fatalf("%d,%d", n.key, n.Level())
	}
	if n := out.right.left; n.key != 4 || n.Level() != 1 {
		t.Fatalf("%d,%d", n.key, n.Level())
	}
	if n := out.right.right; n.key != 6 || n.Level() != 1 {
		t.Fatalf("%d,%d", n.key, n.Level())
	}
	if n := out.right.right.right; n.key != 7 || n.Level() != 1 {
		t.Fatalf("%d,%d", n.key, n.Level())
	}
}
