package main

import "testing"

func TestInsertionSort(t *testing.T) {
	s := []int{5, 2, 4, 1, 3}
	expected := []int{1, 2, 3, 4, 5}

	is := InsertionSort{}
	r, err := is.Sort(s)
	if err != nil {
		t.Errorf("unexpected error: %e", err)
	}

	for i := 0; i < len(s); i++ {
		if expected[i] != r[i] {
			t.Errorf("expected: %d, got: %d", expected[i], r[i])
		}
	}
}
