package main

import (
	"fmt"
	. "majcn.si/advent-of-code-2021/util"
	"strings"
)

type Fold struct {
	along byte
	value int
}

type DataType struct {
	grid Grid
	folds []Fold
}

func parseData(data string) DataType {
	dataSplitGrid := strings.Split(strings.Split(data, "\n\n")[0], "\n")
	dataSplitFolds := strings.Split(strings.Split(data, "\n\n")[1], "\n")

	grid := make(Grid, len(dataSplitGrid))
	for _, line := range dataSplitGrid {
		lineFields := strings.Split(line, ",")
		location := Location{
			X: ParseInt(lineFields[0]),
			Y: ParseInt(lineFields[1]),
		}
		grid.Add(location)
	}

	folds := make([]Fold, len(dataSplitFolds))
	for i, line := range dataSplitFolds {
		folds[i] = Fold{
			along: line[11],
			value: ParseInt(line[13:]),
		}
	}

	return DataType{
		grid:  grid,
		folds: folds,
	}
}

func solveFold(grid Grid, fold Fold) Grid {
	result := make(Grid, len(grid))
	for el := range grid {
		newEl := el

		if fold.along == 'y' && el.Y >= fold.value {
			newEl.Y = 2*fold.value - el.Y
		}

		if fold.along == 'x' && el.X >= fold.value {
			newEl.X = 2*fold.value - el.X
		}

		result.Add(newEl)
	}

	return result
}

func solvePart1(data DataType) (rc int) {
	return len(solveFold(data.grid, data.folds[0]))
}

func solvePart2(data DataType) (rc string) {
	grid := data.grid
	for _, fold := range data.folds {
		grid = solveFold(grid, fold)
	}

	maxX := 0
	maxY := 0
	for el := range grid {
		maxX = Max(maxX, el.X)
		maxY = Max(maxY, el.Y)
	}

	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if grid.Contains(Location{X: x, Y: y}) {
				rc += "#"
			} else {
				rc += " "
			}
		}
		rc += "\n"
	}

	return
}

func main() {
	data := parseData(FetchInputData(13))
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
