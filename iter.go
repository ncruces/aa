//go:build go1.23 || goexperiment.rangefunc

package aa

import "iter"

// All returns an in-order iterator for this tree.
//
//	for node := range tree.All() {
//		fmt.Println(node.Key(), node.Value())
//	})
func (tree *Tree[K, V]) All() iter.Seq[*Tree[K, V]] {
	return func(yield func(*Tree[K, V]) bool) { tree.pull(yield) }
}
