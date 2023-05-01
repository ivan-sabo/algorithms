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

	for i := 0; i < n; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if Less(ss.s[j], ss.s[min]) {
				min = j
			}
		}

		err = Exchange(ss.s, i, min)
		if err != nil {
			log.Fatalf("an error occured: %s, i=%d, min=%d, len(%d)\n", err, i, min, n)
		}
	}

	return ss.s, nil
}
