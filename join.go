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
		return left, node, join(right, tree, tree.right)

	case +1:
		left, node, right = tree.right.Split(key)
		return join(tree.left, tree, left), node, right
	}
}

// Filter returns a tree of nodes for which pred returns true.
func (tree *Tree[K, V]) Filter(pred func(node *Tree[K, V]) bool) *Tree[K, V] {
	if tree == nil {
		return nil
	}

	left := tree.left.Filter(pred)
	right := tree.right.Filter(pred)
	if pred(tree) {
		return join(left, tree, right)
	}
	return join2(left, right)
}

// Partition returns a tree of nodes for which pred returns true,
// and a tree of nodes for which it returns false.
func (tree *Tree[K, V]) Partition(pred func(node *Tree[K, V]) bool) (t, f *Tree[K, V]) {
	if tree == nil {
		return nil, nil
	}

	lt, lf := tree.left.Partition(pred)
	rt, rf := tree.right.Partition(pred)
	if pred(tree) {
		return join(lt, tree, rt), join2(lf, rf)
	}
	return join2(lt, rt), join(lf, tree, rf)
}

func join[K cmp.Ordered, V any](left, node, right *Tree[K, V]) *Tree[K, V] {
	if left == node.left && right == node.right {
		return node
	}

	ll := left.Level()
	rl := right.Level()

	switch {
	case ll < rl:
		left := join(left, node, right.left)
		if left == right.left {
			return right
		}
		copy := *right
		copy.left = left
		return copy.ins_rebalance()

	case ll > rl:
		right := join(left.right, node, right)
		if right == left.right {
			return left
		}
		copy := *left
		copy.right = right
		return copy.ins_rebalance()

	default:
		return &Tree[K, V]{
			left:  left,
			right: right,
			key:   node.key,
			value: node.value,
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
	return join(left, node, right)
}
