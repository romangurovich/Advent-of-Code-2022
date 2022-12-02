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

  var hands = map[string]string{
    "A": ROCK,
    "B": PAPER,
    "C": SCISSORS,
    "X": ROCK,
    "Y": PAPER,
    "Z": SCISSORS,
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

  var data = common.GetData(2)
  var total = 0
  for _, round := range data {
    var roundHands = strings.Split(round, " ")
    var theirHand = hands[roundHands[0]]
    var myHand = hands[roundHands[1]]

    total += handValues[myHand]

    if dominates[myHand] == theirHand {
      total += 6
    } else if myHand == theirHand {
      total += 3
    }
  }
  fmt.Println(total)
}
