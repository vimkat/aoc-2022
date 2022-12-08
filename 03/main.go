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
	priorities := 0
	for line := range util.ReadByLines(path) {
		rucksack := Rucksack(line)

		duplicate, _ := FindDuplicate(rucksack)
		priorities += duplicate.Priority()
	}
	fmt.Println(priorities)
}

func part2(path string) {
	priorities := 0
	currentRucksack := 0
	currentGroup := Group([3]Rucksack{})
	for line := range util.ReadByLines(path) {
		currentGroup[currentRucksack] = Rucksack(line)
		currentRucksack++

		// Group complete, get badge
		if currentRucksack == 3 {
			badge, _ := currentGroup.Badge()
			priorities += badge.Priority()
			currentRucksack = 0
		}
	}
	fmt.Println(priorities)
}
