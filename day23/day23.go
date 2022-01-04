package main

import (
	"container/heap"
	"fmt"
	. "majcn.si/advent-of-code-2021/util"
	"strings"
)

type DataType [8]byte

func parseData(data string) DataType {
	dataSplit := strings.Split(data, "\n")

	result, i := [8]byte{}, 0
	for _, line := range dataSplit {
		for _, v := range line {
			if v == 'A' || v == 'B' || v == 'C' || v == 'D' {
				result[i] = byte(v); i++
			}
		}
	}

	return result
}

type Amphipod struct {
	X int
	Y int
	Type byte
	Moves int
}
func (a *Amphipod) MoveCost() int {
	switch a.Type {
	case 'A':
		return 1
	case 'B':
		return 10
	case 'C':
		return 100
	case 'D':
		return 1000
	}

	return 0
}
func (a *Amphipod) HomeX() int {
	switch a.Type {
	case 'A':
		return 3
	case 'B':
		return 5
	case 'C':
		return 7
	case 'D':
		return 9
	}

	return 0
}

type State [16]Amphipod
func (state *State) isGoal() bool {
	for _, s := range state {
		if s.Moves != 2 {
			return false
		}
	}

	return true
}

func isValidDestinationFromHallwayToRoom(state *State, a *Amphipod, next *Location, maxDepth int) bool {
	if next.Y == 1 {
		return false
	}

	homeX := a.HomeX()

	if next.X != homeX {
		return false
	}

	for _, s := range state {
		if s.X == homeX && s.Type != a.Type {
			return false
		}
	}

	for _, s := range state {
		if s.X == homeX {
			if next.Y + 1  == s.Y {
				return true
			}
		}
	}

	return next.Y == maxDepth
}

func isValidDestinationFromRoomToHallway(next *Location) bool {
	if next.Y != 1 {
		return false
	}

	if next.X == 3 || next.X == 5 || next.X == 7 || next.X == 9 {
		return false
	}

	return true
}

func isValidMove(loc *Location, state *State, grid *Grid) bool {
	if !grid.Contains(*loc) {
		return false
	}

	for _, s := range state {
		if loc.X == s.X && loc.Y == s.Y {
			return false
		}
	}

	return true
}

func getAllNeighbours(a *Amphipod, state *State, grid *Grid) map[Location]int {
	queue := make(map[Location]int)
	queue[Location{X: a.X, Y: a.Y}] = 0

	for true {
		newQueue := make(map[Location]int)
		for queueKey, queueValue := range queue {
			for _, neighbourOffset := range GetNeighbours4() {
				neighbour := queueKey.Add(neighbourOffset)

				if _, ok := queue[neighbour]; ok {
					continue
				}

				if isValidMove(&neighbour, state, grid) {
					newQueue[neighbour] = queueValue + 1
				}
			}
		}

		if len(newQueue) == 0 {
			break
		} else {
			for newQueueKey, newQueueValue := range newQueue {
				queue[newQueueKey] = newQueueValue
			}
		}
	}

	delete(queue, Location{X: a.X, Y: a.Y})

	return queue
}

func getNeighbours(a *Amphipod, state *State, grid *Grid, maxDepth int) map[Location]int {
	if a.Moves == 1 {
		for nLocation, nCost := range getAllNeighbours(a, state, grid) {
			if isValidDestinationFromHallwayToRoom(state, a, &nLocation, maxDepth) {
				return map[Location]int{nLocation: nCost}
			}
		}
	}

	if a.Moves == 0 {
		result := make(map[Location]int)
		for nLocation, nCost := range getAllNeighbours(a, state, grid) {
			if isValidDestinationFromRoomToHallway(&nLocation) {
				result[nLocation] = nCost
			}
		}
		return result
	}

	return map[Location]int{}
}

func solvePartX(initState *State) int {
	maxDepth := 0
	grid := make(Grid)
	for x := 1; x < 12; x++ {
		grid.Add(Location{X: x, Y: 1})
	}
	for _, s := range initState {
		grid.Add(Location{X: s.X, Y: s.Y})
		maxDepth = Max(maxDepth, s.Y)
	}


	h := func(state *State) int {
		result := 0

		finalYDestination := map[byte]int{'A': maxDepth, 'B': maxDepth, 'C': maxDepth, 'D': maxDepth}
		for _, s := range state {
			if s.Moves == 2 {
				finalYDestination[s.Type]--
			}
		}

		for _, s := range state {
			if s.Moves != 2 {
				distanceX := Abs(s.X - s.HomeX())
				distanceY := Abs(s.Y - 1) + Abs(1 - finalYDestination[s.Type])

				result += (distanceX + distanceY) * s.MoveCost()

				finalYDestination[s.Type]--
			}
		}

		return result
	}

	gScore := make(map[State]int)

	q := PriorityQueue{&PriorityQueueItem{Value: *initState}}
	for len(q) > 0 {
		current := heap.Pop(&q).(*PriorityQueueItem).Value.(State)

		if current.isGoal() {
			return gScore[current]
		}

		for i := range current {
			neighbours := getNeighbours(&current[i], &current, &grid, maxDepth)
			for neighbourLoc, neighbourCost := range neighbours {
				neighbour := current
				neighbour[i].X = neighbourLoc.X
				neighbour[i].Y = neighbourLoc.Y
				neighbour[i].Moves++

				newCost := gScore[current] + neighbourCost * current[i].MoveCost()
				if _, ok := gScore[neighbour]; !ok || newCost < gScore[neighbour] {
					gScore[neighbour] = newCost
					fScore := newCost + h(&neighbour)
					heap.Push(&q, &PriorityQueueItem{Value: neighbour, Score: fScore})
				}
			}
		}
	}

	return -1
}

func solvePart1(data DataType) (rc int) {
	initState := State {
		{X: 3, Y: 2, Type: data[0], Moves: 0},
		{X: 3, Y: 3, Type: data[4], Moves: 0},
		{X: 5, Y: 2, Type: data[1], Moves: 0},
		{X: 5, Y: 3, Type: data[5], Moves: 0},
		{X: 7, Y: 2, Type: data[2], Moves: 0},
		{X: 7, Y: 3, Type: data[6], Moves: 0},
		{X: 9, Y: 2, Type: data[3], Moves: 0},
		{X: 9, Y: 3, Type: data[7], Moves: 0},

		{X: 3, Y: 4, Type: 'A', Moves: 2},
		{X: 3, Y: 5, Type: 'A', Moves: 2},
		{X: 5, Y: 4, Type: 'B', Moves: 2},
		{X: 5, Y: 5, Type: 'B', Moves: 2},
		{X: 7, Y: 4, Type: 'C', Moves: 2},
		{X: 7, Y: 5, Type: 'C', Moves: 2},
		{X: 9, Y: 4, Type: 'D', Moves: 2},
		{X: 9, Y: 5, Type: 'D', Moves: 2},
	}

	return solvePartX(&initState)
}

func solvePart2(data DataType) (rc int) {
	initState := State {
		{X: 3, Y: 2, Type: data[0], Moves: 0},
		{X: 3, Y: 5, Type: data[4], Moves: 0},
		{X: 5, Y: 2, Type: data[1], Moves: 0},
		{X: 5, Y: 5, Type: data[5], Moves: 0},
		{X: 7, Y: 2, Type: data[2], Moves: 0},
		{X: 7, Y: 5, Type: data[6], Moves: 0},
		{X: 9, Y: 2, Type: data[3], Moves: 0},
		{X: 9, Y: 5, Type: data[7], Moves: 0},

		{X: 3, Y: 3, Type: 'D', Moves: 0},
		{X: 3, Y: 4, Type: 'D', Moves: 0},
		{X: 5, Y: 3, Type: 'C', Moves: 0},
		{X: 5, Y: 4, Type: 'B', Moves: 0},
		{X: 7, Y: 3, Type: 'B', Moves: 0},
		{X: 7, Y: 4, Type: 'A', Moves: 0},
		{X: 9, Y: 3, Type: 'A', Moves: 0},
		{X: 9, Y: 4, Type: 'C', Moves: 0},
	}

	return solvePartX(&initState)
}

func main() {
	data := parseData(FetchInputData(23))
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
