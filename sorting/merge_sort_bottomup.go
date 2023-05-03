package sorting

import (
	"github.com/ivan-sabo/algorithms/common"
)

type MergeSortBottomup struct {
	s   []int
	aux []int
}

func (msb *MergeSortBottomup) Sort(s []int) []int {
	msb.s = make([]int, len(s))
	msb.aux = make([]int, len(s))

	copy(msb.s, s)

	n := len(msb.s)
	for len := 1; len < n; len *= 2 {
		for low := 0; low < n-len; low += len + len {
			msb.merge(low, low+len-1, common.MinInt((low+(2*len)-1), n-1))
		}
	}

	return msb.s
}

func (msb *MergeSortBottomup) merge(low, mid, high int) error {
	if high >= len(msb.s) {
		return common.ErrorIndexOutOfBound
	}

	for i := low; i <= high; i++ {
		msb.aux[i] = msb.s[i]
	}

	first, second := low, mid+1
	for i := low; i <= high; i++ {
		if first > mid {
			msb.s[i] = msb.aux[second]
			second++
			continue
		}

		if second > high {
			msb.s[i] = msb.aux[first]
			first++
			continue
		}

		if common.Less(msb.aux[second], msb.aux[first]) {
			msb.s[i] = msb.aux[second]
			second++
			continue
		}

		msb.s[i] = msb.aux[first]
		first++
	}

	return nil
}
