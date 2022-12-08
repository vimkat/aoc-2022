package main

import (
	"fmt"
	"strings"

	"github.com/vimkat/aoc-2022/util"
)

func main() {
	fmt.Println("Part 1")
	run1("example.txt")
	run1("puzzle.txt")

	fmt.Println("Part 2")
	run2("example.txt")
	run2("puzzle.txt")
}

func run1(path string) {
	points := 0
	for line := range util.ReadByLines(path) {
		split := strings.Split(line, " ")
		play := parseRPSPlay(util.Str2Rune(split[0]))
		response := parseRPSResponse(util.Str2Rune(split[1]))
		points += getPoints(play, response)
	}
	fmt.Printf("%d\n", points)
}

func run2(path string) {
	points := 0
	for line := range util.ReadByLines(path) {
		split := strings.Split(line, " ")
		play := parseRPSPlay(util.Str2Rune(split[0]))
		outcome := parseOutcome(util.Str2Rune(split[1]))
		response := getChoice(play, outcome)
		points += getPoints(play, response)
	}
	fmt.Printf("%d\n", points)
}
