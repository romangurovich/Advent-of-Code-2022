package main

import (
	common "Advent-of-Code-2022/Common"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	data := common.GetData(4)
	numOverAssignments := 0
	for _, entry := range data {
		elves := strings.Split(entry, ",")
		elf1 := elves[0]
		elf2 := elves[1]
		assignment1 := createAssignment(elf1)
		assignment2 := createAssignment(elf2)
		if rangeOverlapsOther(assignment1, assignment2) {
			numOverAssignments += 1
		}
	}

	fmt.Println(numOverAssignments)
}

type Assignment struct {
	Min int
	Max int
}

func createAssignment(elf string) Assignment {
	slot := strings.Split(elf, "-")
	rangeStart, _ := strconv.Atoi(slot[0])
	rangeEnd, _ := strconv.Atoi(slot[1])
	return Assignment{
		Min: rangeStart,
		Max: rangeEnd,
	}
}

func rangeOverlapsOther(assignment Assignment, otherAssignment Assignment) bool {
	firstOverlapsSecond := assignment.Max >= otherAssignment.Min && assignment.Min <= otherAssignment.Min
	secondOverlapsFirst := otherAssignment.Max >= assignment.Min && otherAssignment.Min <= assignment.Min
	return firstOverlapsSecond || secondOverlapsFirst
}
