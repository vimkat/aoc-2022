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
	part2("puzzle.txt")
}

func part1(path string) {
	total := 0
	root := parseTerminal(path)
	root.WalkFiltered(IsDirectory, func(node FSNode) {
		size := node.Size()
		if size <= 100000 {
			total += size
		}
	})
	fmt.Println(total)
}

func part2(path string) {
	totalSpace := 70000000
	neededSpace := 30000000

	root := parseTerminal(path)
	rootSize := root.Size()
	freeSpace := totalSpace - rootSize
	needsFreeing := neededSpace - freeSpace

	fmt.Printf("%d/%d in use, %d available\n", rootSize, totalSpace, freeSpace)

	dirToDeleteSize := rootSize
	root.WalkFiltered(IsDirectory, func(node FSNode) {
		size := node.Size()
		if size >= needsFreeing && size < dirToDeleteSize {
			dirToDeleteSize = size
		}
	})

	fmt.Println(dirToDeleteSize)
}

func parseTerminal(path string) *Directory {
	root := NewDirectory("/")
	current := root

	for line := range util.ReadByLines(path) {
		split := strings.Split(line, " ")
		if split[0] == "$" {
			// Execute command
			switch split[1] {
			case "cd":
				switch split[2] {
				case "..":
					current = current.Parent()
				case "/":
					current = root
				default:
					node, _ := current.Get(split[2])
					current = node.(*Directory)
				}
			case "ls":
				break // no need to handle this explicitly
			default:
				panic(fmt.Sprintf("command %s unknown", split[1]))
			}

		} else {
			// Receive output
			if split[0] == "dir" {
				// is a directory
				current.Add(NewDirectory(split[1]))
			} else {
				// is a file
				size, _ := strconv.Atoi(split[0])
				current.Add(NewFile(split[1], size))
			}
		}
	}

	return root
}
