package aa

import (
	"slices"
	"testing"
)

func TestPutGetAscendDescend(t *testing.T) {
	var aat *Tree[int, struct{}]
	aat = aat.Add(1).Add(3).Add(5)
	aat = aat.Add(4).Add(2)
	aat = aat.Add(3)

	var j int
	for i := range aat.Add(0).Ascend() {
		if i != j {
			t.Errorf("%d ≠ %d", i, j)
		}
		j++
	}
	for i := range aat.Add(6).Descend() {
		if i != j {
			t.Errorf("%d ≠ %d", i, j)
		}
		j--
	}
}

func TestAscendCeil(t *testing.T) {
	var aat *Tree[int, struct{}]
	aat = aat.Add(1).Add(9)
	aat = aat.Add(5).Add(3)
	aat = aat.Add(7)

	var out []int

	out = out[:0]
	for i := range aat.AscendCeil(4) {
		out = append(out, i)
	}
	if !slices.Equal(out, []int{5, 7, 9}) {
		t.Error(out)
	}

	out = out[:0]
	for i := range aat.AscendCeil(3) {
		out = append(out, i)
	}
	if !slices.Equal(out, []int{3, 5, 7, 9}) {
		t.Error(out)
	}

	out = out[:0]
	for i := range aat.AscendCeil(0) {
		out = append(out, i)
	}
	if !slices.Equal(out, []int{1, 3, 5, 7, 9}) {
		t.Error(out)
	}
}

func TestAscendFloor(t *testing.T) {
	var aat *Tree[int, struct{}]
	aat = aat.Add(1).Add(9)
	aat = aat.Add(5).Add(3)
	aat = aat.Add(7)

	var out []int

	out = out[:0]
	for i := range aat.AscendFloor(4) {
		out = append(out, i)
	}
	if !slices.Equal(out, []int{3, 5, 7, 9}) {
		t.Error(out)
	}

	out = out[:0]
	for i := range aat.AscendFloor(3) {
		out = append(out, i)
	}
	if !slices.Equal(out, []int{3, 5, 7, 9}) {
		t.Error(out)
	}

	out = out[:0]
	for i := range aat.AscendFloor(0) {
		out = append(out, i)
	}
	if !slices.Equal(out, []int{1, 3, 5, 7, 9}) {
		t.Error(out)
	}
}

func TestDescendFloor(t *testing.T) {
	var aat *Tree[int, struct{}]
	aat = aat.Add(1).Add(9)
	aat = aat.Add(5).Add(3)
	aat = aat.Add(7)

	var out []int

	out = out[:0]
	for i := range aat.DescendFloor(6) {
		out = append(out, i)
	}
	if !slices.Equal(out, []int{5, 3, 1}) {
		t.Error(out)
	}

	out = out[:0]
	for i := range aat.DescendFloor(7) {
		out = append(out, i)
	}
	if !slices.Equal(out, []int{7, 5, 3, 1}) {
		t.Error(out)
	}

	out = out[:0]
	for i := range aat.DescendFloor(10) {
		out = append(out, i)
	}
	if !slices.Equal(out, []int{9, 7, 5, 3, 1}) {
		t.Error(out)
	}
}

func TestDescendCeil(t *testing.T) {
	var aat *Tree[int, struct{}]
	aat = aat.Add(1).Add(9)
	aat = aat.Add(5).Add(3)
	aat = aat.Add(7)

	var out []int

	out = out[:0]
	for i := range aat.DescendCeil(6) {
		out = append(out, i)
	}
	if !slices.Equal(out, []int{7, 5, 3, 1}) {
		t.Error(out)
	}

	out = out[:0]
	for i := range aat.DescendCeil(7) {
		out = append(out, i)
	}
	if !slices.Equal(out, []int{7, 5, 3, 1}) {
		t.Error(out)
	}

	out = out[:0]
	for i := range aat.DescendCeil(10) {
		out = append(out, i)
	}
	if !slices.Equal(out, []int{9, 7, 5, 3, 1}) {
		t.Error(out)
	}
}
