package main

import "testing"

func TestLess(t *testing.T) {
	a, b := 1, 2

	if !Less(a, b) {
		t.Errorf("expected %t, got %t", true, false)
	}
}
