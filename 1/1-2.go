package main

import (
	common "Advent-of-Code-2022/Common"
	"fmt"
	"strconv"
)

func main() {
	var data = common.GetData(1)
	currentCalories := 0
	top3CalorieTotals := []int{0, 0, 0}
	sumTop3Calories := 0

	for _, calorieCount := range data {
		if calorieCount == "" {
			updateTop3CalorieTotals(&top3CalorieTotals, currentCalories)
			currentCalories = 0
		}
		calories, _ := strconv.Atoi(calorieCount)
		currentCalories += calories
	}

	updateTop3CalorieTotals(&top3CalorieTotals, currentCalories)

	for _, calories := range top3CalorieTotals {
		sumTop3Calories += calories
	}

	fmt.Println(sumTop3Calories)
}

func updateTop3CalorieTotals(top3CalorieTotalsRef *[]int, calories int) {
	top3CalorieTotals := *top3CalorieTotalsRef
	if calories > top3CalorieTotals[0] {
		*top3CalorieTotalsRef = append([]int{calories}, top3CalorieTotals[:2]...)
	} else if calories > top3CalorieTotals[1] {
		*top3CalorieTotalsRef = append(top3CalorieTotals[:1], calories, top3CalorieTotals[1])
	} else if calories > top3CalorieTotals[2] {
		*top3CalorieTotalsRef = append(top3CalorieTotals[0:2], calories)
	}
}
