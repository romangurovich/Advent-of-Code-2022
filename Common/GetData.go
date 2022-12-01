package common

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func GetData(num int) []string {
	file, err := os.Open(fmt.Sprintf("./Data/%s.txt", strconv.Itoa(num)))
	var result []string
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}
