package aa

type Tree[K ordered, V any] struct {
	left  *Tree[K, V]
	right *Tree[K, V]
	key   K
	value V
	level int
}

func (tree *Tree[K, V]) Get(key K) (value V, found bool) {
	for tree != nil {
		switch {
		default:
			return tree.value, true
		case less(key, tree.key):
			tree = tree.left
		case less(tree.key, key):
			tree = tree.right
		}
	}
	return // zero, false
}

func (tree *Tree[K, V]) Contains(key K) bool {
	_, found := tree.Get(key)
	return found
}

func (tree *Tree[K, V]) All() func(yield func(K, V) bool) {
	return func(yield func(K, V) bool) { tree.all_r(yield) }
}

func (tree *Tree[K, V]) all_r(yield func(K, V) bool) bool {
	if tree == nil {
		return true
	}
	return tree.left.all_r(yield) &&
		yield(tree.key, tree.value) &&
		tree.right.all_r(yield)
}

func (tree *Tree[K, V]) Put(key K, value V) *Tree[K, V] {
	if tree == nil {
		return &Tree[K, V]{key: key, value: value, level: 1}
	}

	copy := *tree
	switch {
	default:
		copy.value = value
		return &copy
	case less(key, tree.key):
		copy.left = tree.left.Put(key, value)
	case less(tree.key, key):
		copy.right = tree.right.Put(key, value)
	}
	return copy.ins_rebalance()
}

func (tree *Tree[K, V]) Add(key K) *Tree[K, V] {
	if tree == nil {
		return &Tree[K, V]{key: key, level: 1}
	}

	copy := *tree
	switch {
	default:
		return tree
	case less(key, tree.key):
		copy.left = tree.left.Add(key)
		if copy.left == tree.left {
			return tree
		}
	case less(tree.key, key):
		copy.right = tree.right.Add(key)
		if copy.right == tree.right {
			return tree
		}
	}
	return copy.ins_rebalance()
}

func (tree *Tree[K, V]) Delete(key K) *Tree[K, V] {
	if tree == nil {
		return nil
	}

	copy := *tree
	switch {
	case less(key, tree.key):
		copy.left = tree.left.Delete(key)
		if copy.left == tree.left {
			return tree
		}
	case less(tree.key, key):
		copy.right = tree.right.Delete(key)
		if copy.right == tree.right {
			return tree
		}
	default:
		if tree.left == nil {
			return tree.right
		}
		heir := tree.left
		for heir.right != nil {
			heir = heir.right
		}
		copy.key = heir.key
		copy.value = heir.value
		copy.left = tree.left.Delete(heir.key)
	}
	return copy.del_rebalance()
}

func (tree Tree[K, V]) ins_rebalance() *Tree[K, V] {
	if tree.need_raise() {
		tree.level++
		return &tree
	}
	return tree.skew().split()
}

func (tree Tree[K, V]) del_rebalance() *Tree[K, V] {
	var want int
	if tree.left == nil || tree.right == nil {
		want = 1
	} else {
		want = 1 + min(tree.left.level, tree.right.level)
	}
	if tree.level > want {
		tree.level = want
		if tree.right != nil && tree.right.level > want {
			copy := *tree.right
			copy.level = want
			tree.right = &copy
		}
		return tree.skew_r().split_r()
	}
	return &tree
}

func (tree Tree[K, V]) need_skew() bool {
	return tree.left != nil && tree.left.level == tree.level
}

func (tree Tree[K, V]) need_split() bool {
	return tree.right != nil && tree.right.right != nil && tree.right.right.level == tree.level
}

func (tree Tree[K, V]) need_raise() bool {
	return tree.left != nil && tree.right != nil &&
		tree.left.level == tree.level &&
		tree.right.level == tree.level
}

func (tree Tree[K, V]) skew() *Tree[K, V] {
	if tree.need_skew() {
		copy := *tree.left
		tree.left = copy.right
		copy.right = &tree
		return &copy
	}
	return &tree
}

func (tree Tree[K, V]) skew_r() *Tree[K, V] {
	if tree.need_skew() {
		copy := *tree.left
		tree.left = copy.right
		copy.right = tree.skew_r()
		return &copy
	}
	if tree.right != nil && tree.right.need_skew() {
		copy := *tree.right
		tree.right = copy.skew_r()
	}
	return &tree
}

func (tree Tree[K, V]) split() *Tree[K, V] {
	if tree.need_split() {
		copy := *tree.right
		tree.right = copy.left
		copy.left = &tree
		copy.level++
		return &copy
	}
	return &tree
}

func (tree Tree[K, V]) split_r() *Tree[K, V] {
	if tree.need_split() {
		copy := *tree.right
		tree.right = copy.left
		copy.right = copy.right.split_r()
		copy.left = &tree
		copy.level++
		return &copy
	}
	return &tree
}
