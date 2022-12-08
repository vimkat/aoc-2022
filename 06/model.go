package main

import (
	"github.com/vimkat/aoc-2022/util"
)

type RingBuffer[T any] []T

func NewRingBuffer[T any](size int) RingBuffer[T] {
	return make([]T, size)
}

func (b *RingBuffer[T]) Push(newElement T) {
	for i := 0; i < len(*b)-1; i++ {
		(*b)[i] = (*b)[i+1]
	}
	(*b)[len(*b)-1] = newElement
}

func (b *RingBuffer[T]) Last() T {
	return (*b)[len(*b)-1]
}

func (b *RingBuffer[T]) LastN(n int) []T {
	return (*b)[len(*b)-n:]
}

func findFirstUniqueN[T comparable](elements []T, n int) int {
	// No need to process short singals
	if len(elements) < n {
		return -1
	}

	buf := RingBuffer[T](elements[:n])

	for i := n; i < len(elements); i++ {
		buf.Push(elements[i])

		// SOP marker found
		if util.Unique(buf.LastN(n)) {
			return i
		}
	}

	return -1
}

func FindSOP(s string) int {
	return findFirstUniqueN([]rune(s), 4)
}

func FindSOM(s string) int {
	return findFirstUniqueN([]rune(s), 14)
}
