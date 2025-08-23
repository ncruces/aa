package aa

import (
	"math/rand"
	"slices"
	"testing"
)

func TestUnion(t *testing.T) {
	var t1, t2 *Tree[int, string]
	t1 = t1.Put(1, "one").Put(2, "two").Put(3, "three").Put(5, "five")
	t2 = t2.Put(0, "zero").Put(2, "").Put(4, "four")

	if a := Union(t1, nil); a != t1 {
		t.Errorf("%p ≠ %p", a, t1)
	}
	if a := Union(nil, t2); a != t2 {
		t.Errorf("%p ≠ %p", a, t2)
	}

	var j int
	for i, v := range Union(t1, t2).Ascend() {
		if i != j {
			t.Errorf("%d ≠ %d", i, j)
		}
		if v == "" {
			t.Error()
		}
		j++
	}
}

func TestIntersection(t *testing.T) {
	var t1, t2 *Tree[int, string]
	t1 = t1.Put(1, "one").Put(2, "two").Put(3, "three").Put(5, "five")
	t2 = t2.Put(0, "zero").Put(2, "").Put(4, "four")

	var out []int
	for i, v := range Intersection(t1, t2).Ascend() {
		out = append(out, i)
		if v == "" {
			t.Error()
		}
	}
	if !slices.Equal(out, []int{2}) {
		t.Error(out)
	}
}

func TestDifference(t *testing.T) {
	var t1, t2 *Tree[int, string]
	t1 = t1.Put(1, "one").Put(2, "two").Put(3, "three").Put(5, "five")
	t2 = t2.Put(0, "zero").Put(2, "").Put(4, "four")

	var out []int
	for i, v := range Difference(t1, t2).Ascend() {
		out = append(out, i)
		if v == "" {
			t.Error()
		}
	}
	if !slices.Equal(out, []int{1, 3, 5}) {
		t.Error(out)
	}
}

func Test_join(t *testing.T) {
	r := rand.New(rand.NewSource(42))

	for range 1000 {
		var left, right *Tree[int8, struct{}]

		for range r.Intn(200) {
			left = left.Add(-1 - int8(r.Intn(100)))
		}
		for range r.Intn(200) {
			right = right.Add(+1 + int8(r.Intn(100)))
		}

		tree := join(0, struct{}{}, left, right)
		tree.check()
	}
}

func Test_join2(t *testing.T) {
	r := rand.New(rand.NewSource(42))

	for range 1000 {
		var left, right *Tree[int8, struct{}]

		for range r.Intn(200) {
			left = left.Add(-1 - int8(r.Intn(100)))
		}
		for range r.Intn(200) {
			right = right.Add(+1 + int8(r.Intn(100)))
		}

		tree := join2(left, right)
		tree.check()
	}
}
