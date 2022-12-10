package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/vimkat/aoc-2022/util"
)

func main() {
	part1("example.txt")
	part1("puzzle.txt")
	fmt.Println()
	part2("example.txt")
	part2("example-2.txt")
	part2("puzzle.txt")
}

func part1(path string) { simulateRope(path, 2) }
func part2(path string) { simulateRope(path, 10) }

func simulateRope(path string, ropeLength int) {
	rope := NewRope(ropeLength)

	for line := range util.ReadByLines(path) {
		parsed := strings.Split(line, " ")
		repeat := util.MustValue(strconv.Atoi(parsed[1]))

		for i := 0; i < repeat; i++ {
			switch parsed[0] {
			case "L":
				rope.Move(-1, 0)
			case "R":
				rope.Move(1, 0)
			case "U":
				rope.Move(0, -1)
			case "D":
				rope.Move(0, 1)
			}
		}

		// Print steps
		// if strings.HasPrefix(path, "example") {
		// 	rope.PrettyPrint(35, 35, 15, 15)
		// 	fmt.Println()
		// }
	}

	fmt.Println(len(rope.Visited()))
}
