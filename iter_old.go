//go:build !(go1.23 || goexperiment.rangefunc)

package aa

// All returns an in-order iterator for this tree.
//
//	tree.All()(func(node *aa.Tree[int, any]) bool {
//		fmt.Println(node.Key(), node.Value())
//		return true
//	})
func (tree *Tree[K, V]) All() func(yield func(node *Tree[K, V]) bool) {
	return func(yield func(*Tree[K, V]) bool) { tree.pull(yield) }
}
