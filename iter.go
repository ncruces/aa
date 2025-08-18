package aa

import (
	"cmp"
	"iter"
)

// Ascend returns an ascending iterator for this tree.
func (tree *Tree[K, V]) Ascend() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) { tree.ascend(yield) }
}

func (tree *Tree[K, V]) ascend(yield func(K, V) bool) bool {
	if tree == nil {
		return true
	}
	return tree.left.ascend(yield) && yield(tree.key, tree.value) && tree.right.ascend(yield)
}

// Descend returns a descending iterator for this tree.
func (tree *Tree[K, V]) Descend() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) { tree.descend(yield) }
}

func (tree *Tree[K, V]) descend(yield func(K, V) bool) bool {
	if tree == nil {
		return true
	}
	return tree.right.descend(yield) && yield(tree.key, tree.value) && tree.left.descend(yield)
}

// AscendCeil returns an ascending iterator for this tree,
// starting at the least key in this tree greater-than or equal-to pivot.
func (tree *Tree[K, V]) AscendCeil(pivot K) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) { tree.ascendCeil(pivot, yield) }
}

func (tree *Tree[K, V]) ascendCeil(pivot K, yield func(K, V) bool) bool {
	for tree != nil {
		if !cmp.Less(tree.key, pivot) {
			return tree.left.ascendCeil(pivot, yield) && yield(tree.key, tree.value) && tree.right.ascend(yield)
		}
		tree = tree.right
	}
	return true
}

// AscendFloor returns an ascending iterator for this tree,
// starting at the greatest key in this tree less-than or equal-to pivot.
func (tree *Tree[K, V]) AscendFloor(pivot K) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) { tree.ascendFloor(pivot, nil, yield) }
}

func (tree *Tree[K, V]) ascendFloor(pivot K, node *Tree[K, V], yield func(K, V) bool) bool {
	for tree != nil {
		if cmp.Less(pivot, tree.key) {
			return tree.left.ascendFloor(pivot, node, yield) && yield(tree.key, tree.value) && tree.right.ascend(yield)
		}
		node = tree
		tree = tree.right
	}
	if node != nil {
		return yield(node.key, node.value)
	}
	return true
}

// DescendFloor returns a descending iterator for this tree,
// starting at the greatest key in this tree less-than or equal-to pivot.
func (tree *Tree[K, V]) DescendFloor(pivot K) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) { tree.descendFloor(pivot, yield) }
}

func (tree *Tree[K, V]) descendFloor(pivot K, yield func(K, V) bool) bool {
	for tree != nil {
		if !cmp.Less(pivot, tree.key) {
			return tree.right.descendFloor(pivot, yield) && yield(tree.key, tree.value) && tree.left.descend(yield)
		}
		tree = tree.left
	}
	return true
}

// DescendCeil returns a descending iterator for this tree,
// starting at the least key in this tree greater-than or equal-to pivot.
func (tree *Tree[K, V]) DescendCeil(pivot K) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) { tree.descendCeil(pivot, nil, yield) }
}

func (tree *Tree[K, V]) descendCeil(pivot K, node *Tree[K, V], yield func(K, V) bool) bool {
	for tree != nil {
		if cmp.Less(tree.key, pivot) {
			return tree.right.descendCeil(pivot, node, yield) && yield(tree.key, tree.value) && tree.left.descend(yield)
		}
		node = tree
		tree = tree.left
	}
	if node != nil {
		return yield(node.key, node.value)
	}
	return true
}
