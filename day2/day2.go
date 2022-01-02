package main

import (
	"fmt"
	. "majcn.si/advent-of-code-2021/util"
	"strings"
)

type DataType []Command

const (
	FORWARD = iota
	DOWN
	UP
)

type Command struct {
	direction int
	value int
}

func fromLine(s string) (c Command) {
	sFields := strings.Fields(s)

	c.value = ParseInt(sFields[1])

	switch sFields[0] {
	case "forward":
		c.direction = FORWARD
	case "down":
		c.direction = DOWN
	case "up":
		c.direction = UP
	}

	return
}

func parseData(data string) DataType {
	dataSplit := strings.Split(data, "\n")

	result := make(DataType, len(dataSplit))
	for i, v := range dataSplit {
		result[i] = fromLine(v)
	}

	return result
}

func solvePart1(data DataType) int {
	position, depth := 0, 0

	for _, cmd := range data {
		switch cmd.direction {
		case FORWARD:
			position += cmd.value
		case DOWN:
			depth += cmd.value
		case UP:
			depth -= cmd.value
		}
	}

	return position * depth
}

func solvePart2(data DataType) int {
	aim, position, depth := 0, 0, 0

	for _, cmd := range data {
		switch cmd.direction {
		case FORWARD:
			position += cmd.value
			depth += aim * cmd.value
		case DOWN:
			aim += cmd.value
		case UP:
			aim -= cmd.value
		}
	}

	return position * depth
}

func main() {
	data := parseData(FetchInputData(2))
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
