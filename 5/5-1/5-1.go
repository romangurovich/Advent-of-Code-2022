package main

import (
	common "Advent-of-Code-2022/Common"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data := common.GetData(5)
	crates := generateCrates()

	for _, entry := range data[10:] {
		command := parseCommand(entry)
		fromColumn := crates[command.From]
		toColumn := crates[command.To]
		num := command.NumCrates

		for i := 0; i < num; i++ {
			toColumn = append(toColumn, fromColumn[len(fromColumn)-1])
			fromColumn = fromColumn[:len(fromColumn)-1]
			crates[command.To] = toColumn
			crates[command.From] = fromColumn
		}
	}

	fmt.Println(topCrates(crates))
}

type Command struct {
	NumCrates int
	From      int
	To        int
}

func topCrates(crates [][]string) string {
	var result strings.Builder
	for _, stack := range crates {
		result.WriteString(stack[len(stack)-1])
	}
	return result.String()
}

func parseCommand(entry string) Command {
	r, _ := regexp.Compile(`\d+`)
	matches := r.FindAllString(entry, -1)
	numCrates, _ := strconv.Atoi(matches[0])
	from, _ := strconv.Atoi(matches[1])
	to, _ := strconv.Atoi(matches[2])
	return Command{numCrates, from - 1, to - 1}
}

func generateCrates() [][]string {
	data := common.GetData(5)
	var allCrates = make([][]string, 9)

	for i := 0; i < 8; i++ {
		row := data[i]
		stackIndex := 0
		for j := 1; j < len(row); j += 4 {
			val := string(row[j])
			if val != " " {
				allCrates[stackIndex] = append([]string{val}, allCrates[stackIndex]...)
			}
			stackIndex += 1
		}
	}

	return allCrates
}
