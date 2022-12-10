package util

func Abs(val int) int {
	if val < 0 {
		return -val
	} else {
		return val
	}
}

func Sign(val int) int {
	if val < 0 {
		return -1
	} else {
		return 1
	}
}
