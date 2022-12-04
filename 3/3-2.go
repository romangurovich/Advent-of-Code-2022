package main

import (
  common "Advent-of-Code-2022/Common"
  "fmt"
)

func main() {
  data := common.GetData(3)
  var chunkedData [][]string

  for i := 2; i < len(data); i += 3 {
    chunk := []string{data[i-2], data[i-1], data[i]}
    chunkedData = append(chunkedData, chunk)
  }

  var priorities = map[rune]int{}
  var lowercaseLetters = "abcdefghijklmnopqrstuvwxyz"
  var uppercaseLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
  for _, r := range lowercaseLetters {
    priorities[r] = int(r) - 96
  }
  for _, r := range uppercaseLetters {
    priorities[r] = int(r) - 38
  }
  var prioritiesSum = 0

  for _, group := range chunkedData {
    itemsSoFar := map[rune][3]bool{}
    for i, sack := range group {
      for _, item := range sack {
        presentAtSackIndex := itemsSoFar[item]
        presentAtSackIndex[i] = true
        itemsSoFar[item] = presentAtSackIndex
      }
    }

    var badge rune
    for k, v := range itemsSoFar {
      if v == [3]bool{true, true, true} {
        badge = k
        fmt.Println(group)
        fmt.Println(string(badge))
        break
      }
    }

    prioritiesSum += priorities[badge]
  }

  fmt.Println(prioritiesSum)
}
