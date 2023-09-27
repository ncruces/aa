//go:build go1.21

package aa

import "cmp"

type ordered = cmp.Ordered

func less[T ordered](x, y T) bool {
	return cmp.Less[T](x, y)
}
