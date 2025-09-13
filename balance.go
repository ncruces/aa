package aa

import (
	"math"
	"math/bits"
)

const mask = bits.UintSize - 1

// Balance stores both the level and the cardinality of an AA tree,
// both offset by one: the zero value represents a leaf node of level one.
//
// The lowest 5 or 6 bits store the level,
// the highest 27 or 58 bits store cardinality.
//
// This is sufficient for trees up to 32 or 64 levels,
// with 134217728 or 288230376151711744 elements.
//
// With the smallest possible node size (4 machine words),
// you'd need to exhaust at least half the address space (2GiB or 8EiB)
// before you could hit these limits.
type balance uint

func (b balance) level() int {
	return 1 + int(b&mask)
}

func (b balance) len() int {
	return 1 + int(b>>bits.Len(mask))
}

func (tree *Tree[K, V]) setLevel(level int) {
	tree.balance = balance(level-1) | tree.balance&^mask
}

func (tree *Tree[K, V]) fixup() *Tree[K, V] {
	var sum uint64

	if tree.left != nil {
		sum += 1<<bits.Len(mask) + uint64(tree.left.balance) + 1
	}
	if tree.right != nil {
		sum += 1<<bits.Len(mask) + uint64(tree.right.balance)&^mask
	}
	if sum > math.MaxUint {
		panic("overflow")
	}

	tree.balance = balance(sum)
	return tree
}
