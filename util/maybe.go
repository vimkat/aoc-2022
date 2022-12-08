package util

type maybe[T any] struct {
	value T
	err   error
}

func Maybe[T any](value T, err error) maybe[T] {
	return maybe[T]{value, err}
}

func (m maybe[T]) Value() T {
	return m.value
}

func (m maybe[T]) Err() error {
	return m.err
}
