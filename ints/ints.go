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
