package main

import "fmt"

type HeightMap [][]Tree

func (m HeightMap) PrettyPrint() {
	for i := 0; i < m.Height(); i++ {
		for j := 0; j < m.Width(); j++ {
			fmt.Print(m[i][j])
		}
		fmt.Println()
	}
}

func (m HeightMap) Height() int {
	return len(m)
}

func (m HeightMap) Width() int {
	return len(m[0])
}

func (m *HeightMap) AddRow(row []Tree) {
	*m = append(*m, row)
}

func (m HeightMap) Row(row int) []Tree {
	return m[row]
}

func (m HeightMap) Column(column int) []Tree {
	col := make([]Tree, m.Height())
	for i, row := range m {
		col[i] = row[column]
	}
	return col
}

func (m HeightMap) Get(row, col int) Tree {
	return m[row][col]
}

func (m HeightMap) VisibleTrees() []Tree {
	visible := make([]Tree, 0, m.Height()*2+m.Width()*2)

	// First and last row are visible
	visible = append(visible, m.Row(0)...)
	visible = append(visible, m.Row(m.Height()-1)...)

	// Same for column (however the first and last tree have already been added)
	visible = append(visible, m.Column(0)[1:m.Height()-1]...)
	visible = append(visible, m.Column(m.Width() - 1)[1:m.Height()-1]...)

	// Check inside trees
	for row := 1; row < m.Height()-1; row++ {
		for col := 1; col < m.Width()-1; col++ {
			if m.IsVisible(row, col) {
				visible = append(visible, m[row][col])
			}
		}
	}

	return visible
}

func (m HeightMap) IsVisible(row, col int) bool {
	// Outside trees are visible
	if col == 0 || row == 0 || col == m.Width()-1 || row == m.Height()-1 {
		return true
	}

	tree := m[row][col]

	// Left trees
	if height, _ := maxTree(m.Row(row)[:col]); height < tree {
		return true
	}

	// Right trees
	if height, _ := maxTree(m.Row(row)[col+1:]); height < tree {
		return true
	}

	// Top trees
	if height, _ := maxTree(m.Column(col)[:row]); height < tree {
		return true
	}

	// Bottom trees
	if height, _ := maxTree(m.Column(col)[row+1:]); height < tree {
		return true
	}

	return false
}

func (m HeightMap) ScenicScore(row, col int) int {
	return m.viewingDistance(row, col, 1, 0) *
		m.viewingDistance(row, col, 0, 1) *
		m.viewingDistance(row, col, -1, 0) *
		m.viewingDistance(row, col, 0, -1)
}

func maxTree(trees []Tree) (Tree, int) {
	max := trees[0]
	for _, tree := range trees {
		if tree > max {
			max = tree
		}
	}
	return max, 0
}

func (m HeightMap) viewingDistance(row, col, walkX, walkY int) int {
	// Trees on the edge have viewing distance 0 when going outside the grid
	if row == 0 && walkY == -1 {
		return 0
	}
	if row == m.Height()-1 && walkY == 1 {
		return 0
	}
	if col == 0 && walkX == -1 {
		return 0
	}
	if col == m.Width()-1 && walkX == 1 {
		return 0
	}

	tree := m[row][col]
	distance := 0

	// walk along x
	if walkX != 0 {
		for i := col + walkX; i < m.Width() && i >= 0; i += walkX {
			distance += 1
			if m[row][i] >= tree {
				return distance
			}
		}
	}

	if walkY != 0 {
		for i := row + walkY; i < m.Height() && i >= 0; i += walkY {
			distance += 1
			if m[i][col] >= tree {
				return distance
			}
		}
	}

	return distance
}

type Tree int
