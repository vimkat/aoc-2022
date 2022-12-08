package util

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

func MustValue[T any](value T, err error) T {
	Must(err)
	return value
}
