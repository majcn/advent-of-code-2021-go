package main

import (
	. "../util"
	"fmt"
	"strings"
)

type DataType map[string][]string

func parseData(data string) DataType {
	dataSplit := strings.Split(data, "\n")

	result := make(DataType)
	for _, line := range dataSplit {
		lineSplit := strings.Split(line, "-")
		result[lineSplit[0]] = append(result[lineSplit[0]], lineSplit[1])
		result[lineSplit[1]] = append(result[lineSplit[1]], lineSplit[0])
	}

	return result
}

func findPath(grid map[string][]string, path []string, visited StringSet, twice bool) int {
	if path[len(path)-1] == "end" {
		return 1
	}

	result := 0
	for _, adj := range grid[path[len(path)-1]] {
		if adj == "start" {
			continue
		}

		if adj == strings.ToLower(adj) && visited.Contains(adj) && twice {
			continue
		}

		newTwice := twice || (adj == strings.ToLower(adj) && visited.Contains(adj) && !twice)

		newPath := make([]string, len(path) + 1)
		copy(newPath, path)
		newPath[len(path)] = adj

		newVisited := NewStringSet(visited.AsSlice())
		newVisited.Add(adj)

		result += findPath(grid, newPath, newVisited, newTwice)
	}

	return result
}

func solvePart1(data DataType) (rc int) {
	return findPath(data, []string{"start"}, NewStringSet([]string{"start"}), true)
}

func solvePart2(data DataType) (rc int) {
	return findPath(data, []string{"start"}, NewStringSet([]string{"start"}), false)
}

func main() {
	data := parseData(FetchInputData(12))
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
