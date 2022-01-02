package main

import (
	"fmt"
	. "majcn.si/advent-of-code-2021/util"
	"strings"
)

type DataType []struct {
	input [10]StringSet
	output []StringSet
}

func parseData(data string) DataType {
	dataSplit := strings.Split(data, "\n")

	result := make(DataType, len(dataSplit))
	for i, line := range dataSplit {
		v := strings.Split(line, "|")
		inputV, outputV := strings.Fields(v[0]), strings.Fields(v[1])

		for j, cs := range inputV {
			result[i].input[j] = NewStringSet(strings.Split(cs, ""))
		}

		result[i].output = make([]StringSet, len(outputV))
		for j, cs := range outputV {
			result[i].output[j] = NewStringSet(strings.Split(cs, ""))
		}
	}

	return result
}

func next(data [10]StringSet, predicate func(StringSet)bool) StringSet {
	for _, x := range data {
		if predicate(x) {
			return x
		}
	}

	return nil
}

func solvePartX(data DataType) [][10]StringSet {
	result := make([][10]StringSet, len(data))

	for i, dataEl := range data {
		numbers := [10]StringSet{}

		numbers[1] = next(dataEl.input, func(x StringSet) bool {
			return len(x) == 2
		})
		numbers[7] = next(dataEl.input, func(x StringSet) bool {
			return len(x) == 3
		})
		numbers[4] = next(dataEl.input, func(x StringSet) bool {
			return len(x) == 4
		})
		numbers[8] = next(dataEl.input, func(x StringSet) bool {
			return len(x) == 7
		})

		numbers[9] = next(dataEl.input, func(x StringSet) bool {
			return len(x) == 6 && numbers[4].IsSubset(&x)
		})
		numbers[0] = next(dataEl.input, func(x StringSet) bool {
			return len(x) == 6 && numbers[1].IsSubset(&x) && !x.Equals(&numbers[9])
		})
		numbers[6] = next(dataEl.input, func(x StringSet) bool {
			return len(x) == 6 && !x.Equals(&numbers[9]) && !x.Equals(&numbers[0])
		})

		numbers[5] = next(dataEl.input, func(x StringSet) bool {
			return len(x) == 5 && x.IsSubset(&numbers[6])
		})
		numbers[3] = next(dataEl.input, func(x StringSet) bool {
			return len(x) == 5 && x.IsSubset(&numbers[9]) && !x.Equals(&numbers[5])
		})
		numbers[2] = next(dataEl.input, func(x StringSet) bool {
			return len(x) == 5 && !x.Equals(&numbers[5]) && !x.Equals(&numbers[3])
		})

		result[i] = numbers
	}

	return result
}

func GetIndex(slice [10]StringSet, element StringSet) int {
	for i, s := range slice {
		if s.Equals(&element) {
			return i
		}
	}
	return -1
}

func solvePart1(data DataType) (rc int) {
	numbers := solvePartX(data)

	for i, dataEl := range data {
		for _, x := range dataEl.output {
			n := GetIndex(numbers[i], x)
			if n == 1 || n == 4 || n == 7 || n == 8 {
				rc++
			}
		}
	}

	return
}

func solvePart2(data DataType) (rc int) {
	numbers := solvePartX(data)

	for i, dataEl := range data {
		outputNumber := make([]int, len(dataEl.output))
		for j, v := range dataEl.output {
			outputNumber[j] = GetIndex(numbers[i], v)
		}
		rc += ParseInt(outputNumber)
	}

	return
}

func main() {
	data := parseData(FetchInputData(8))
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
