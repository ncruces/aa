package aa

import (
	"cmp"
	"slices"
)

func MakeSet[K cmp.Ordered](keys ...K) *Tree[K, struct{}] {
	if !increasing(keys) {
		keys = slices.Clone(keys)
		slices.Sort(keys)
		keys = slices.Compact(keys)
	}
	return makeTree[K, struct{}](keys, nil)
}

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
	switch len(keys) {
	case 0:
		return nil
	case 1:
		return &Tree[K, V]{key: keys[0], value: m[keys[0]]}
	}
	key := keys[len(keys)/2]
	left := makeTree(keys[:len(keys)/2], m)
	right := makeTree(keys[len(keys)/2+1:], m)
	return join(left, &Tree[K, V]{key: key, value: m[key]}, right)
}

func increasing[K cmp.Ordered](keys []K) bool {
	for i := len(keys) - 1; i > 0; i-- {
		if !cmp.Less(keys[i-1], keys[i]) {
			return false
		}
	}
	return true
}
