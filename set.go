package aa

import "cmp"

// Union returns the set union of two trees.
// For keys in both trees, the value from t1 is retained.
func Union[K cmp.Ordered, V any](t1, t2 *Tree[K, V]) *Tree[K, V] {
	switch {
	case t1 == t2 || t2 == nil:
		return t1
	case t1 == nil:
		return t2
	}
	left, _, right := t2.Split(t1.key)
	left = Union(t1.left, left)
	right = Union(t1.right, right)
	return join(left, t1, right)
}

// Intersection returns the set intersection of two trees.
// Values are taken from t1.
func Intersection[K cmp.Ordered, V any](t1, t2 *Tree[K, V]) *Tree[K, V] {
	switch {
	case t1 == t2:
		return t1
	case t1 == nil || t2 == nil:
		return nil
	}
	left, node, right := t1.Split(t2.key)
	left = Intersection(left, t2.left)
	right = Intersection(right, t2.right)
	if node == nil {
		return join2(left, right)
	}
	return join(left, node, right)
}

// Difference returns the set difference of two trees.
func Difference[K cmp.Ordered, V any](t1, t2 *Tree[K, V]) *Tree[K, V] {
	switch {
	case t1 == t2 || t1 == nil:
		return nil
	case t2 == nil:
		return t1
	}
	left, _, right := t1.Split(t2.key)
	left = Difference(left, t2.left)
	right = Difference(right, t2.right)
	return join2(left, right)
}

// SymmetricDifference returns the set symmetric difference of two trees.
func SymmetricDifference[K cmp.Ordered, V any](t1, t2 *Tree[K, V]) *Tree[K, V] {
	switch {
	case t1 == t2:
		return nil
	case t1 == nil:
		return t2
	case t2 == nil:
		return t1
	}
	left, node, right := t1.Split(t2.key)
	left = SymmetricDifference(left, t2.left)
	right = SymmetricDifference(right, t2.right)
	if node == nil {
		return join(left, t2, right)
	}
	return join2(left, right)
}
