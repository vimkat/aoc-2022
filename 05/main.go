package main

import (
	"fmt"
	"strings"

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
	shipLines := []string{}
	ship := Ship{}

	for line := range util.ReadByLines(path) {
		if len(ship) == 0 {
			// Parse ship
			if line != "" {
				shipLines = append(shipLines, line)
			} else {
				ship, _ = ParseShipByLines(shipLines)
			}
		} else {
			// Parse move ops
			op, _ := ParseMoveOp(line)
			ship.Execute9000(op)
		}
	}

	builder := strings.Builder{}
	for _, stack := range ship {
		builder.WriteRune(rune(stack.Peek()))
	}
	fmt.Println(builder.String())
}

func part2(path string) {
	shipLines := []string{}
	ship := Ship{}

	for line := range util.ReadByLines(path) {
		if len(ship) == 0 {
			// Parse ship
			if line != "" {
				shipLines = append(shipLines, line)
			} else {
				ship, _ = ParseShipByLines(shipLines)
			}
		} else {
			// Parse move ops
			op, _ := ParseMoveOp(line)
			ship.Execute9001(op)
		}
	}

	builder := strings.Builder{}
	for _, stack := range ship {
		builder.WriteRune(rune(stack.Peek()))
	}
	fmt.Println(builder.String())
}
