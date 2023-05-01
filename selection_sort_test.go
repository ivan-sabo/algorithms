package main

import "testing"

func TestSelectionSort(t *testing.T) {
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
