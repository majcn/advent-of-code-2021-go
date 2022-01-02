package main

import (
	. "../util"
	"fmt"
	"sort"
	"strings"
)

type DataType map[Location]int

func parseData(data string) DataType {
	dataSplit := strings.Split(data, "\n")

	result := make(DataType)
	for y, line := range dataSplit {
		for x, v := range line {
			result[Location{X: x, Y: y}] = ParseInt(v)
		}
	}

	return result
}

func getValidNeighbours(grid map[Location]int, location Location) []Location {
	result := make([]Location, 0, 4)
	for _, n := range GetNeighbours4() {
		nLocation := location.Add(n)
		if _, ok := grid[nLocation]; ok {
			result = append(result, nLocation)
		}
	}

	return result
}

func solvePart1(data DataType) (rc int) {
	mainLoop:
	for el := range data {
		for _, n := range getValidNeighbours(data, el) {
			if data[el] >= data[n] {
				continue mainLoop
			}
		}

		rc += data[el] + 1
	}

	return
}

func solvePart2(data DataType) (rc int) {
	allVisited := make(Grid)
	for el := range data {
		if data[el] == 9 {
			allVisited.Add(el)
		}
	}

	result := make([]int, 0)
	for len(data) > len(allVisited) {
		nextEl := Location{}
		for el := range data {
			if !allVisited.Contains(el) {
				nextEl = el
				break
			}
		}

		basinSize := 1
		allVisited.Add(nextEl)
		queue := NewGrid(nextEl)
		for len(queue) > 0 {
			el := queue.Pop()

			for _, n := range getValidNeighbours(data, el) {
				if !allVisited.Contains(n) {
					basinSize++
					allVisited.Add(n)
					queue.Add(n)
				}
			}
		}
		result = append(result, basinSize)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(result)))

	return result[0] * result[1] * result[2]
}

func main() {
	data := parseData(FetchInputData(9))
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
