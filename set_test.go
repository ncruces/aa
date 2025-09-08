package aa

import (
	"slices"
	"testing"
)

func TestUnion(t *testing.T) {
	var t1, t2 *Tree[int, string]
	t1 = t1.Put(1, "one").Put(2, "two").Put(3, "three").Put(5, "five")
	t2 = t2.Put(0, "zero").Put(2, "").Put(4, "four")

	if a := Union(t1, t1); a != t1 {
		t.Errorf("%p ≠ %p", a, t1)
	}
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

	if a := Intersection(t1, t1); a != t1 {
		t.Errorf("%p ≠ %p", a, t1)
	}
	if a := Intersection(t1, nil); a != nil {
		t.Error()
	}
	if a := Intersection(nil, t2); a != nil {
		t.Error()
	}

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

	if a := Difference(t1, t1); a != nil {
		t.Error()
	}
	if a := Difference(nil, t2); a != nil {
		t.Error()
	}
	if a := Difference(t1, nil); a != t1 {
		t.Errorf("%p ≠ %p", a, t1)
	}

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

func TestSymmetricDifference(t *testing.T) {
	var t1, t2 *Tree[int, string]
	t1 = t1.Put(1, "one").Put(2, "two").Put(3, "three").Put(5, "five")
	t2 = t2.Put(0, "zero").Put(2, "").Put(4, "four")

	if a := SymmetricDifference(t1, t1); a != nil {
		t.Error()
	}
	if a := SymmetricDifference(nil, t2); a != t2 {
		t.Errorf("%p ≠ %p", a, t2)
	}
	if a := SymmetricDifference(t1, nil); a != t1 {
		t.Errorf("%p ≠ %p", a, t1)
	}

	var out []int
	for i, v := range SymmetricDifference(t1, t2).Ascend() {
		out = append(out, i)
		if v == "" {
			t.Error()
		}
	}
	if !slices.Equal(out, []int{0, 1, 3, 4, 5}) {
		t.Error(out)
	}
}
