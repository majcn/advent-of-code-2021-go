package main

import (
	"fmt"
	. "majcn.si/advent-of-code-2021/util"
	"strings"
)

type DataType []int

func parseData(data string) DataType {
	dataSplit := strings.Split(data, ",")

	result := make(DataType, len(dataSplit))
	for i, v := range dataSplit {
		result[i] = ParseInt(v)
	}

	return result
}

func solvePartX(data DataType, timeLimit int) int {
	result := [9]int{}
	for _, d := range data {
		result[d]++
	}

	for t := 0; t < timeLimit; t++ {
		zeros := result[0]
		result[0] = 0

		for i := 1; i < 9; i++ {
			result[i-1] += result[i]
			result[i] = 0
		}

		result[6] += zeros
		result[8] += zeros
	}

	return Sum(result[:])
}

func solvePart1(data DataType) (rc int) {
	return solvePartX(data, 80)
}

func solvePart2(data DataType) (rc int) {
	return solvePartX(data, 256)
}

func main() {
	data := parseData(FetchInputData(6))
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
