package main

import (
	"fmt"
	. "majcn.si/advent-of-code-2021/util"
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
	result := make([]Location, 0, 8)
	for _, n := range GetNeighbours8() {
		nLocation := location.Add(n)
		if _, ok := grid[nLocation]; ok {
			result = append(result, nLocation)
		}
	}

	return result
}

func solvePartX(data DataType, visitor func(int, Grid)bool) {
	grid := make(map[Location]int, len(data))
	for k, v := range data {
		grid[k] = v
	}

	for i := 1; true; i++ {
		for el := range grid {
			grid[el]++
		}

		flashed := make(Grid)
		flashedNew := NewGrid(Location{})
		for len(flashedNew) > 0 {
			flashedNew = make(Grid)
			for el := range grid {
				if grid[el] > 9 && !flashed.Contains(el) {
					flashedNew.Add(el)
				}
			}

			for el := range flashedNew {
				flashed.Add(el)

				for _, nEl := range getValidNeighbours(grid, el) {
					grid[nEl]++
				}
			}
		}

		for el := range grid {
			if grid[el] > 9 {
				grid[el] = 0
			}
		}

		if !visitor(i, flashed) {
			break
		}
	}
}

func solvePart1(data DataType) (rc int) {
	solvePartX(data, func(i int, flashed Grid) bool {
		rc += len(flashed)
		return i < 100
	})

	return
}

func solvePart2(data DataType) (rc int) {
	solvePartX(data, func(i int, flashed Grid) bool {
		if len(data) == len(flashed) {
			rc = i
			return false
		}

		return true
	})

	return
}

func main() {
	data := parseData(FetchInputData(11))
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
