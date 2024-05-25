package aa

import "cmp"

func (tree *Tree[K, V]) check() {
	if tree == nil {
		return
	}

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

	tree.Left().check()
	tree.Right().check()
}

func node[K cmp.Ordered](key K, level int, children ...*Tree[K, struct{}]) *Tree[K, struct{}] {
	if children == nil {
		return &Tree[K, struct{}]{key: key, level: level - 1}
	}
	if len(children) != 2 {
		panic("node must have two children")
	}
	return &Tree[K, struct{}]{
		key: key, level: level - 1,
		left: children[0], right: children[1]}
}
