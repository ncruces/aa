package aa

import "cmp"

// Select finds the node at index i of this tree and returns it.
func (tree *Tree[K, V]) Select(i int) *Tree[K, V] {
	for tree != nil {
		p := tree.left.Len()
		switch cmp.Compare(i, p) {
		case -1:
			tree = tree.left
		case +1:
			tree = tree.right
			i -= p + 1
		default:
			return tree
		}
	}
	return tree
}

// Rank finds the rank of key,
// the number of nodes in this tree less than key.
func (tree *Tree[K, V]) Rank(key K) int {
	k := 0
	for tree != nil {
		switch cmp.Compare(key, tree.key) {
		case -1:
			tree = tree.left

		case +1:
			k += tree.left.Len() + 1
			tree = tree.right

		default:
			return k + tree.left.Len()
		}
	}
	return k
}
