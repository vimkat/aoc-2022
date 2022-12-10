package util

func Sum(arr []int) int {
	s := 0
	for _, v := range arr {
		s += v
	}
	return s
}

func Max(arr []int) int {
	max := arr[0]
	for _, val := range arr {
		if val > max {
			max = val
		}
	}
	return max
}

func FindDupliacte[T comparable](slices ...[]T) (result T, found bool) {
	for _, element := range slices[0] {
		if AllHaveElement(element, slices[1:]) {
			return element, true
		}
	}
	return // zero-value, false
}

func AllHaveElement[T comparable](element T, slices [][]T) bool {
	for _, slice := range slices {
		contains := false
		for _, otherElement := range slice {
			if otherElement == element {
				contains = true
				break
			}
		}
		if !contains {
			return false
		}
	}
	return true
}

func Unique[T comparable](slice []T) bool {
	for i := 0; i < len(slice); i++ {
		for j := 0; j < i; j++ {
			if slice[i] == slice[j] {
				return false
			}
		}
	}

	return true
}

func Any[T any](slice []T, predicate func(*T) bool) bool {
	for _, element := range slice {
		if predicate(&element) {
			return true
		}
	}
	return false
}
