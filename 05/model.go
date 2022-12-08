package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Crate rune

func (c Crate) String() string {
	return fmt.Sprintf("[%v]", string(c))
}

func ParseCrate(s string) (Crate, error) {
	if len(s) != 3 {
		return 0, fmt.Errorf("a crate is identified by a single letter, %v given", s)
	}

	crate := rune(s[1])
	return Crate(crate), nil
}

type Stack []Crate

func (s Stack) String() string {
	builder := strings.Builder{}
	for _, crate := range s {
		builder.WriteString(crate.String())
	}
	return builder.String()
}

func (s Stack) Equals(other Stack) bool {
	if len(s) != len(other) {
		return false
	}
	for i, e := range s {
		if e != other[i] {
			return false
		}
	}
	return true
}

func (s *Stack) Push(crate Crate) {
	(*s) = append(*s, crate)
}

func (s *Stack) PushN(crates []Crate) {
	(*s) = append(*s, crates...)
}

func (s Stack) Peek() Crate {
	return s[len(s)-1]
}

func (s *Stack) Pop() Crate {
	topIndex := len(*s) - 1
	top := (*s)[topIndex]
	*s = (*s)[:topIndex]
	return top
}

func (s *Stack) PopN(n int) []Crate {
	crates := (*s)[len(*s)-n:]
	*s = (*s)[:len(*s)-n]
	return crates
}

type Ship []Stack

func (s Ship) String() string {
	builder := strings.Builder{}
	for i, st := range s {
		builder.WriteString(fmt.Sprintf("%d:\t%v\n", i+1, st))
	}
	return builder.String()
}

func (s *Ship) Execute9000(op MoveOp) {
	for i := 0; i < op.count; i++ {
		crate := (*s)[op.fromStack].Pop()
		(*s)[op.toStack].Push(crate)
	}
}

func (s *Ship) Execute9001(op MoveOp) {
	crates := (*s)[op.fromStack].PopN(op.count)
	(*s)[op.toStack].PushN(crates)
}

func (s *Ship) ExecuteAll(ops []MoveOp) {
	for _, op := range ops {
		s.Execute9000(op)
	}
}

func ParseShip(s string) (Ship, error) {
	lines := strings.Split(s, "\n")
	return ParseShipByLines(lines)
}

func ParseShipByLines(lines []string) (Ship, error) {
	split := (strings.Split(strings.TrimSpace(lines[len(lines)-1]), " "))
	numberOfStacks, err := strconv.Atoi(split[len(split)-1])
	if err != nil {
		return nil, fmt.Errorf("couldn't parse index line: %v", err)
	}

	s := make([]Stack, numberOfStacks)

	// start at the second last line since the last line has been parsed above
	for i := len(lines) - 2; i >= 0; i-- {
		byStack := cratesByStack(lines[i])
		for j, cStr := range byStack {
			// Ignore stacks that have been filled
			if cStr == "" {
				continue
			}

			// Stack new crate
			crate, err := ParseCrate(cStr)
			if err != nil {
				return nil, fmt.Errorf("couldn't parse ship line %v: %v", len(lines)-1, err)
			}
			s[j].Push(crate)
		}
	}

	return s, nil
}

func cratesByStack(s string) []string {
	split := []string{}
	for i := 0; i < len(s); i += 4 {
		if strings.HasPrefix(s[i:], "   ") {
			// empty space
			split = append(split, "")
		} else {
			// crate
			split = append(split, s[i:i+3])
		}
	}
	return split
}

type MoveOp struct {
	count     int
	fromStack int
	toStack   int
}

func (m MoveOp) String() string {
	return fmt.Sprintf("move %d from %d to %d", m.count, m.fromStack, m.toStack)
}

func (m MoveOp) Equals(other MoveOp) bool {
	return m.count == other.count &&
		m.fromStack == other.fromStack &&
		m.toStack == other.toStack
}

func ParseMoveOp(str string) (MoveOp, error) {
	var amount, from, to string
	p := NewParser(str)
	p.
		Consume("move ").
		Capture(&amount).
		Consume(" from ").
		Capture(&from).
		Consume(" to ").
		Capture(&to)

	if p.Err() != nil {
		return MoveOp{}, fmt.Errorf("couldn't parse %v: %v", str, p.Err())
	}

	amountI, err := strconv.Atoi(amount)
	if err != nil {
		return MoveOp{}, fmt.Errorf("couldn't parse AMOUNT, %v given", amount)
	}
	fromI, err := strconv.Atoi(from)
	if err != nil {
		return MoveOp{}, fmt.Errorf("couldn't parse FROM, %v given", from)
	}
	toI, err := strconv.Atoi(to)
	if err != nil {
		return MoveOp{}, fmt.Errorf("couldn't parse TO %v given", to)
	}

	return MoveOp{amountI, fromI - 1, toI - 1}, nil
}

type Parser struct {
	str    string
	cursor int
	err    error
}

func NewParser(s string) *Parser {
	return &Parser{s, 0, nil}
}

func (p *Parser) Err() error {
	return p.err
}

func (p *Parser) Input() string {
	return p.str[p.cursor:]
}

func (p *Parser) Consume(s string) *Parser {
	// Nop on error
	if p.err != nil {
		return p
	}

	if strings.HasPrefix(p.Input(), s) {
		p.cursor += len(s)
	} else {
		p.err = fmt.Errorf("can't consume %v", s)
	}
	return p
}

func (p *Parser) Capture(dest *string) *Parser {
	// Nop on error
	if p.err != nil {
		return p
	}

	indexOfSpace := 0
	for _, c := range p.Input() {
		if c == ' ' {
			break
		}
		indexOfSpace += 1
	}

	*dest = p.Input()[:indexOfSpace]
	p.cursor += indexOfSpace
	return p
}
