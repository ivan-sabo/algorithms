// MergeSort implementation.
//
// MergeSort uses between  ~1/2 * N * lg(N) and N * lg(N) compares to sort any array of
// length N.
//
// Drawback is that it requires extra space proportional to N for the auxiliary array for
// merging.
package main

import "fmt"

type MergeSort struct {
	s   []int
	aux []int
}

func (ms *MergeSort) Sort(s []int) ([]int, error) {
	ms.s = make([]int, len(s))
	ms.aux = make([]int, len(s))
	copy(ms.s, s)
	copy(ms.aux, s)

	ms.helperSort(0, len(ms.s)-1)

	fmt.Println(ToString(ms.s))

	return ms.s, nil
}

func (ms *MergeSort) helperSort(low, high int) {
	if high <= low {
		return
	}

	mid := low + (high-low)/2

	ms.helperSort(low, mid)
	ms.helperSort(mid+1, high)
	ms.merge(low, mid, high)
}

func (ms *MergeSort) merge(low, mid, high int) error {
	if high >= len(ms.s) {
		return ErrorIndexOutOfBound
	}

	for i := low; i <= high; i++ {
		ms.aux[i] = ms.s[i]
	}

	first, second := low, mid+1
	for i := low; i <= high; i++ {
		if first > mid {
			ms.s[i] = ms.aux[second]
			second++
			continue
		}

		if second > high {
			ms.s[i] = ms.aux[first]
			first++
			continue
		}

		if Less(ms.aux[second], ms.aux[first]) {
			ms.s[i] = ms.aux[second]
			second++
			continue
		}

		ms.s[i] = ms.aux[first]
		first++
	}

	return nil
}
