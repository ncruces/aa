package aa

func (tree *Tree[K, V]) ins_rebalance() *Tree[K, V] {
	if tree.need_raise() { // avoid 2 rotations and allocs
		tree.level++
		return tree
	}
	return tree.skew().split()
}

func (tree *Tree[K, V]) del_rebalance() *Tree[K, V] {
	var want int8
	if tree.left != nil && tree.right != nil {
		want = 1 + min(tree.left.level, tree.right.level)
	}
	if tree.level > want {
		tree.level = want
		if tree.right != nil && tree.right.level > want {
			copy := *tree.right
			copy.level = want
			tree.right = &copy
		}
		return tree.skew_rec().split_rec()
	}
	return tree
}

func (tree *Tree[K, V]) need_skew() bool {
	return tree != nil && tree.left != nil &&
		tree.left.level == tree.level
}

func (tree *Tree[K, V]) need_split() bool {
	return tree != nil && tree.right != nil && tree.right.right != nil &&
		tree.right.right.level == tree.level
}

func (tree *Tree[K, V]) need_raise() bool {
	return tree != nil && tree.left != nil && tree.right != nil &&
		tree.left.level == tree.level &&
		tree.right.level == tree.level
}

func (tree *Tree[K, V]) skew() *Tree[K, V] {
	if tree.need_skew() {
		copy := *tree.left
		tree.left = copy.right
		copy.right = tree
		return &copy
	}
	return tree
}

func (tree *Tree[K, V]) skew_rec() *Tree[K, V] {
	if tree.need_skew() {
		copy := *tree.left
		tree.left = copy.right
		copy.right = tree.skew_rec()
		return &copy
	}
	if tree.right.need_skew() {
		node := *tree.right
		tree.right = node.skew_rec()
	}
	return tree
}

func (tree *Tree[K, V]) split() *Tree[K, V] {
	if tree.need_split() {
		copy := *tree.right
		tree.right = copy.left
		copy.left = tree
		copy.level++
		return &copy
	}
	return tree
}

func (tree *Tree[K, V]) split_rec() *Tree[K, V] {
	if tree.need_split() {
		copy := *tree.right
		tree.right = copy.left
		copy.left = tree
		copy.level++
		if copy.right.need_split() {
			node := *copy.right
			copy.right = node.split()
		}
		return &copy
	}
	return tree
}
