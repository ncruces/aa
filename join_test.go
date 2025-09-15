package aa

import (
	"math/rand"
	"testing"
)

func TestFilter(t *testing.T) {
	var tt *Tree[int, struct{}]
	tt = tt.Add(0).Add(1).Add(2)
	tt = tt.Add(3).Add(4).Add(5)

	even := tt.Filter(func(node *Tree[int, struct{}]) bool {
		return node.key%2 == 0
	})

	var j int
	for i := range even.Ascend() {
		if i != j {
			t.Errorf("%d ≠ %d", i, j)
		}
		j += 2
	}
}

func TestPartition(t *testing.T) {
	var tt *Tree[int, struct{}]
	tt = tt.Add(0).Add(1).Add(2)
	tt = tt.Add(3).Add(4).Add(5)

	even, odd := tt.Partition(func(node *Tree[int, struct{}]) bool {
		return node.key%2 == 0
	})

	var j int
	for i := range even.Ascend() {
		if i != j {
			t.Errorf("%d ≠ %d", i, j)
		}
		j += 2
	}

	j = 1
	for i := range odd.Ascend() {
		if i != j {
			t.Errorf("%d ≠ %d", i, j)
		}
		j += 2
	}
}

func Test_join(t *testing.T) {
	var tt *Tree[int, struct{}]
	tt = tt.Add(0).Add(1).Add(2)

	if a := join(tt.left, tt, tt.right); a != tt {
		t.Errorf("%p ≠ %p", a, tt)
	}

	r := rand.New(rand.NewSource(42))
	var buf [512]byte
	for range 1000 {
		r.Read(buf[:])
		join3Fuzzer(t, buf[:r.Intn(512)])
	}
}

func Fuzz_join3(f *testing.F) {
	f.Fuzz(join3Fuzzer)
}

func join3Fuzzer(t *testing.T, ints []byte) {
	var left, right *Tree[int8, struct{}]

	for _, b := range ints {
		switch i := int8(b); {
		case i > 0:
			right = right.Add(i)
		case i < 0:
			left = left.Add(i)
		}
	}

	var zero Tree[int8, struct{}]
	tree := join(left, &zero, right)
	tree.check()
}

func Test_join2(t *testing.T) {
	var tt *Tree[int, struct{}]
	tt = tt.Add(0).Add(1).Add(2)

	if a := join2(tt, nil); a != tt {
		t.Errorf("%p ≠ %p", a, tt)
	}
	if a := join2(nil, tt); a != tt {
		t.Errorf("%p ≠ %p", a, tt)
	}

	r := rand.New(rand.NewSource(42))
	var buf [512]byte
	for range 1000 {
		r.Read(buf[:])
		join2Fuzzer(t, buf[:r.Intn(512)])
	}
}

func Fuzz_join2(f *testing.F) {
	f.Fuzz(join2Fuzzer)
}

func join2Fuzzer(t *testing.T, ints []byte) {
	var left, right *Tree[int8, struct{}]

	for _, b := range ints {
		switch i := int8(b); {
		case i > 0:
			right = right.Add(i)
		case i < 0:
			left = left.Add(i)
		}
	}

	tree := join2(left, right)
	tree.check()
}
