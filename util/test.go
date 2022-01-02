package util

import (
	"os"
	"testing"
)

type AOCTest interface {
	ParseData(string) interface{}

	SolvePart1(interface{}) interface{}
	SolvePart2(interface{}) interface{}

	ExpectedPart1() interface{}
	ExpectedPart2() interface{}
}

func InitData(parseData func(data string)interface{}) interface{} {
	dat, _ := os.ReadFile("./input.txt")
	return parseData(string(dat))
}

func AssertTestPartX(t *testing.T, expected interface{}, actual interface{}) {
	if actual != expected {
		t.Errorf("Result should be %d, got %d.", expected, actual)
	}
}
