package main

import (
	"container/heap"
	"fmt"
	. "majcn.si/advent-of-code-2021/util"
	"strings"
)

type DataType struct {
	grid map[Location]int
	size Location
}

func parseData(data string) DataType {
	dataSplit := strings.Split(data, "\n")

	grid := make(map[Location]int)
	size := Location{X: len(dataSplit[0]), Y: len(dataSplit)}
	for y, line := range dataSplit {
		for x, v := range line {
			grid[Location{X: x, Y: y}] = ParseInt(v)
		}
	}

	return DataType{
		grid: grid,
		size: size,
	}
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

func getLowestRiskPath(grid map[Location]int, start Location, goal Location) int {
	visited := make(Grid)
	gScore := make(map[Location]int)

	q := PriorityQueue{&PriorityQueueItem{Value: start}}
	for len(q) > 0 {
		el := heap.Pop(&q).(*PriorityQueueItem)
		current, score := el.Value.(Location), el.Score

		if current == goal {
			return score
		}

		visited.Add(current)

		for _, neighbour := range getValidNeighbours(grid, current) {
			if visited.Contains(neighbour) {
				continue
			}

			newCost := score + grid[neighbour]
			if _, ok := gScore[neighbour]; !ok || newCost < gScore[neighbour] {
				heap.Push(&q, &PriorityQueueItem{Value: neighbour, Score: newCost})
				gScore[neighbour] = newCost
			}
		}
	}

	return -1
}

func solvePart1(data DataType) (rc int) {
	return getLowestRiskPath(data.grid, Location{X: 0, Y: 0}, data.size.Add(Location{X: -1, Y: -1}))
}

func solvePart2(data DataType) (rc int) {
	nSize := data.size.Mul(5)

	grid := make(map[Location]int, len(data.grid) * 5)
	for x := 0; x < nSize.X; x++ {
		for y := 0; y < nSize.Y; y++ {
			divX, modX := x / data.size.X, x % data.size.X
			divY, modY := y / data.size.Y, y % data.size.Y

			newValue := data.grid[Location{X: modX, Y: modY}] + divX + divY
			newValue = (newValue - 1) % 9 + 1

			grid[Location{X: x, Y: y}] = newValue
		}
	}

	return getLowestRiskPath(grid, Location{X: 0, Y: 0}, nSize.Add(Location{X: -1, Y: -1}))
}

func main() {
	data := parseData(FetchInputData(15))
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
