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
	heightMap := buildMap(path)
	visibles := heightMap.VisibleTrees()
	visibleCount := len(visibles)
	fmt.Println(visibleCount)
}

func part2(path string) {
	heightMap := buildMap(path)
	maxScore := 0
	for row := 1; row < heightMap.Height()-1; row++ {
		for col := 1; col < heightMap.Width()-1; col++ {
			score := heightMap.ScenicScore(row, col)
			if score > maxScore {
				maxScore = score
			}
		}
	}
	fmt.Println(maxScore)
}

func buildMap(path string) *HeightMap {
	heightMap := HeightMap{}
	for line := range util.ReadByLines(path) {
		row := []Tree{}
		for _, treeRune := range line {
			row = append(row, Tree(treeRune-'0'))
		}
		heightMap.AddRow(row)
	}
	return &heightMap
}
