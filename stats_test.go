package aa

import "testing"

func TestTree_Select(t *testing.T) {
	var tt *Tree[int, string]
	tt = tt.Put(1, "one").Put(3, "three").Put(5, "five")
	tt = tt.Put(4, "four").Put(2, "two")
	tt = tt.Put(3, "THREE")

	for i := range 5 {
		if n := tt.Select(i); i+1 != n.Key() {
			t.Errorf("%d ≠ %d", i+1, n.Key())
		}
	}

	if tt.Select(-1) != nil {
		t.Error()
	}
	if tt.Select(tt.Len()) != nil {
		t.Error()
	}
}

func TestTree_Rank(t *testing.T) {
	var tt *Tree[int, string]
	tt = tt.Put(1, "one").Put(3, "three").Put(5, "five")
	tt = tt.Put(4, "four").Put(2, "two")
	tt = tt.Put(3, "THREE")

	if i := tt.Rank(0); i != 0 {
		t.Error(i)
	}
	if i := tt.Rank(1); i != 0 {
		t.Error(i)
	}
	if i := tt.Rank(3); i != 2 {
		t.Error(i)
	}
	if i := tt.Rank(5); i != 4 {
		t.Error(i)
	}
	if i := tt.Rank(6); i != 5 {
		t.Error(i)
	}
}
