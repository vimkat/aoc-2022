package main

import (
	"fmt"
	"strconv"
	"strings"
)

type CPU struct {
	x        int
	ip       int
	cycle    int
	pipeline []Operation
}

func NewCPU() *CPU {
	return &CPU{
		x:        1,
		ip:       0,
		cycle:    1,
		pipeline: []Operation{},
	}
}

func (cpu CPU) X() int              { return cpu.x }
func (cpu CPU) Cycle() int          { return cpu.cycle }
func (cpu CPU) SignalStrength() int { return cpu.cycle * cpu.x }

func (cpu *CPU) ScheduleOperation(op Operation) {
	cpu.pipeline = append(cpu.pipeline, op)
}

func (cpu *CPU) Tick() {
	done := cpu.pipeline[cpu.ip].tick(cpu)
	if done {
		cpu.ip += 1
	}
	cpu.cycle += 1
}

func (cpu CPU) IsHalted() bool {
	return cpu.ip == len(cpu.pipeline)-1
}

type Operation struct {
	source     string
	cyclesLeft int
	f          func(*CPU)
}

func (o *Operation) tick(cpu *CPU) bool {
	o.cyclesLeft -= 1
	if o.cyclesLeft <= 0 {
		// fmt.Printf("#%3d: executing '%v'\n", cpu.cycle, o.source)
		o.f(cpu)
		return true
	}
	return false
}

func NewOperation(cycles int, f func(*CPU)) Operation {
	return Operation{"", cycles, f}
}

func ParseInstruction(s string) (op Operation, err error) {
	op, err = parseInstruction(s)
	op.source = s
	return
}

func parseInstruction(s string) (Operation, error) {
	split := strings.Split(s, " ")

	switch split[0] {

	case "noop":
		return noop(), nil

	case "addx":
		v, err := strconv.Atoi(split[1])
		if err != nil {
			return noop(), err
		}
		return addx(v), nil

	}

	return noop(), fmt.Errorf("operation %v not implemented", split[0])
}

////////////////////////////////////////////////////////////////////////////////

func addx(v int) Operation {
	return NewOperation(2, func(cpu *CPU) {
		cpu.x += v
	})
}

func noop() Operation { return NewOperation(1, func(cpu *CPU) {}) }
