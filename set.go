package aa

import "cmp"

// Split partitions this tree around a key. It returns
// a left tree with keys less than key,
// a right tree with keys greater than key,
// and the node for that key.
func (tree *Tree[K, V]) Split(key K) (left, node, right *Tree[K, V]) {
	if tree == nil {
		return nil, nil, nil
	}

	switch cmp.Compare(key, tree.key) {
	default:
		return tree.left, tree, tree.right

	case -1:
		left, node, right = tree.left.Split(key)
		return left, node, join(tree.key, tree.value, right, tree.right)

	case +1:
		left, node, right = tree.right.Split(key)
		return join(tree.key, tree.value, tree.left, left), node, right
	}
}

// Union returns the set union of two trees.
// For keys in both trees, the value from t1 is retained.
func Union[K cmp.Ordered, V any](t1, t2 *Tree[K, V]) *Tree[K, V] {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	left, _, right := t2.Split(t1.key)
	return join(t1.key, t1.value, Union(t1.left, left), Union(t1.right, right))
}

// Intersection returns the set intersection of two trees.
// Values are taken from t1.
func Intersection[K cmp.Ordered, V any](t1, t2 *Tree[K, V]) *Tree[K, V] {
	if t1 == nil {
		return nil
	}
	if t2 == nil {
		return nil
	}
	left, node, right := t1.Split(t2.key)
	left = Intersection(left, t2.left)
	right = Intersection(right, t2.right)
	if node == nil {
		return join2(left, right)
	}
	return join(node.key, node.value, left, right)
}

// Difference returns the set difference of two trees.
func Difference[K cmp.Ordered, V any](t1, t2 *Tree[K, V]) *Tree[K, V] {
	if t1 == nil {
		return nil
	}
	if t2 == nil {
		return t1
	}
	left, _, right := t1.Split(t2.key)
	left = Difference(left, t2.left)
	right = Difference(right, t2.right)
	return join2(left, right)
}

func join[K cmp.Ordered, V any](key K, value V, left, right *Tree[K, V]) *Tree[K, V] {
	ll := left.Level()
	rl := right.Level()

	switch {
	case ll < rl:
		copy := *right
		copy.left = join(key, value, left, copy.left)
		return copy.ins_rebalance()

	case ll > rl:
		copy := *left
		copy.right = join(key, value, copy.right, right)
		return copy.ins_rebalance()

	default:
		return &Tree[K, V]{
			left:  left,
			right: right,
			key:   key,
			value: value,
			level: int8(ll), // left.level + 1
		}
	}
}

func join2[K cmp.Ordered, V any](left, right *Tree[K, V]) *Tree[K, V] {
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	left, node := left.DeleteMax()
	return join(node.key, node.value, left, right)
}
