package aa

import "testing"

func TestTree_Select(t *testing.T) {
	var aat *Tree[int, string]
	aat = aat.Put(1, "one").Put(3, "three").Put(5, "five")
	aat = aat.Put(4, "four").Put(2, "two")
	aat = aat.Put(3, "THREE")

	for i := range 5 {
		if n := aat.Select(i); i+1 != n.Key() {
			t.Errorf("%d ≠ %d", i+1, n.Key())
		}
	}

	if aat.Select(-1) != nil {
		t.Error()
	}
	if aat.Select(aat.Len()) != nil {
		t.Error()
	}
}

func TestTree_Rank(t *testing.T) {
	var aat *Tree[int, string]
	aat = aat.Put(1, "one").Put(3, "three").Put(5, "five")
	aat = aat.Put(4, "four").Put(2, "two")
	aat = aat.Put(3, "THREE")

	if i := aat.Rank(0); i != 0 {
		t.Error(i)
	}
	if i := aat.Rank(1); i != 0 {
		t.Error(i)
	}
	if i := aat.Rank(3); i != 2 {
		t.Error(i)
	}
	if i := aat.Rank(5); i != 4 {
		t.Error(i)
	}
	if i := aat.Rank(6); i != 5 {
		t.Error(i)
	}
}
