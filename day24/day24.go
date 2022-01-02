package main

import (
	. "../util"
	"fmt"
	"strings"
)

type DataType [][2]int

func parseData(data string) DataType {
	dataSplit := strings.Split(data, "inp w\n")[1:]

	result := make([][2]int, len(dataSplit))
	for i, lines := range dataSplit {
		linesSplit := strings.Split(lines, "\n")
		result[i][0] = ParseInt(strings.Split(linesSplit[4], " ")[2])
		result[i][1] = ParseInt(strings.Split(linesSplit[14], " ")[2])
	}

	return result
}

func program(data DataType, input [7]int) ([7]int, bool) {
	magicNumber := 26
	output := [7]int{}
	z := 0

	inputI, outputI := 0, 0

	for _, d := range data {
		x, y := d[0], d[1]

		if x > 0 {
			z = (z * magicNumber) + y + input[inputI]; inputI++
		} else {
			outputResult := (z % magicNumber) + x
			if outputResult < 1 || outputResult > 9 {
				return [7]int{}, false
			}

			z /= magicNumber
			output[outputI] = outputResult; outputI++
		}
	}

	return output, true
}

func generateModel(i [7]int, o [7]int) int {
	model := []int{i[0], i[1], i[2], i[3], i[4], o[0], i[5], o[1], o[2], i[6], o[3], o[4], o[5], o[6]}
	return ParseInt(model)
}

func solvePartX(data DataType, searchRange []int) (rc int) {
	for _, i1 := range searchRange {
		for _, i2 := range searchRange {
			for _, i3 := range searchRange {
				for _, i4 := range searchRange {
					for _, i5 := range searchRange {
						for _, i6 := range searchRange {
							for _, i7 := range searchRange {
								input := [7]int{i1, i2, i3, i4, i5, i6, i7}
								if output, successful := program(data, input); successful {
									return generateModel(input, output)
								}
							}
						}
					}
				}
			}
		}
	}
	return
}

func solvePart1(data DataType) (rc int) {
	searchRange := make([]int, 9)
	for i := 0; i < 9; i++ {
		searchRange[i] = 9 - i
	}

	return solvePartX(data, searchRange)
}

func solvePart2(data DataType) (rc int) {
	searchRange := make([]int, 9)
	for i := 0; i < 9; i++ {
		searchRange[i] = i + 1
	}

	return solvePartX(data, searchRange)
}

func main() {
	data := parseData(FetchInputData(24))
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
