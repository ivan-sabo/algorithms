package main

import "testing"

func TestMergeSort(t *testing.T) {
	unsorted := []int{5, 4, 1, 3, 2}

	ms := MergeSort{}
	sorted, err := ms.Sort(unsorted)
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

func TestMerge(t *testing.T) {
	ms := MergeSort{
		s:   []int{1, 2, 4, 3, 5, 6, 7},
		aux: []int{1, 2, 4, 3, 5, 6, 7},
	}

	err := ms.merge(0, 2, 6)
	if err != nil {
		t.Errorf("unexpected error: %e\n", err)
	}

	expected := []int{1, 2, 3, 4, 5, 6, 7}
	if len(ms.aux) != len(expected) {
		t.Fatalf("resulting array size(%d) is different from expected size(%d)",
			len(ms.aux),
			len(expected),
		)
	}

	for i := 0; i < len(ms.s); i++ {
		if expected[i] != ms.s[i] {
			t.Errorf("expected: %d, got: %d", expected[i], ms.aux[i])
		}
	}
}

func TestErrorIndexOutOfBound(t *testing.T) {
	ms := MergeSort{
		s:   []int{1, 2, 4, 3, 5, 6, 7},
		aux: []int{1, 2, 4, 3, 5, 6, 7},
	}

	err := ms.merge(0, 2, 7)
	if err != ErrorIndexOutOfBound {
		t.Errorf("expected error: %e, got: %e\n", ErrorIndexOutOfBound, err)
	}
}
