package aa

import (
	"cmp"
	"iter"
)

// Equal reports whether two trees contain the same key/value pairs.
func Equal[K cmp.Ordered, V comparable](t1, t2 *Tree[K, V]) bool {
	if t1 == t2 {
		return true
	}
	l1 := t1.Len()
	l2 := t2.Len()
	if l1 != l2 {
		return false
	}

	n1, s1 := iter.Pull2(t1.Ascend())
	defer s1()
	n2, s2 := iter.Pull2(t2.Ascend())
	defer s2()

	for range l1 {
		k1, v1, _ := n1()
		k2, v2, _ := n2()
		if (k1 != k2 && (k1 == k1 || k2 == k2)) || v1 != v2 {
			return false
		}
	}
	return true
}

// Subset reports whether t2 contains every t1 key/value pair.
func Subset[K cmp.Ordered, V comparable](t1, t2 *Tree[K, V]) bool {
	if t1 == t2 {
		return true
	}
	l1 := t1.Len()
	if l1 == 0 {
		return true
	}
	l2 := t2.Len()
	if l1 > l2 {
		return false
	}

	n1, s1 := iter.Pull2(t1.Ascend())
	defer s1()
	n2, s2 := iter.Pull2(t2.Ascend())
	defer s2()

loop1:
	for ; l1 <= l2 && l1 > 0; l1-- {
		k1, v1, _ := n1()

		for ; l1 <= l2 && l2 > 0; l2-- {
			k2, v2, _ := n2()

			switch cmp.Compare(k1, k2) {
			case -1:
				return false // won't find k1 in t2
			case +1:
				continue
			}
			if v1 == v2 {
				continue loop1
			}
			return false // found k1, v1 != v2
		}
		return false // couldn't find k1 in t2
	}
	return l1 == 0 // checked all k1
}

// Overlap reports whether t1 contains any keys from t2.
func Overlap[K cmp.Ordered, V any](t1, t2 *Tree[K, V]) bool {
	if t1 == nil || t2 == nil {
		return false
	}
	if t1 == t2 {
		return true
	}

	n1, s1 := iter.Pull2(t1.Ascend())
	defer s1()
	n2, s2 := iter.Pull2(t2.Ascend())
	defer s2()

	k1, _, ok1 := n1()
	k2, _, ok2 := n2()
	for ok1 && ok2 {
		switch cmp.Compare(k1, k2) {
		case -1:
			k1, _, ok1 = n1()
		case +1:
			k2, _, ok2 = n2()
		default:
			return true
		}
	}
	return false
}
