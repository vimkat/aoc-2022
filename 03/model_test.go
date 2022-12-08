package main

import (
	"strconv"
	"testing"
)

var (
	validR    = Rucksack("jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL")
	invalidR  = Rucksack("A")
	rucksacks = []Rucksack{
		Rucksack("vJrwpWtwJgWrhcsFMMfFFhFp"),
		Rucksack("jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL"),
		Rucksack("PmmdzqPrVvPwwTWBwg"),
		Rucksack("wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn"),
		Rucksack("ttgJtRGJQctTZtZT"),
		Rucksack("CrZsJsPPZsGzwwsLwLmpwMDw"),
	}
	duplicates = []Item{'p', 'L', 'P', 'v', 't', 's'}
	priorities = []int{16, 38, 42, 22, 20, 19}
)

func TestRucksackString(t *testing.T) {
	if validR.String() != "jqHRNqRjqzjGDLGL | rsFMfFZSrLrFZsSL" {
		t.Fatal()
	}
}

func TestRucksackCompartments(t *testing.T) {
	total := string(validR.Compartment1()) + string(validR.Compartment2())
	if total != string(validR) {
		t.Fatalf("%v couldn't be split into compartmetns correctly", validR)
	}
}

func TestRucksackValid(t *testing.T) {
	if !validR.Valid() {
		t.Fatalf("%v is not a valid rucksack even tho it should be", validR)
	}
	if invalidR.Valid() {
		t.Fatalf("%v is a valid rucksack even tho it shouldn't be", validR)
	}
}

func TestFindDuplicates(t *testing.T) {
	for i, r := range rucksacks {
		t.Run(string(duplicates[i]), func(t *testing.T) {
			item, found := FindDuplicate(r)
			if !found {
				t.Fatalf("%v should contain the duplicate %v, none found", r, duplicates[i])
			}
			if item != duplicates[i] {
				t.Fatalf("%v should contain the duplicate %v, %v found", r, duplicates[i], item)
			}
		})
	}
}

func TestItemPriority(t *testing.T) {
	for i, item := range duplicates {
		t.Run(string(item), func(t *testing.T) {
			p := item.Priority()
			if p != priorities[i] {
				t.Fatalf("%v should have priority %v, %v caluclated", item, priorities[i], p)
			}
		})
	}
}

func TestGroupBadge(t *testing.T) {
	var groupA, groupB Group
	copy(groupA[:], rucksacks[:3])
	copy(groupB[:], rucksacks[3:])
	badges := []Item{'r', 'Z'}
	groups := []Group{groupA, groupB}

	for i, group := range groups {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b, _ := group.Badge()
			if b != badges[i] {
				t.Fatalf("%v has badge %v, expected %v", group, b, badges[i])
			}
		})
	}
}
