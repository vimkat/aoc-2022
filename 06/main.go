package main

import (
	"fmt"
	"os"
)

func main() {
	part1("puzzle.txt")
	part2("puzzle.txt")
}

func part1(path string) {
	content, _ := os.ReadFile(path)
	pos := FindSOP(string(content)) + 1
	fmt.Println(pos)
}

func part2(path string) {
	content, _ := os.ReadFile(path)
	pos := FindSOM(string(content)) + 1
	fmt.Println(pos)
}
