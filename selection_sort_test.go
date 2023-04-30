package main

import "testing"

func TestExchange(t *testing.T) {
	a := []int{1, 2}
	ss := SelectionSort{
		s: a,
	}
	ss.exchange(0, 1)

	expected := []int{2, 1}
	for i := 0; i < len(ss.s); i++ {
		if ss.s[i] != expected[i] {
			t.Errorf("expected: %d, got: %d\n", expected[i], ss.s[i])
		}
	}
}

func TestSort(t *testing.T) {
	unsorted := []int{5, 4, 1, 3, 2}

	ss := SelectionSort{}
	sorted, err := ss.Sort(unsorted)

	if err != nil {
		t.Fatalf("unexpected error: %e", err)
	}

	expected := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(sorted); i++ {
		if sorted[i] != expected[i] {
			t.Errorf("unexpected error, slice is not sorted: got=%d, expected=%d",
				sorted[i],
				expected[i],
			)
		}
	}
}
