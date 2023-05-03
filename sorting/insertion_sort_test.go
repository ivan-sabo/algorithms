package sorting

import (
	"math/rand"
	"testing"
	"time"
)

func TestInsertionSort(t *testing.T) {
	s := rand.New(rand.NewSource(time.Now().Unix())).Perm(5)
	is := InsertionSort{}
	sorted, err := is.Sort(s)
	if err != nil {
		t.Errorf("unexpected error: %e", err)
	}

	expected := []int{0, 1, 2, 3, 4}
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
