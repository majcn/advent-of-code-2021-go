package main

import (
	. "../util"
	"fmt"
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

func solvePartX(data DataType, resultFunc func(int)int) (rc int) {
	rc = MaxInt

	start, end := Min(data...), Max(data...)
	for i := start; i <= end; i++ {
		s := 0
		for _, d := range data {
			s += resultFunc(Abs(d-i))
		}

		rc = Min(rc, s)
	}
	return
}

func solvePart1(data DataType) (rc int) {
	return solvePartX(data, func(x int)int {
		return x
	})
}

func solvePart2(data DataType) (rc int) {
	return solvePartX(data, func(x int)int {
		return x * (x + 1) / 2
	})
}

func main() {
	data := parseData(FetchInputData(7))
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
