package main

import (
	. "../util"
	"fmt"
	"strings"
)

type DataType []int

func parseData(data string) DataType {
	dataSplit := strings.Split(data, "\n")

	result := make(DataType, len(dataSplit))
	for i, v := range dataSplit {
		result[i] = ParseInt(v)
	}

	return result
}

func solvePart1(data DataType) (rc int) {
	for i := 1; i < len(data); i++ {
		if data[i] > data[i-1] {
			rc++
		}
	}
	return
}

func solvePart2(data DataType) (rc int) {
	prev := 0
	for i := 0; i < len(data) - 3; i++ {
		cur := Sum(data[i:i+3])
		if cur > prev {
			rc++
		}
		prev = cur
	}
	return
}

func main() {
	data := parseData(FetchInputData(1))
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
