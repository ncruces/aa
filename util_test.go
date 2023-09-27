package aa

func (tree *Tree[K, V]) check() {
	if tree == nil {
		return
	}

	if tree.left == nil && tree.right == nil && tree.level != 1 {
		panic("The level of every leaf node is one.")
	}
	//
	if (tree.left == nil || tree.right == nil) && tree.level > 1 {
		panic("Every node of level greater than one has two children.")
	}

	tree_left_level := 0
	tree_right_level := 0
	tree_right_right_level := 0
	if tree.left != nil {
		tree_left_level = tree.left.level
	}
	if tree.right != nil {
		tree_right_level = tree.right.level
		if tree.right.right != nil {
			tree_right_right_level = tree.right.right.level
		}
	}

	if !(tree_left_level == tree.level-1) {
		panic("The level of every left child is exactly one less than that of its parent.")
	}
	if !(tree_right_level == tree.level || tree_right_level == tree.level-1) {
		panic("The level of every right child is equal to or one less than that of its parent.")
	}
	if !(tree_right_right_level < tree.level) {
		panic("The level of every right grandchild is strictly less than that of its grandparent.")
	}
}
