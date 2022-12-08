package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Range struct {
	start, end int
}

func ParseRange(str string) (Range, error) {
	split := strings.SplitN(str, "-", 2)
	start, err := strconv.Atoi(split[0])
	if err != nil {
		return Range{}, fmt.Errorf("couldn't parse start of range: %v", err)
	}
	end, err := strconv.Atoi(split[1])
	if err != nil {
		return Range{}, fmt.Errorf("couldn't parse end of range: %v", err)
	}

	if end < start {
		return Range{}, fmt.Errorf("invalid range, end (%v) < start (%v)", end, start)
	}

	return Range{start, end}, nil
}

func (r Range) String() string {
	return fmt.Sprintf("%d-%d", r.start, r.end)
}

func (r Range) Equals(other Range) bool {
	return r.start == other.start && r.end == other.end
}

func ContainsRange(outer, inner Range) bool {
	return outer.start <= inner.start && outer.end >= inner.end
}

func RangeOverlaps(a, b Range) bool {
	return a.start >= b.start && a.start <= b.end || a.end >= b.start && a.end <= b.end
}

type Pair struct {
	elf1 Range
	elf2 Range
}

func ParsePair(str string) (Pair, error) {
	split := strings.SplitN(str, ",", 2)
	elf1, err := ParseRange(split[0])
	if err != nil {
		return Pair{}, fmt.Errorf("couldn't parse first range: %v", err)
	}
	elf2, err := ParseRange(split[1])
	if err != nil {
		return Pair{}, fmt.Errorf("couldn't parse second range: %v", err)
	}

	return Pair{elf1, elf2}, nil
}

func (p Pair) String() string {
	return fmt.Sprintf("%s,%s", p.elf1, p.elf2)
}

func (p Pair) Equals(other Pair) bool {
	return p.elf1.Equals(other.elf1) && p.elf2.Equals(other.elf2)
}

func (p Pair) FullyOverlap() bool {
	return ContainsRange(p.elf1, p.elf2) || ContainsRange(p.elf2, p.elf1)
}

func (p Pair) PartiallyOverlap() bool {
	return RangeOverlaps(p.elf1, p.elf2) || RangeOverlaps(p.elf2, p.elf1)
}
