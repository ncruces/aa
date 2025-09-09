package aa

import (
	"cmp"
	"slices"
)

// MakeSet builds a tree from a set of keys.
func MakeSet[K cmp.Ordered](keys ...K) *Tree[K, struct{}] {
	if !increasing(keys) {
		keys = slices.Clone(keys)
		slices.Sort(keys)
		keys = slices.Compact(keys)
	}
	return makeTree[K, struct{}](keys, nil)
}

// MakeMap builds a tree from a key-value map.
func MakeMap[K cmp.Ordered, V any](m map[K]V) *Tree[K, V] {
	i, keys := 0, make([]K, len(m))
	for key := range m {
		keys[i] = key
		i++
	}
	slices.Sort(keys)
	return makeTree(keys, m)
}

func makeTree[K cmp.Ordered, V any](keys []K, m map[K]V) *Tree[K, V] {
	if len(keys) == 0 {
		return nil
	}

	// AA trees lean right, so round down.
	mid := (len(keys) - 1) / 2
	left := makeTree(keys[:mid], m)
	right := makeTree(keys[mid+1:], m)

	key := keys[mid]
	return &Tree[K, V]{
		left:  left,
		right: right,
		key:   key,
		value: m[key],
		level: int8(left.Level()), // left.level + 1
	}
}

func increasing[K cmp.Ordered](keys []K) bool {
	for i := len(keys) - 1; i > 0; i-- {
		if !cmp.Less(keys[i-1], keys[i]) {
			return false
		}
	}
	return true
}

// Collect collects key-value pairs from this tree
// into a new map and returns it.
func (tree *Tree[K, V]) Collect() map[K]V {
	m := make(map[K]V, 1<<tree.Level())
	tree.collect(m)
	return m
}

func (tree *Tree[K, V]) collect(m map[K]V) {
	// AA trees lean right, so recurse to the left.
	for tree != nil {
		m[tree.key] = tree.value
		tree.left.collect(m)
		tree = tree.right
	}
}
