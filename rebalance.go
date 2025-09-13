package aa

func (tree *Tree[K, V]) ins_rebalance() *Tree[K, V] {
	if tree.need_raise() { // Avoid 2 rotations and allocs.
		return tree.fixup()
	}
	return tree.skew().split()
}

func (tree *Tree[K, V]) del_rebalance() *Tree[K, V] {
	max := 1 + min(tree.left.Level(), tree.right.Level())
	if tree.Level() > max {
		tree.setLevel(max)
		if tree.right.Level() > max {
			copy := *tree.right
			copy.setLevel(max)
			tree.right = &copy
		}
		return tree.skew_rec().split_rec()
	}
	return tree.fixup()
}

func (tree *Tree[K, V]) need_skew() bool {
	return tree != nil && tree.Level() == tree.left.Level()
}

func (tree *Tree[K, V]) need_split() bool {
	return tree != nil && tree.right != nil &&
		tree.Level() == tree.right.right.Level()
}

func (tree *Tree[K, V]) need_raise() bool {
	return tree != nil &&
		tree.Level() == tree.left.Level() &&
		tree.Level() == tree.right.Level()
}

func (tree *Tree[K, V]) skew() *Tree[K, V] {
	if tree.need_skew() {
		// Rotate right.
		copy := *tree.left
		tree.left = copy.right
		copy.right = tree.fixup()
		tree = &copy
	}
	return tree.fixup()
}

func (tree *Tree[K, V]) skew_rec() *Tree[K, V] {
	if tree.need_skew() {
		// Rotate right.
		copy := *tree.left
		tree.left = copy.right
		copy.right = tree.skew_rec() // Recurse.
		tree = &copy
	}
	if tree.right.need_skew() {
		node := *tree.right
		tree.right = node.skew_rec() // Recurse.
	}
	return tree.fixup()
}

func (tree *Tree[K, V]) split() *Tree[K, V] {
	if tree.need_split() {
		// Rotate left.
		copy := *tree.right
		tree.right = copy.left
		copy.left = tree.fixup()
		tree = &copy
	}
	return tree.fixup()
}

func (tree *Tree[K, V]) split_rec() *Tree[K, V] {
	if tree.need_split() {
		// Rotate left.
		copy := *tree.right
		tree.right = copy.left
		copy.left = tree.fixup()
		tree = &copy
	}
	if tree.right.need_split() {
		node := *tree.right
		tree.right = node.split() // Recurse once.
	}
	return tree.fixup()
}
