package aa

import "cmp"

func isNaN[T cmp.Ordered](x T) bool {
	return x != x
}

func eq[T cmp.Ordered](x, y T) bool {
	if isNaN(x) {
		return isNaN(y)
	}
	return x == y
}
