package aa

import (
	"testing"
)

func TestMakeSet(t *testing.T) {
	aat := MakeSet(0, 2, 4, 6, 5, 3, 1)

	if n := aat; n.key != 3 || n.Level() != 3 {
		t.Fatalf("%d,%d", n.key, n.Level())
	}
	if n := aat.left; n.key != 1 || n.Level() != 2 {
		t.Fatalf("%d,%d", n.key, n.Level())
	}
	if n := aat.right; n.key != 5 || n.Level() != 2 {
		t.Fatalf("%d,%d", n.key, n.Level())
	}
	if n := aat.left.left; n.key != 0 || n.Level() != 1 {
		t.Fatalf("%d,%d", n.key, n.Level())
	}
	if n := aat.left.right; n.key != 2 || n.Level() != 1 {
		t.Fatalf("%d,%d", n.key, n.Level())
	}
	if n := aat.right.left; n.key != 4 || n.Level() != 1 {
		t.Fatalf("%d,%d", n.key, n.Level())
	}
	if n := aat.right.right; n.key != 6 || n.Level() != 1 {
		t.Fatalf("%d,%d", n.key, n.Level())
	}
}

func BenchmarkMakeSet(b *testing.B) {
	var slice []int
	// r := rand.New(rand.NewSource(42))
	// for range 1000 * b.N {
	// 	slice = append(slice, r.Int())
	// }
	for i := range 1000 * b.N {
		slice = append(slice, i)
	}
	b.ResetTimer()
	MakeSet(slice...).check()
}
