package aa

import "cmp"

func (tree *Tree[K, V]) check() int {
	if tree == nil {
		return 0
	}

	// BST invariants.
	if tree.Left() != nil && !cmp.Less(tree.Left().Key(), tree.Key()) {
		panic("the left child's key must be less than this key")
	}
	if tree.Right() != nil && !cmp.Less(tree.Key(), tree.Right().Key()) {
		panic("this key must be less than the right child's key")
	}

	// AA tree invariants.
	if tree.Left() == nil && tree.Right() == nil && tree.Level() != 1 {
		panic("the level of every leaf node is one")
	}
	if (tree.Left() == nil || tree.Right() == nil) && tree.Level() > 1 {
		panic("every node of level greater than one has two children")
	}
	if !(tree.Left().Level() == tree.Level()-1) {
		panic("the level of every left child is exactly one less than that of its parent")
	}
	if !(tree.Right().Level() == tree.Level() || tree.Right().Level() == tree.Level()-1) {
		panic("the level of every right child is equal to or one less than that of its parent")
	}
	if !(tree.Right().Right().Level() < tree.Level()) {
		panic("the level of every right grandchild is strictly less than that of its grandparent")
	}

	// OST invariant.
	len := 1 + tree.Left().check() + tree.Right().check()
	if len != tree.Len() {
		panic("the length of this tree is one plus the length of both children")
	}
	return len
}

func node[K cmp.Ordered](key K, level int, children ...*Tree[K, struct{}]) *Tree[K, struct{}] {
	tree := &Tree[K, struct{}]{key: key}
	if children != nil {
		if len(children) != 2 {
			panic("node must have two children")
		}
		tree.left = children[0]
		tree.right = children[1]
	}
	tree.fixup()
	tree.setLevel(level)
	return tree
}
