package main

import "testing"

func TestParseRange(t *testing.T) {
	r, err := ParseRange("2-4")
	if err != nil {
		t.Fatalf("error parsing range: %v", err)
	}
	if r.start != 2 {
		t.Fatalf("couldn't parse range start: got %v, expected 2", r.start)
	}
	if r.end != 4 {
		t.Fatalf("couldn't parse range end: got %v, expected 2", r.start)
	}
}

func TestParsePair(t *testing.T) {
	p, err := ParsePair("2-4,6-8")
	if err != nil {
		t.Fatalf("error parsing pair: %v", err)
	}
	if !p.elf1.Equals(Range{2, 4}) {
		t.Fatalf("couldn't parse first range: got %v, expected 2-4", p.elf1)
	}
	if !p.elf2.Equals(Range{6, 8}) {
		t.Fatalf("couldn't parse second range: got %v, expected 6-8", p.elf1)
	}
}

func TestRangeFullyOverlap(t *testing.T) {
	cases := map[Pair]bool{
		{Range{2, 4}, Range{6, 8}}: false,
		{Range{2, 3}, Range{4, 5}}: false,
		{Range{5, 7}, Range{7, 9}}: false,
		{Range{2, 8}, Range{3, 7}}: true,
		{Range{6, 6}, Range{4, 6}}: true,
		{Range{2, 6}, Range{4, 8}}: false,
	}

	for pair, expected := range cases {
		t.Run(pair.String(), func(t *testing.T) {
			hasOverlap := pair.FullyOverlap()
			if hasOverlap != expected {
				t.Fatalf("%s full overlap resulted in %v, expected %v", pair, hasOverlap, expected)
			}
		})
	}
}

func TestRangeOverlaps(t *testing.T) {
	cases := map[Pair]bool{
		{Range{2, 4}, Range{6, 8}}: false,
		{Range{2, 3}, Range{4, 5}}: false,
		{Range{5, 7}, Range{7, 9}}: true,
		{Range{2, 8}, Range{3, 7}}: true,
		{Range{6, 6}, Range{4, 6}}: true,
		{Range{2, 6}, Range{4, 8}}: true,
	}

	for pair, expected := range cases {
		t.Run(pair.String(), func(t *testing.T) {
			hasOverlap := pair.PartiallyOverlap()
			if hasOverlap != expected {
				t.Fatalf("%s partial overlap resulted in %v, expected %v", pair, hasOverlap, expected)
			}
		})
	}
}
