package main

import (
	. "../util"
	"fmt"
	"regexp"
	"sort"
	"strings"
)

type Command struct {
	state bool
	x1 int
	x2 int
	y1 int
	y2 int
	z1 int
	z2 int
}

type DataType []Command

func fromLine(s string) Command {
	re := `^(on|off) x=(-?\d+)..(-?\d+),y=(-?\d+)..(-?\d+),z=(-?\d+)..(-?\d+)$`
	m := regexp.MustCompile(re).FindStringSubmatch(s)

	return Command{
		state: m[1] == "on",
		x1:    ParseInt(m[2]),
		x2:    ParseInt(m[3]),
		y1:    ParseInt(m[4]),
		y2:    ParseInt(m[5]),
		z1:    ParseInt(m[6]),
		z2:    ParseInt(m[7]),
	}
}

func parseData(data string) DataType {
	dataSplit := strings.Split(data, "\n")

	result := make([]Command, len(dataSplit))
	for i, line := range dataSplit {
		result[i] = fromLine(line)
	}

	return result
}

func filter(commands []Command, filter func(Command)bool) []Command {
	result := make([]Command, 0, len(commands))
	for _, command := range commands {
		if filter(command) {
			result = append(result, command)
		}
	}
	return result
}

func first(commands []Command, filter func(Command)bool) Command {
	for _, command := range commands {
		if filter(command) {
			return command
		}
	}
	return Command{}
}

func solvePartX(commands []Command) (rc int) {
	reverseCommands := make([]Command, 0, len(commands))
	xList := make([]int, 0, len(commands) * 2)
	yList := make([]int, 0, len(commands) * 2)
	zList := make([]int, 0, len(commands) * 2)

	for i := len(commands) - 1; i >= 0; i-- {
		reverseCommands = append(reverseCommands, commands[i])
		xList = append(xList, commands[i].x1, commands[i].x2 + 1)
		yList = append(yList, commands[i].y1, commands[i].y2 + 1)
		zList = append(zList, commands[i].z1, commands[i].z2 + 1)
	}

	sort.Ints(xList)
	sort.Ints(yList)
	sort.Ints(zList)

	for xi := 0; xi < len(xList) - 1; xi++ {
		tmpX := filter(reverseCommands, func(command Command) bool { return command.x1 <= xList[xi] && xList[xi] <= command.x2 })
		for yi := 0; yi < len(yList) - 1; yi++ {
			tmpY := filter(tmpX, func(command Command) bool { return command.y1 <= yList[yi] && yList[yi] <= command.y2 })
			for zi := 0; zi < len(zList) -1; zi++ {
				if first(tmpY, func(command Command) bool { return command.z1 <= zList[zi] && zList[zi] <= command.z2 }).state {
					dx := xList[xi+1] - xList[xi]
					dy := yList[yi+1] - yList[yi]
					dz := zList[zi+1] - zList[zi]
					rc += dx * dy * dz
				}
			}
		}
	}

	return
}

func solvePart1(data DataType) (rc int) {
	return solvePartX(filter(data, func(command Command) bool {
		return command.x1 >= -50 && command.x2 <= 50 &&
			command.y1 >= -50 && command.y2 <= 50 &&
			command.z1 >= -50 && command.z2 <= 50
	}))
}

func solvePart2(data DataType) (rc int) {
	return solvePartX(data)
}

func main() {
	data := parseData(FetchInputData(22))
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
