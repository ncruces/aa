//go:build !go1.21

package aa

type ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

func less[T ordered](x, y T) bool {
	return x < y || (x != x && y == y)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
