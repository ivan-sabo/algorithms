package sorting

import "github.com/ivan-sabo/algorithms/common"

type QuickSort struct {
	s []int
}

func (qs *QuickSort) Sort(s []int) []int {
	qs.s = make([]int, len(s))
	copy(qs.s, s)

	qs.helperSort(0, len(s)-1)

	return qs.s
}

func (qs *QuickSort) helperSort(low, high int) {
	if high < low {
		return
	}
	j := qs.partition(low, high)
	qs.helperSort(low, j-1)
	qs.helperSort(j+1, high)
}

func (qs *QuickSort) partition(low, high int) int {
	left, right := low, high+1
	v := qs.s[low]
	for {
		for {
			left++
			if common.Less(qs.s[left], v) || (left == high) {
				break
			}
		}
		for {
			right--
			if common.Less(v, qs.s[right]) || (right == low) {
				break
			}
		}
		if left >= right {
			break
		}
		common.Exchange(qs.s, left, right)
	}
	common.Exchange(qs.s, low, right)

	return right
}
