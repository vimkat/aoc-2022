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
	cpu := NewCPU()
	total := 0

	lineNr := 1
	for line := range util.ReadByLines(path) {
		op, err := ParseInstruction(line)
		if err != nil {
			fmt.Printf("Parse error line %d: %v\n", lineNr, err)
			return
		}

		cpu.ScheduleOperation(op)
		cpu.Tick()
		lineNr += 1

		// check reporting
		if cpu.Cycle()%40 == 20 {
			// fmt.Printf("#%3d: %d\n", cpu.Cycle(), cpu.SignalStrength())
			total += cpu.SignalStrength()
		}
	}

	for !cpu.IsHalted() {
		cpu.Tick()
		if cpu.Cycle()%40 == 20 {
			// fmt.Printf("#%3d: %d\n", cpu.Cycle(), cpu.SignalStrength())
			total += cpu.SignalStrength()
		}
	}

	fmt.Println(total)
	// fmt.Println("-----")
}

func part2(path string) {
	computer := NewComputer()

	lineNr := 1
	for line := range util.ReadByLines(path) {
		err := computer.ReadInstruction(line)
		if err != nil {
			fmt.Printf("error line %d: %v", lineNr, err)
			return
		}
		lineNr += 1
	}
	computer.RunToHalt()
	computer.Show()
	fmt.Println()
}
