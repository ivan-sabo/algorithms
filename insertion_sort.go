package main

type InsertionSort struct {
	s []int
}

func (is *InsertionSort) Sort(s []int) ([]int, error) {
	is.s = make([]int, len(s))
	copy(is.s, s)

	l := len(is.s)
	for i := 1; i < l; i++ {
		for j := i; j > 0; j-- {
			if Less(is.s[j], is.s[j-1]) {
				Exchange(is.s, j, j-1)
			}
		}
	}

	return is.s, nil
}
