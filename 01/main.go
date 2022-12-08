package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/vimkat/aoc-2022/util"
)

func main() {
	c := caloriesPerElve("puzzle.txt")
	sort.Sort(sort.Reverse(sort.IntSlice(c)))
	fmt.Println(c[0])
	fmt.Println(util.Sum(c[0:3]))
}

func caloriesPerElve(path string) []int {
	calories := []int{0}
	elf := 0

	for line := range util.ReadByLines(path) {
		if line == "" {
			elf++
			calories = append(calories, 0)
		} else {
			i, err := strconv.Atoi(line)
			util.Must(err)
			calories[elf] += i
		}
	}

	return calories
}
