package main

import (
	"fmt"
	. "majcn.si/advent-of-code-2021/util"
	"regexp"
)

type DataType struct {
	X1 int
	X2 int
	Y1 int
	Y2 int
}

func parseData(data string) DataType {
	re := `^target area: x=(-?\d+)..(-?\d+), y=(-?\d+)..(-?\d+)$`
	m := regexp.MustCompile(re).FindStringSubmatch(data)

	return DataType{
		X1: ParseInt(m[1]),
		X2: ParseInt(m[2]),
		Y1: ParseInt(m[3]),
		Y2: ParseInt(m[4]),
	}
}

type Probe struct {
	Position Location
	Velocity Location
	MaxY int
}
func (p *Probe) nextStep() {
	p.Position = p.Position.Add(p.Velocity)

	p.Velocity = p.Velocity.Add(Location{X: -1,Y: -1})
	if p.Velocity.X < 0 {
		p.Velocity.X = 0
	}

	p.MaxY = Max(p.MaxY, p.Position.Y)
}
func (p *Probe) canReachGoal(data DataType) bool {
	if p.Velocity.X == 0 && data.X1 > p.Position.X && p.Position.X > data.X2 {
		return false
	}

	if p.Velocity.X > 0 && p.Position.X > data.X2 {
		return false
	}

	if p.Velocity.Y < 0 {
		return p.Position.Y > data.Y1
	}

	return true
}
func (p *Probe) inGoal(data DataType) bool {
	return data.X1 <= p.Position.X && p.Position.X <= data.X2 && data.Y1 <= p.Position.Y && p.Position.Y <= data.Y2
}

func solvePartX(data DataType, visitor func(Probe)) {
	magicNumber := 107

	bruteForce := make([]Probe, 0)
	for vx := 0; vx <= data.X2; vx++ {
		for vy := data.Y1; vy <= magicNumber; vy++ {
			bruteForce = append(bruteForce, Probe{Velocity: Location{X: vx, Y: vy}})
		}
	}

	for _, probe := range bruteForce {
		for true {
			probe.nextStep()
			if probe.inGoal(data) {
				visitor(probe)
				break
			}
			if !probe.canReachGoal(data) {
				break
			}
		}
	}
}

func solvePart1(data DataType) (rc int) {
	solvePartX(data, func(probe Probe) {
		rc = Max(rc, probe.MaxY)
	})
	return
}

func solvePart2(data DataType) (rc int) {
	solvePartX(data, func(probe Probe) {
		rc++
	})
	return
}

func main() {
	data := parseData(FetchInputData(17))
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
