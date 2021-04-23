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

func TestIntersection(t *testing.T) {
	if len(Intersection([]int{}, []int{1})) > 0 {
		t.Fail()
	}
	if len(Intersection([]int{1}, []int{})) > 0 {
		t.Fail()
	}
	if len(Intersection([]int{1}, []int{1, 2})) != 1 || Intersection([]int{1}, []int{1, 2})[0] != 1 {
		t.Fail()
	}
	if len(Intersection([]int{1, 2}, []int{1})) != 1 || Intersection([]int{1}, []int{1, 2})[0] != 1 {
		t.Fail()
	}
}

func TestUnion(t *testing.T) {
	if len(Union([]int{}, []int{1})) != 1 {
		t.Fail()
	}
	if len(Union([]int{1}, []int{})) != 1 {
		t.Fail()
	}
	var ints = Union([]int{1, 2}, []int{1})
	if len(ints) != 2 || ints[0] != 1 || ints[1] != 2 {
		t.Fail()
	}
	ints = Union([]int{1}, []int{1, 2})
	if len(ints) != 2 || ints[0] != 1 || ints[1] != 2 {
		t.Fail()
	}
	ints = Union([]int{1, 3}, []int{2, 3})
	if len(ints) != 3 || ints[0] != 1 || ints[1] != 2 || ints[2] != 3 {
		t.Fail()
	}
}

func TestMinMax(t *testing.T) {
	var a, b = MinMax(1, 2)
	if a != 1 || b != 2 {
		t.Fail()
	}
	a, b = MinMax(2, 1)
	if a != 1 || b != 2 {
		t.Fail()
	}
}
