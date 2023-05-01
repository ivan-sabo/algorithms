package main

import "testing"

func TestLess(t *testing.T) {
	a, b := 1, 2

	if !Less(a, b) {
		t.Errorf("expected %t, got %t", true, false)
	}
}

func TestExchange(t *testing.T) {
	a := []int{1, 2}
	ss := SelectionSort{
		s: a,
	}
	err := Exchange(a, 0, 1)
	if err != nil {
		t.Errorf("unexpected error: %e\n", err)
	}

	expected := []int{2, 1}
	for i := 0; i < len(ss.s); i++ {
		if ss.s[i] != expected[i] {
			t.Errorf("expected: %d, got: %d\n", expected[i], ss.s[i])
		}
	}
}

func TestExchangeErrorExchangeIndexOutOfBound(t *testing.T) {
	a := []int{1, 2}

	err := Exchange(a, 0, 3)
	expected := ErrorExchangeIndexOutOfBound
	if err != expected {
		t.Errorf("expected: %e, got: %e", ErrorExchangeIndexOutOfBound, err)
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
