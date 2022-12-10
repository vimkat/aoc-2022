package main

import (
	"fmt"

	"github.com/vimkat/aoc-2022/util"
)

type Pos struct{ x, y int }

func (p Pos) String() string   { return fmt.Sprintf("(%d/%d)", p.x, p.y) }
func (p Pos) Add(x, y int) Pos { return Pos{p.x + x, p.y + y} }

func (p Pos) IsAdjacent(other Pos) bool {
	return util.Abs(p.x-other.x) <= 1 && util.Abs(p.y-other.y) <= 1
}

func (p Pos) MakeAdjacentTo(other Pos) Pos {
	if p.IsAdjacent(other) {
		return p
	}

	if p.x == other.x {
		return Pos{other.x, other.y - util.Sign(other.y-p.y)}
	}

	if p.y == other.y {
		return Pos{other.x - util.Sign(other.x-p.x), other.y}
	}

	dx := util.Abs(other.x - p.x)
	dy := util.Abs(other.y - p.y)

	if dx < dy {
		return Pos{other.x, other.y - util.Sign(other.y-p.y)}
	}

	if dy < dx {
		return Pos{other.x - util.Sign(other.x-p.x), other.y}
	}

	return Pos{
		other.x - util.Sign(other.x-p.x),
		other.y - util.Sign(other.y-p.y),
	}
}

type Rope struct {
	rope    []Pos
	visited map[Pos]int
}

func NewRope(ropeLength int) *Rope {
	r := Rope{
		rope:    make([]Pos, ropeLength),
		visited: make(map[Pos]int),
	}

	// Start position has already been visited
	r.visited[Pos{0, 0}] = 1

	return &r
}

func (r *Rope) Move(dx, dy int) {
	// Move the head
	r.rope[0] = r.rope[0].Add(dx, dy)

	// Update the tail(s)
	for i := 1; i < len(r.rope); i++ {
		r.rope[i] = r.rope[i].MakeAdjacentTo(r.rope[i-1])
	}

	// Mark visited
	val := r.visited[r.rope[len(r.rope)-1]]
	val += 1
	r.visited[r.rope[len(r.rope)-1]] = val
}

func (r Rope) Visited() map[Pos]int {
	return r.visited
}

func (r Rope) PrettyPrint(width, height, centerX, centerY int) {
	grid := make([][]rune, height)
	for row := 0; row < height; row++ {
		grid[row] = make([]rune, width)
		for col := 0; col < width; col++ {
			grid[row][col] = '.'
		}
	}

	grid[centerY][centerX] = 'S'

	// Print knots
	for i, knot := range r.rope {
		grid[knot.y+centerY][knot.x+centerX] = rune(i + '0')
	}

	for row := 0; row < height; row++ {
		fmt.Println(string(grid[row]))
	}
}
