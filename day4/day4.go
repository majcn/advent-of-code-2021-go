package main

import (
	. "../util"
	"fmt"
	"strings"
)

const (
	SIZE = 5
)

type DataType struct {
	line []int
	grids []map[Location]int
}

func parseData(data string) DataType {
	dataSplit := strings.Split(data, "\n")

	lineFields := strings.Split(dataSplit[0], ",")
	line := make([]int, len(lineFields))
	for i, v := range lineFields {
		line[i] = ParseInt(v)
	}

	grids := make([]map[Location]int, 0)
	offset := 2
	for (offset + SIZE) <= len(dataSplit) {
		grid := make(map[Location]int)
		for y, xLine := range dataSplit[offset:offset+SIZE] {
			for x, v := range strings.Fields(xLine) {
				grid[Location{X: x, Y: y}] = ParseInt(v)
			}
		}
		offset += SIZE + 1
		grids = append(grids, grid)
	}

	return DataType{
		line:  line,
		grids: grids,
	}
}

func checkForWin(grid Grid) bool {
	for x := 0; x < SIZE; x++ {
		c := true
		for y := 0; y < SIZE; y++ {
			if !grid.Contains(Location{X: x, Y: y}) {
				c = false
				break
			}
		}

		if c {
			return true
		}
	}

	for y := 0; y < SIZE; y++ {
		c := true
		for x := 0; x < SIZE; x++ {
			if !grid.Contains(Location{X: x, Y: y}) {
				c = false
				break
			}
		}

		if c {
			return true
		}
	}

	return false
}

func solvePartX(data DataType, resultFunc func(int, int, map[Location]int, Grid)bool) (map[Location]int, Grid, int) {
	winners := make([]Grid, len(data.grids))
	for i := 0; i < len(data.grids); i++ {
		winners[i] = make(Grid)
	}

	for _, n := range data.line {
		for i, grid := range data.grids {
			for x := 0; x < SIZE; x++ {
				for y := 0; y < SIZE; y++ {
					l := Location{X: x, Y: y}
					if grid[l] == n {
						winners[i].Add(l)
						if checkForWin(winners[i]) {
							if resultFunc(n, i, grid, winners[i]) {
								return grid, winners[i], n
							}
						}
					}
				}
			}
		}
	}

	return nil, nil, -1
}

func score(grid map[Location]int, winner Grid, n int) int {
	c := 0
	for x := 0; x < SIZE; x++ {
		for y:= 0; y < SIZE; y++ {
			l := Location{X: x, Y: y}
			if !winner.Contains(l) {
				c += grid[l]
			}
		}
	}

	return c * n
}

func solvePart1(data DataType) int {
	grid, winner, n := solvePartX(data, func(in int, ii int, igrid map[Location]int, iwinner Grid)bool {
		return true
	})

	return score(grid, winner, n)
}

func solvePart2(data DataType) int {
	wGrids := make(map[int]bool)
	grid, winner, n := solvePartX(data, func(in int, ii int, igrid map[Location]int, iwinner Grid)bool {
		wGrids[ii] = true
		if len(wGrids) == len(data.grids) {
			return true
		}
		return false
	})

	return score(grid, winner, n)
}

func main() {
	data := parseData(FetchInputData(4))
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
