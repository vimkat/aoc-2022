package main

import "github.com/vimkat/aoc-2022/util"

type Item rune

func (i Item) Priority() int {
	// Lowercsae -> Priority 1 - 26
	if i >= 'a' && i <= 'z' {
		return int(i-'a') + 1
	}

	// Uppercase -> Priority 27 - 52
	if i >= 'A' && i <= 'Z' {
		return int(i-'A') + 27
	}

	return -1
}

func (i Item) String() string {
	return string(i)
}

type Rucksack []Item

func (r Rucksack) String() string {
	return string(r.Compartment1()) + " | " + string(r.Compartment2())
}

func (r Rucksack) Compartment1() []Item {
	return r[:len(r)/2]
}

func (r Rucksack) Compartment2() []Item {
	return r[len(r)/2:]
}

func (r Rucksack) Valid() bool {
	return len(r)%2 == 0
}

func FindDuplicate(rucksack Rucksack) (Item, bool) {
	return util.FindDupliacte(rucksack.Compartment1(), rucksack.Compartment2())
}

type Group [3]Rucksack

func (g Group) Badge() (Item, bool) {
	return util.FindDupliacte(g[0], g[1], g[2])
}
