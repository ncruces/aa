package aa

import (
	"slices"
	"testing"
)

func TestUnion(t *testing.T) {
	t1 := FromSorted([]Entry[int, int]{{1, 1}, {2, 2}, {3, 3},})
	t2 := FromSorted([]Entry[int, int]{{2, 22}, {3, 33}, {4, 4},})
	got := Union(t1, t2).Entries()
	want := []Entry[int, int]{{1, 1}, {2, 2}, {3, 3}, {4, 4}}
	if !slices.Equal(got, want) {
		t.Errorf("Union() = %v, want %v", got, want)
	}
}

func TestIntersection(t *testing.T) {
	t1 := FromSorted([]Entry[int, int]{{1, 1}, {2, 2}, {3, 3},})
	t2 := FromSorted([]Entry[int, int]{{2, 22}, {3, 33}, {4, 4},})
	got := Intersection(t1, t2).Entries()
	want := []Entry[int, int]{{2, 2}, {3, 3}}
	if !slices.Equal(got, want) {
		t.Errorf("Intersection() = %v, want %v", got, want)
	}
}

func TestDifference(t *testing.T) {
	t1 := FromSorted([]Entry[int, int]{{1, 1}, {2, 2}, {3, 3},})
	t2 := FromSorted([]Entry[int, int]{{2, 22}, {3, 33}, {4, 4},})
	got := Difference(t1, t2).Entries()
	want := []Entry[int, int]{{1, 1}}
	if !slices.Equal(got, want) {
		t.Errorf("Difference() = %v, want %v", got, want)
	}
}

func TestSymmetricDifference(t *testing.T) {
	t1 := FromSorted([]Entry[int, int]{{1, 1}, {2, 2}, {3, 3},})
	t2 := FromSorted([]Entry[int, int]{{2, 22}, {3, 33}, {4, 4},})
	got := SymmetricDifference(t1, t2).Entries()
	want := []Entry[int, int]{{1, 1}, {4, 4}}
	if !slices.Equal(got, want) {
		t.Errorf("SymmetricDifference() = %v, want %v", got, want)
	}
}