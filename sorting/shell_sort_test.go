package sorting

import (
	"math/rand"
	"testing"
	"time"
)

func TestShellSort(t *testing.T) {
	unsorted := rand.New(rand.NewSource(time.Now().Unix())).Perm(5)

	ss := ShellSort{}
	sorted, err := ss.Sort(unsorted)

	if err != nil {
		t.Fatalf("unexpected error: %e", err)
	}

	expected := []int{0, 1, 2, 3, 4}
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
