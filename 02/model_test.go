package main

import (
	"fmt"
	"testing"
)

type round struct{ play, response RPS }

func (r round) String() string {
	return fmt.Sprintf("%v vs. %v", r.play, r.response)
}

type strategy struct {
	play    RPS
	outcome Outcome
}

func TestParseRPSPlay(t *testing.T) {
	cases := map[rune]RPS{
		'A': Rock,
		'B': Paper,
		'C': Scissors,
	}

	for in, out := range cases {
		t.Run(string(in), func(t *testing.T) {
			res := parseRPSPlay(in)
			if res != out {
				t.Fatalf("parsing %v resulted in %v, expected %v", in, res, out)
			}
		})
	}
}

func TestParseRPSResponse(t *testing.T) {
	cases := map[rune]RPS{
		'X': Rock,
		'Y': Paper,
		'Z': Scissors,
	}

	for in, out := range cases {
		t.Run(string(in), func(t *testing.T) {
			res := parseRPSResponse(in)
			if res != out {
				t.Fatalf("parsing %v resulted in %v, expected %v", in, res, out)
			}
		})
	}
}

func TestParseOutcome(t *testing.T) {
	cases := map[rune]Outcome{
		'X': Loose,
		'Y': Draw,
		'Z': Win,
	}

	for in, out := range cases {
		outcome := parseOutcome(in)
		if outcome != out {
			t.Fatalf("parsing %v resulted in %v, expected %v", in, outcome, out)
		}
	}
}

func TestGetOutcome(t *testing.T) {
	cases := map[round]Outcome{
		{Rock, Rock}:         Draw,
		{Rock, Paper}:        Loose,
		{Rock, Scissors}:     Win,
		{Paper, Rock}:        Win,
		{Paper, Paper}:       Draw,
		{Paper, Scissors}:    Loose,
		{Scissors, Rock}:     Loose,
		{Scissors, Paper}:    Win,
		{Scissors, Scissors}: Draw,
	}

	for round, expected := range cases {
		t.Run(round.String(), func(t *testing.T) {
			outcome := getOutcome(round.play, round.response)
			if outcome != expected {
				t.Fatalf("%v resulted in %v, expected %v", round, outcome, expected)
			}
		})
	}
}

func TestGetPoints(t *testing.T) {
	cases := map[round]int{
		{Rock, Paper}:        8,
		{Paper, Rock}:        1,
		{Scissors, Scissors}: 6,
	}

	for round, expected := range cases {
		t.Run(round.String(), func(t *testing.T) {
			points := getPoints(round.play, round.response)
			if points != expected {
				t.Fatalf("%v resulted in %v points, expected %v", round, points, expected)
			}
		})
	}
}

func TestGetChoice(t *testing.T) {
	cases := map[strategy]RPS{
		{Rock, Loose}:     Scissors,
		{Rock, Draw}:      Rock,
		{Rock, Win}:       Paper,
		{Paper, Loose}:    Rock,
		{Paper, Draw}:     Paper,
		{Paper, Win}:      Scissors,
		{Scissors, Loose}: Paper,
		{Scissors, Draw}:  Scissors,
		{Scissors, Win}:   Rock,
	}

	for strategy, expected := range cases {
		t.Run(fmt.Sprintf("%v_%v", strategy.play, strategy.outcome), func(t *testing.T) {
			choice := getChoice(strategy.play, strategy.outcome)
			if choice != expected {
				t.Fatalf("chose %v in response to %v in order to %v, expected %v",
					choice, strategy.play, strategy.outcome, expected)
			}
		})
	}
}
