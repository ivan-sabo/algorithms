package main

import "testing"

func TestShellSort(t *testing.T) {
	unsorted := []int{5, 4, 1, 3, 2}

	ss := ShellSort{}
	sorted, err := ss.Sort(unsorted)

	if err != nil {
		t.Fatalf("unexpected error: %e", err)
	}

	expected := []int{1, 2, 3, 4, 5}
	if len(sorted) != len(expected) {
		t.Fatalf("resulting array size(%d) is different from expected size(%d)",
			len(sorted),
			len(expected),
		)
	}

	for i := 0; i < len(expected); i++ {
		if sorted[i] != expected[i] {
			t.Errorf("unexpected error, slice is not sorted: got=%d, expected=%d",
				sorted[i],
				expected[i],
			)
		}
	}
}
