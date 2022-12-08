package main

import "testing"

func TestRingBufferPush(t *testing.T) {
	buf := NewRingBuffer[int](4)
	buf.Push(1)
	buf.Push(2)
	buf.Push(3)
	buf.Push(4)

	// Check equality
	for i := range buf {
		if buf[i] != i+1 {
			t.Fatalf("buffer at position %d is %d, expected %d", i, buf[i], i+1)
		}
	}

	// Check after overflow
	buf.Push(5)
	for i := range buf {
		if buf[i] != i+2 {
			t.Fatalf("buffer at position %d is %d, expected %d", i, buf[i], i+2)
		}
	}
}

func TestFindSOP(t *testing.T) {
	cases := map[string]int{
		"mjqjpqmgbljsphdztnvjfqwrcgsmlb":    7,
		"bvwbjplbgvbhsrlpgdmjqwftvncz":      5,
		"nppdvjthqldpwncqszvftbrmjlhg":      6,
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg": 10,
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw":  11,
	}

	for input, pos := range cases {
		t.Run(input, func(t *testing.T) {
			found := FindSOP(input) + 1
			if found != pos {
				t.Errorf("found SOP marker for %v at position %v (%v), expected %v",
					input, found, input[found:found+4], pos)
			}
		})
	}
}

func TestFindSOM(t *testing.T) {
	cases := map[string]int{
		"mjqjpqmgbljsphdztnvjfqwrcgsmlb":    19,
		"bvwbjplbgvbhsrlpgdmjqwftvncz":      23,
		"nppdvjthqldpwncqszvftbrmjlhg":      23,
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg": 29,
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw":  26,
	}

	for input, pos := range cases {
		t.Run(input, func(t *testing.T) {
			found := FindSOM(input) + 1
			if found != pos {
				t.Errorf("found SOM marker for %v at position %v (%v), expected %v",
					input, found, input[found:found+14], pos)
			}
		})
	}
}

func TestRingBufferLastN(t *testing.T) {
	buf := NewRingBuffer[int](4)
	buf.Push(1)
	buf.Push(2)
	buf.Push(3)
	buf.Push(4)

	last4 := buf.LastN(4)

	// Check equality
	for i := range last4 {
		if buf[i] != i+1 {
			t.Fatalf("buffer at position %d is %d, expected %d", i, buf[i], i+1)
		}
	}
}
