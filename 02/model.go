package main

import "fmt"

type Outcome int

func (o Outcome) String() string {
	switch o {
	case Loose:
		return "Loose"
	case Draw:
		return "Draw"
	case Win:
		return "Win"
	}
	panic(fmt.Sprintf("%d is not a valid outcome", o))
}

const (
	Loose Outcome = 0
	Draw  Outcome = 3
	Win   Outcome = 6
)

type RPS int

func (rps RPS) String() string {
	switch rps {
	case Rock:
		return "Rock"
	case Paper:
		return "Paper"
	case Scissors:
		return "Scissors"
	}

	panic(fmt.Sprintf("%d is not a valid RPS value", rps))
}

const (
	Rock RPS = iota + 1
	Paper
	Scissors
)

func parseRPSPlay(rps rune) RPS {
	return RPS(rps-'A') + 1
}

func parseRPSResponse(rps rune) RPS {
	return RPS(rps-'X') + 1
}

func parseOutcome(outcome rune) Outcome {
	switch outcome {
	case 'X':
		return Loose
	case 'Y':
		return Draw
	case 'Z':
		return Win
	}

	panic(fmt.Sprintf("%v is not a parsable outcome", outcome))
}

func getOutcome(myPlay, theirPlay RPS) Outcome {
	switch (myPlay - theirPlay + 3) % 3 {
	case 0:
		return Draw
	case 1:
		return Win
	case 2:
		return Loose
	}

	panic("unreachable")
}

func getPoints(opponentPlay, reaction RPS) int {
	return int(reaction) + int(getOutcome(reaction, opponentPlay))
}

func getChoice(opponentPlay RPS, outcome Outcome) RPS {
	offset := int(outcome / 3)
	return RPS((int(opponentPlay)+offset+1)%3) + 1
}
