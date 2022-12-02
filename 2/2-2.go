package main

import (
	common "Advent-of-Code-2022/Common"
	"fmt"
	"strings"
)

func main() {
	const ROCK = "rock"
	const PAPER = "paper"
	const SCISSORS = "scissors"

	const LOSE = 0
	const DRAW = 3
	const WIN = 6

	var hands = map[string]string{
		"A": ROCK,
		"B": PAPER,
		"C": SCISSORS,
	}

	var outcomes = map[string]int{
		"X": LOSE,
		"Y": DRAW,
		"Z": WIN,
	}

	var handValues = map[string]int{
		ROCK:     1,
		PAPER:    2,
		SCISSORS: 3,
	}

	var dominates = map[string]string{
		ROCK:     SCISSORS,
		SCISSORS: PAPER,
		PAPER:    ROCK,
	}

	var dominatedBy = map[string]string{
		ROCK:     PAPER,
		PAPER:    SCISSORS,
		SCISSORS: ROCK,
	}

	var data = common.GetData(2)
	var total = 0
	for _, round := range data {
		var roundValues = strings.Split(round, " ")
		var theirHand = hands[roundValues[0]]
		var myOutcome = outcomes[roundValues[1]]

		var myHand = SCISSORS
		total += myOutcome

		switch myOutcome {
		case LOSE:
			myHand = dominates[theirHand]
		case DRAW:
			myHand = theirHand
		case WIN:
			myHand = dominatedBy[theirHand]
		}
		total += handValues[myHand]
	}

	fmt.Println(total)
}
