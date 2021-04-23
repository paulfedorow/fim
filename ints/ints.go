package ints

// Less returns true if ints1 is lexicographically less than ints2, false otherwise.
func Less(ints1 []int, ints2 []int) bool {
	if len(ints1) == len(ints2) {
		for i := 0; i < len(ints1); i += 1 {
			if ints1[i] < ints2[i] {
				return true
			}
		}
		return false
	} else {
		return len(ints1) < len(ints2)
	}
}

// Equals returns true if the contents of ints1 und ints2 are identical, false otherwise.
func Equals(ints1 []int, ints2 []int) bool {
	if len(ints1) == len(ints2) {
		for i := 0; i < len(ints1); i += 1 {
			if ints1[i] != ints2[i] {
				return false
			}
		}
		return true
	} else {
		return false
	}
}

// Intersection return the intersection of ints1 und ints2. Both arguments are expected to be sorted.
func Intersection(ints1 []int, ints2 []int) []int {
	var result []int
	var i = 0
	var j = 0
	for i < len(ints1) && j < len(ints2) {
		switch {
		case ints1[i] < ints2[j]:
			i += 1
		case ints1[i] > ints2[j]:
			j += 1
		case ints1[i] == ints2[j]:
			result = append(result, ints1[i])
			i += 1
			j += 1
		}
	}
	return result
}

// Union return the union of ints1 und ints2. Both arguments are expected to be sorted.
func Union(ints1 []int, ints2 []int) []int {
	var result []int
	var i = 0
	var j = 0
	for i < len(ints1) && j < len(ints2) {
		switch {
		case ints1[i] < ints2[j]:
			result = append(result, ints1[i])
			i += 1
		case ints1[i] > ints2[j]:
			result = append(result, ints2[j])
			j += 1
		case ints1[i] == ints2[j]:
			result = append(result, ints1[i])
			i += 1
			j += 1
		}
	}
	for i < len(ints1) {
		result = append(result, ints1[i])
		i += 1
	}
	for j < len(ints2) {
		result = append(result, ints2[j])
		j += 1
	}
	return result
}

// MinMax returns the smaller integer first and the larger integer second.
func MinMax(a int, b int) (int, int) {
	if a < b {
		return a, b
	} else {
		return b, a
	}
}
