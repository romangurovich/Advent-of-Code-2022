package main

import (
	common "Advent-of-Code-2022/Common"
	"fmt"
)

func main() {
	data := common.GetData(6)[0]

	index := 0
	var discreteChars []uint8

	for index < len(data) {
		for contains(discreteChars, data[index]) {
			discreteChars = discreteChars[1:]
		}
		discreteChars = append(discreteChars, data[index])
		index += 1
		if len(discreteChars) == 4 {
			fmt.Println(index)
			break
		}
	}
}

func contains(s []uint8, search uint8) bool {
	for _, val := range s {
		if val == search {
			return true
		}
	}

	return false
}
