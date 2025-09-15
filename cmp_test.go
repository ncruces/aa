package aa

import "testing"

func TestEqual(t *testing.T) {
	equal := func(a, b *Tree[int, struct{}], contains bool) {
		t.Helper()
		if Equal(a, b) != contains {
			t.Error()
		}
	}

	set := MakeSet(1, 2, 3, 4, 5)

	equal(nil, nil, true)
	equal(set, set, true)
	equal(nil, set, false)
	equal(set, nil, false)

	equal(set, set.Add(0), false)
	equal(set, set.Add(6), false)
	equal(set.Add(0), set, false)
	equal(set.Add(6), set, false)

	equal(set.Add(0), set.Add(6), false)
	equal(set.Add(6), set.Add(7), false)

	equal(set.Add(0).Add(6), set.Add(6).Add(0), true)
}

func TestOverlap(t *testing.T) {
	overlap := func(a, b *Tree[int, struct{}], contains bool) {
		t.Helper()
		if Overlap(a, b) != contains {
			t.Error()
		}
	}

	set := MakeSet(1, 2, 3, 4, 5)

	overlap(nil, nil, false)
	overlap(set, set, true)
	overlap(nil, set, false)
	overlap(set, nil, false)

	overlap(set, set.Add(0), true)
	overlap(set, set.Add(6), true)
	overlap(set.Add(0), set, true)
	overlap(set.Add(6), set, true)

	overlap(MakeSet(3), set, true)
	overlap(set, MakeSet(3), true)

	overlap(MakeSet(2, 4), set, true)
	overlap(set, MakeSet(2, 4), true)

	overlap(MakeSet(0, 1), set, true)
	overlap(set, MakeSet(5, 6), true)

	overlap(MakeSet(0, 6), set, false)
	overlap(set, MakeSet(0, 6), false)
}

func TestSubset(t *testing.T) {
	subset := func(a, b *Tree[int, struct{}], contains bool) {
		t.Helper()
		if Subset(a, b) != contains {
			t.Error()
		}
	}

	set := MakeSet(1, 2, 3, 4, 5)

	subset(nil, nil, true)
	subset(set, set, true)
	subset(nil, set, true)
	subset(set, nil, false)

	subset(set, set.Add(0), true)
	subset(set, set.Add(6), true)
	subset(set.Add(0), set, false)
	subset(set.Add(6), set, false)

	subset(MakeSet(3), set, true)
	subset(set, MakeSet(3), false)

	subset(MakeSet(2, 4), set, true)
	subset(set, MakeSet(2, 4), false)

	subset(MakeSet(1, 2), set, true)
	subset(MakeSet(4, 5), set, true)
	subset(MakeSet(0, 1, 2), set, false)
	subset(MakeSet(4, 5, 6), set, false)
}

func TestEqualSubset_value(t *testing.T) {
	var tt *Tree[int, string]
	tt = tt.Put(1, "one").Put(3, "three").Put(5, "five")
	tt = tt.Put(4, "four").Put(2, "two")

	if Equal(tt, tt.Put(3, "THREE")) {
		t.Error()
	}
	if Subset(tt, tt.Put(3, "THREE")) {
		t.Error()
	}
	if !Overlap(tt, tt.Put(3, "THREE")) {
		t.Error()
	}
}
