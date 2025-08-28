package aa

import (
	"math/rand"
	"testing"
)

func Test_join(t *testing.T) {
	r := rand.New(rand.NewSource(42))
	var buf [512]byte
	for range 1000 {
		r.Read(buf[:])
		join3Fuzzer(t, buf[:r.Intn(len(buf))])
	}
}

func Fuzz_join3(f *testing.F) {
	f.Fuzz(join3Fuzzer)
}

func join3Fuzzer(t *testing.T, ints []byte) {
	var left, right *Tree[int8, struct{}]

	for _, b := range ints {
		if i := int8(b); i > 0 {
			right.Add(i)
		} else if i < 0 {
			left.Add(i)
		} else {
			break
		}
	}

	var zero Tree[int8, struct{}]
	tree := join(left, &zero, right)
	tree.check()
}

func Test_join2(t *testing.T) {
	r := rand.New(rand.NewSource(42))
	var buf [512]byte
	for range 1000 {
		r.Read(buf[:])
		join2Fuzzer(t, buf[:r.Intn(len(buf))])
	}
}

func Fuzz_join2(f *testing.F) {
	f.Fuzz(join2Fuzzer)
}

func join2Fuzzer(t *testing.T, ints []byte) {
	var left, right *Tree[int8, struct{}]

	for _, b := range ints {
		if i := int8(b); i > 0 {
			right.Add(i)
		} else if i < 0 {
			left.Add(i)
		} else {
			break
		}
	}

	tree := join2(left, right)
	tree.check()
}
