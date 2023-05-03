package sorting

import "testing"

func TestMergeSortBottomup(t *testing.T) {
	s := []int{5, 2, 4, 1, 3}

	is := MergeSortBottomup{}
	sorted := is.Sort(s)

	expected := []int{1, 2, 3, 4, 5}
	if len(sorted) != len(expected) {
		t.Fatalf("resulting array size(%d) is different from expected size(%d)",
			len(sorted),
			len(expected),
		)
	}

	for i := 0; i < len(s); i++ {
		if expected[i] != sorted[i] {
			t.Errorf("expected: %d, got: %d", expected[i], sorted[i])
		}
	}
}
