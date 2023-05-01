// ShellSort implementation
//
// ShellSort is the only sorting method considered in "Algorithms" book whose permormance
// on randomly ordered arrays has not been precisely characterized.
//
// Much faster than InsertionSort and SelectionSort.
// It uses no extra space.
package main

type ShellSort struct {
	s []int
}

func (ss *ShellSort) Sort(s []int) ([]int, error) {
	ss.s = make([]int, len(s))
	copy(ss.s, s)

	n := len(s)
	h := 1
	for h < n/3 {
		h = 3*h + 1
	}

	for h >= 1 {
		for i := h; i < n; i++ {
			for j := i; (j >= h) && Less(ss.s[j], ss.s[j-h]); j -= h {
				Exchange(ss.s, j, j-h)
			}
		}
		h /= 3
	}

	return ss.s, nil
}
