package aa

func (tree *Tree[K, V]) check() {
	if tree == nil {
		return
	}

	if tree.Left() == nil && tree.Right() == nil && tree.Level() != 1 {
		panic("The level of every leaf node is one.")
	}
	if (tree.Left() == nil || tree.Right() == nil) && tree.Level() > 1 {
		panic("Every node of level greater than one has two children.")
	}
	if !(tree.Left().Level() == tree.Level()-1) {
		panic("The level of every left child is exactly one less than that of its parent.")
	}
	if !(tree.Right().Level() == tree.Level() || tree.Right().Level() == tree.Level()-1) {
		panic("The level of every right child is equal to or one less than that of its parent.")
	}
	if !(tree.Right().Right().Level() < tree.Level()) {
		panic("The level of every right grandchild is strictly less than that of its grandparent.")
	}

	tree.Left().check()
	tree.Right().check()
}
