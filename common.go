package main

import "fmt"

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

func Show(a []int) {
	for i, v := range a {
		if i == len(a)-1 {
			fmt.Printf("%d\n", v)
			break
		}
		fmt.Printf("%d ", v)
	}
}
