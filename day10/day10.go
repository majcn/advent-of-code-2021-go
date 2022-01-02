package main

import (
	. "../util"
	"fmt"
	"sort"
	"strings"
)

type DataType [][]byte

func parseData(data string) DataType {
	dataSplit := strings.Split(data, "\n")

	result := make(DataType, len(dataSplit))
	for i, line := range dataSplit {
		result[i] = []byte(line)
	}

	return result
}

func invert(el byte) byte {
	switch el {
	case ']': return '['
	case ')': return '('
	case '}': return '{'
	case '>': return '<'
	}

	return 0
}

func solvePart1(data DataType) (rc int) {
	score := func(el byte) int {
		switch el {
		case ')': return 3
		case ']': return 57
		case '}': return 1197
		case '>': return 25137
		}

		return 0
	}

	for _, line := range data {
		stack := make([]byte, 0, len(line))
		for _, el := range line {
			if el == '[' || el == '(' || el == '{' || el == '<' {
				stack = append(stack, el)
			} else if invert(el) == stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			} else {
				rc += score(el)
				break
			}
		}
	}

	return
}

func solvePart2(data DataType) (rc int) {
	score := func(el byte) int {
		switch el {
		case '(': return 1
		case '[': return 2
		case '{': return 3
		case '<': return 4
		}

		return 0
	}

	result := make([]int, 0, len(data))
	for _, line := range data {
		resultPart := 0
		corrupted := false
		stack := make([]byte, 0, len(line))
		for _, el := range line {
			if el == '[' || el == '(' || el == '{' || el == '<' {
				stack = append(stack, el)
			} else if invert(el) == stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			} else {
				corrupted = true
				break
			}
		}

		if !corrupted && len(stack) > 0 {
			for i := len(stack) - 1; i >= 0; i-- {
				resultPart = resultPart * 5 + score(stack[i])
			}
			if resultPart > 0 {
				result = append(result, resultPart)
			}
		}
	}

	sort.Ints(result)

	return result[len(result)/2]
}

func main() {
	data := parseData(FetchInputData(10))
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
