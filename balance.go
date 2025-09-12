package aa

import (
	"math"
	"math/bits"
)

const mask = bits.UintSize - 1

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
