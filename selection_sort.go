package main

import "log"

// @todo: use generics and interfaces
type SelectionSort struct {
	s []int
}

func (ss *SelectionSort) Sort(a []int) ([]int, error) {
	ss.s = make([]int, len(a))
	copy(ss.s, a)
	n := len(ss.s)
	var err error

	Show(ss.s)

	for i := 0; i < n; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if Less(ss.s[j], ss.s[min]) {
				min = j
			}
		}

		err = ss.exchange(i, min)
		if err != nil {
			log.Fatalf("an error occured: %s, i=%d, min=%d, len(%d)\n", err, i, min, n)
		}
	}

	Show(ss.s)

	return ss.s, nil
}

// @todo: use generics and interfaces
func (ss *SelectionSort) exchange(a, b int) error {
	if a >= len(ss.s) || b >= len(ss.s) {
		return ErrorExchangeIndexOutOfBound
	}

	ss.s[a], ss.s[b] = ss.s[b], ss.s[a]

	return nil
}
