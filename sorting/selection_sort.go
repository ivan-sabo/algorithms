// SelectionSort implementation.
//
// Uses ~N^2/2 compares and n exchanges to sort an array of length n.
//
// It takes about as long to run selection sort for an array that is already in order or for an
// array with all keys exaual as it does for a randomly-ordered array.
//
// Data movement is minimal. The number of exchanges is a linear function of the array length.
//
// Running time is N^2 for randomly ordered arrays.
package sorting

import (
	"log"

	"github.com/ivan-sabo/algorithms/common"
)

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
			if common.Less(ss.s[j], ss.s[min]) {
				min = j
			}
		}

		err = common.Exchange(ss.s, i, min)
		if err != nil {
			log.Fatalf("an error occured: %s, i=%d, min=%d, len(%d)\n", err, i, min, n)
		}
	}

	return ss.s, nil
}
