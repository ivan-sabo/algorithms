package common

import (
	"testing"
)

func TestLess(t *testing.T) {
	a, b := 1, 2

	if !Less(a, b) {
		t.Errorf("expected %t, got %t", true, false)
	}
}

func TestExchange(t *testing.T) {
	a := []int{1, 2}

	err := Exchange(a, 0, 1)
	if err != nil {
		t.Errorf("unexpected error: %e\n", err)
	}

	expected := []int{2, 1}
	for i := 0; i < len(a); i++ {
		if a[i] != expected[i] {
			t.Errorf("expected: %d, got: %d\n", expected[i], a[i])
		}
	}
}

func TestExchangeErrorExchangeIndexOutOfBound(t *testing.T) {
	a := []int{1, 2}

	err := Exchange(a, 0, 3)
	expected := ErrorIndexOutOfBound
	if err != expected {
		t.Errorf("expected: %e, got: %e", ErrorIndexOutOfBound, err)
	}
}

func TestToString(t *testing.T) {
	a := []int{1, 2, 3, 4}
	s := ToString(a)
	expected := "1 2 3 4"

	if s != expected {
		t.Errorf("expected: %s, got: %s", expected, s)
	}
}
