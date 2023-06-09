package common

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

// @todo: use generics and interfaces
func IsSorted(a []int) bool {
	if len(a) < 2 {
		return true
	}

	for i := 1; i < len(a); i++ {
		if Less(a[i], a[i-1]) {
			return false
		}
	}

	return true
}

// @todo: use generics and interfaces
func Less(a, b int) bool {
	return a < b
}

func ToString(a []int) string {
	b := strings.Builder{}
	for i, v := range a {
		if i == len(a)-1 {
			b.WriteString(fmt.Sprint(v))
			break
		}
		b.WriteString(fmt.Sprint(v) + " ")
	}

	return b.String()
}

// @todo: use generics and interfaces
func Exchange(s []int, a, b int) error {
	if a >= len(s) || b >= len(s) {
		return ErrorIndexOutOfBound
	}

	s[a], s[b] = s[b], s[a]

	return nil
}

func MinInt(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

var (
	ErrorIndexOutOfBound = errors.New("referenced index out of bound")
)
