package main

import (
	common "Advent-of-Code-2022/Common"
	"fmt"
	"strconv"
)

func main() {
	var data = common.GetData(1)
	maxCalories := 0
	currentCalories := 0

	for _, calorieCount := range data {
		if calorieCount == "" {
			maxCalories = common.Max(maxCalories, currentCalories)
			currentCalories = 0
		}
		calories, _ := strconv.Atoi(calorieCount)
		currentCalories += calories
	}

	maxCalories = common.Max(maxCalories, currentCalories)
	fmt.Println(maxCalories)
}
