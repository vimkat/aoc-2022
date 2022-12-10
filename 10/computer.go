package main

import "fmt"

type Computer struct {
	cpu *CPU
	crt *CRT
}

func NewComputer() *Computer {
	c := Computer{
		NewCPU(),
		NewCRT(40, 6),
	}
	return &c
}

func (c *Computer) ReadInstruction(s string) error {
	op, err := ParseInstruction(s)
	if err != nil {
		return fmt.Errorf("Parse error: %v", err)
	}

	c.cpu.ScheduleOperation(op)
	return nil
}

func (c *Computer) Tick() {
	// Update crt
	c.crt.SetSpritePos(c.cpu.X())
	c.crt.Tick()

	// Next CPU tick
	c.cpu.Tick()
}

func (c *Computer) RunToHalt() {
	for !c.cpu.IsHalted() {
		c.Tick()
	}
}

func (c *Computer) Show() {
	c.crt.Display()
}
