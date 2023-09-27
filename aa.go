// Package aa implements immutable AA trees.
package aa

// Tree is an immutable AA tree,
// a form of self-balancing binary search tree.
//
// Use nil for an empty *Tree.
// The zero value for Tree is a tree with a single zero key/value pair.
type Tree[K ordered, V any] struct {
	left  *Tree[K, V]
	right *Tree[K, V]
	key   K
	value V
	level int
}

// Key returns the key at the root of this tree.
func (tree *Tree[K, V]) Key() K {
	if tree == nil {
		var key K
		return key
	}
	return tree.key
}

// Value returns the value at the root of this tree.
func (tree *Tree[K, V]) Value() V {
	if tree == nil {
		var value V
		return value
	}
	return tree.value
}

// Left returns the left subtree of this tree,
// containing all keys less than Key().
func (tree *Tree[K, V]) Left() *Tree[K, V] {
	if tree == nil {
		return nil
	}
	return tree.left
}

// Right returns the right subtree of this tree,
// containing all keys greater than Key().
func (tree *Tree[K, V]) Right() *Tree[K, V] {
	if tree == nil {
		return nil
	}
	return tree.right
}

// Level returns the level of this AA tree.
// The number of nodes in a tree of level:
//
//	1 is between 1 and 2
//	2 is between 3 and 8
//	3 is between 7 and 26
//	4 is between 15 and 80
//	n is between 2ⁿ-1 and 3ⁿ-1
func (tree *Tree[K, V]) Level() int {
	if tree == nil {
		return 0
	}
	return tree.level + 1
}

// Get retrieves the value for a given key;
// found indicates whether key exists in this tree.
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

// Contains reports whether key exists in this tree.
func (tree *Tree[K, V]) Contains(key K) bool {
	_, found := tree.Get(key)
	return found
}

// All returns an in-order iterator for this tree.
func (tree *Tree[K, V]) All() func(yield func(*Tree[K, V]) bool) {
	return func(yield func(*Tree[K, V]) bool) { tree.pull(yield) }
}

func (tree *Tree[K, V]) pull(yield func(*Tree[K, V]) bool) bool {
	if tree == nil {
		return true
	}
	return tree.left.pull(yield) && yield(tree) && tree.right.pull(yield)
}

// Put returns a modified tree with key set to value.
//
//	tree.Put(key, value).Get(key) ⟹ (value, true)
func (tree *Tree[K, V]) Put(key K, value V) *Tree[K, V] {
	if tree == nil {
		return &Tree[K, V]{key: key, value: value}
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

// Add returns a modified tree that contains key.
//
//	tree.Add(key).Contains(key) ⟹ true
func (tree *Tree[K, V]) Add(key K) *Tree[K, V] {
	if tree == nil {
		return &Tree[K, V]{key: key}
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

// Delete returns a modified tree with key removed from it.
//
//	tree.Delete(key).Contains(key) ⟹ false
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
