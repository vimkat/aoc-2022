package main

import (
	"fmt"

	"github.com/vimkat/aoc-2022/util"
)

func main() {
	part1("example.txt")
	part1("puzzle.txt")
	fmt.Println()
	part2("example.txt")
	part2("puzzle.txt")
}

func part1(path string) {
	count := 0
	for line := range util.ReadByLines(path) {
		p, _ := ParsePair(line)
		if p.FullyOverlap() {
			count++
		}
	}
	fmt.Println(count)
}

func part2(path string) {
	count := 0
	for line := range util.ReadByLines(path) {
		p, _ := ParsePair(line)
		if p.PartiallyOverlap() {
			count++
		}
	}
	fmt.Println(count)
}
