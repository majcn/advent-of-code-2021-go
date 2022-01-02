package main

import (
	"fmt"
	. "majcn.si/advent-of-code-2021/util"
	"strconv"
	"strings"
)

type DataType []string

func parseData(data string) DataType {
	dataSplit := strings.Split(data, "\n")

	result := make(DataType, len(dataSplit))
	for i, line := range dataSplit {
		result[i] = line
	}

	return result
}

func solvePart1(data DataType) int {
	gamma, epsilon := "", ""

	for x := 0; x < len(data[0]); x++ {
		ones := 0
		zeros := 0
		for y := 0; y < len(data); y++ {
			if data[y][x] == '1' {
				ones++
			} else {
				zeros++
			}
		}

		if ones >= zeros {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}

	gammaInt, _ := strconv.ParseInt(gamma, 2, 64)
	epsilonInt, _ := strconv.ParseInt(epsilon, 2, 64)

	return int(gammaInt * epsilonInt)
}


func getElement(data DataType, predicate func(int, int, byte)bool) string {
	element := data
	for x := 0; x < len(element[0]); x++ {
		ones := 0
		zeros := 0
		for y := 0; y < len(element); y++ {
			if element[y][x] == '1' {
				ones++
			} else {
				zeros++
			}
		}

		nextElement := make([]string, 0)
		for y := 0; y < len(element); y++ {
			if predicate(ones, zeros, element[y][x]) {
				nextElement = append(nextElement, element[y])
			}
		}
		element = nextElement

		if len(element) == 1 {
			return element[0]
		}
	}

	return ""
}

func solvePart2(data DataType) int {
	oxygen := getElement(data, func(ones int, zeros int, x byte) bool {
		return (ones >= zeros && x == '1') || (ones < zeros && x == '0')
	})

	co2 := getElement(data, func(ones int, zeros int, x byte) bool {
		return (ones >= zeros && x == '0') || (ones < zeros && x == '1')
	})

	oxygenInt, _ := strconv.ParseInt(oxygen, 2, 64)
	co2Int, _ := strconv.ParseInt(co2, 2, 64)

	return int(oxygenInt * co2Int)
}

func main() {
	data := parseData(FetchInputData(3))
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
