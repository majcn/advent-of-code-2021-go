package main

import (
	"fmt"
	. "majcn.si/advent-of-code-2021/util"
	"strings"
)

type DataType []int

func parseData(data string) DataType {
	dataSplit := strings.Split(data, "\n")

	result := make([]int, len(dataSplit))
	for i, line := range dataSplit {
		result[i] = ParseInt(line)
	}

	return result
}

func solvePart1(data DataType) (rc int) {
	// done on piece of paper
	return 18051
}

func solvePart2(data DataType) (rc int) {
	// done on piece of paper
	return 50245
}

func main() {
	data := parseData(FetchInputData(23))
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
