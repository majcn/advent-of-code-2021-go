package main

import (
	"fmt"
	. "majcn.si/advent-of-code-2021/util"
	"strings"
)

type DataType struct {
	eastGrid Grid
	southGrid Grid
	size Location
}

func parseData(data string) DataType {
	dataSplit := strings.Split(data, "\n")

	eastGrid := make(Grid)
	southGrid := make(Grid)
	size := Location{X: len(dataSplit[0]), Y: len(dataSplit)}
	for y, line := range dataSplit {
		for x, v := range line {
			if v == '>' {
				eastGrid.Add(Location{X: x, Y: y})
			}

			if v == 'v' {
				southGrid.Add(Location{X: x, Y: y})
			}
		}
	}

	return DataType{
		eastGrid:  eastGrid,
		southGrid: southGrid,
		size:      size,
	}
}

func nextStep(eastGrid Grid, southGrid Grid, size Location) (Grid, Grid) {
	newEastGrid := make(Grid, len(eastGrid))
	for l := range eastGrid {
		nl := Location{X: (l.X + 1) % size.X, Y: l.Y}
		if eastGrid.Contains(nl) || southGrid.Contains(nl) {
			newEastGrid.Add(l)
		} else {
			newEastGrid.Add(nl)
		}
	}

	newSouthGrid := make(Grid, len(southGrid))
	for l := range southGrid {
		nl := Location{X: l.X, Y: (l.Y + 1) % size.Y }
		if newEastGrid.Contains(nl) || southGrid.Contains(nl) {
			newSouthGrid.Add(l)
		} else {
			newSouthGrid.Add(nl)
		}
	}

	return newEastGrid, newSouthGrid
}

func isSameGrid(g1 Grid, g2 Grid) bool {
	if len(g1) != len(g2) {
		return false
	}

	for el := range g1 {
		if !g2.Contains(el) {
			return false
		}
	}

	return true
}

func solvePart1(data DataType) (rc int) {
	eastGrid, southGrid, size := data.eastGrid, data.southGrid, data.size

	for i := 0; true ; i++ {
		newEastGrid, newSouthGrid := nextStep(eastGrid, southGrid, size)
		if isSameGrid(eastGrid, newEastGrid) && isSameGrid(southGrid, newSouthGrid) {
			return i + 1
		}
		eastGrid, southGrid = newEastGrid, newSouthGrid
	}

	return -1
}

func main() {
	data := parseData(FetchInputData(25))
	fmt.Println(solvePart1(data))
}
