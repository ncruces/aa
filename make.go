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
	for k := range m {
		keys[i] = k
		i++
	}
	slices.Sort(keys)
	return makeTree(keys, m)
}

func makeTree[K cmp.Ordered, V any](keys []K, m map[K]V) *Tree[K, V] {
	if len(keys) == 0 {
		return nil
	}

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
