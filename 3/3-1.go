package main

import (
	common "Advent-of-Code-2022/Common"
	"fmt"
)

func main() {
	data := common.GetData(3)

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

	for _, sack := range data {
		itemsSoFar := map[rune]bool{}
		half := len(sack) / 2
		compartment1 := sack[:half]
		compartment2 := sack[half:]

		for _, item := range compartment1 {
			itemsSoFar[item] = true
		}

		var duplicateItem rune
		for _, item := range compartment2 {
			if itemsSoFar[item] {
				duplicateItem = item
				break
			}
		}

		prioritiesSum += priorities[duplicateItem]
	}

	fmt.Println(prioritiesSum)
}
