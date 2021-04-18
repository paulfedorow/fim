package ints

import "testing"

func TestEquals(t *testing.T) {
	if !Equals([]int{1, 2, 3}, []int{1, 2, 3}) {
		t.Fail()
	}
	if Equals([]int{}, []int{1, 2, 3}) {
		t.Fail()
	}
	if Equals([]int{1, 2, 3}, []int{}) {
		t.Fail()
	}
	if Equals([]int{3, 2, 1}, []int{1, 2, 3}) {
		t.Fail()
	}
}

func TestLess(t *testing.T) {
	if !Less([]int{}, []int{1}) {
		t.Fail()
	}
	if Less([]int{1}, []int{}) {
		t.Fail()
	}
	if !Less([]int{1}, []int{2}) {
		t.Fail()
	}
	if Less([]int{2}, []int{1}) {
		t.Fail()
	}
	if !Less([]int{1}, []int{1, 2}) {
		t.Fail()
	}
	if Less([]int{1, 2}, []int{1}) {
		t.Fail()
	}
}