package main

import (
	"fmt"
	. "majcn.si/advent-of-code-2021/util"
	"strings"
)

type DataType struct {
	template string
	rules map[string]byte
}

func parseData(data string) DataType {
	dataSplit := strings.Split(data, "\n")

	template := dataSplit[0]

	rules := make(map[string]byte, len(dataSplit) - 2)
	for _, line := range dataSplit[2:] {
		lineFields := strings.Split(line, " -> ")
		rules[lineFields[0]] = lineFields[1][0]
	}

	return DataType{
		template: template,
		rules:    rules,
	}
}

type CacheKey struct {
	a byte
	b byte
	i int
}

type Counter map[byte]int
func (c Counter) Add(c2 Counter) {
	for k := range c2 {
		c[k] += c2[k]
	}
}

func f(data DataType, a byte, b byte, limit int) Counter {
	cache := make(map[CacheKey]Counter)

	var f func(a byte, b byte, limit int, i int) Counter
	f = func(a byte, b byte, limit int, i int) Counter {
		if i == limit {
			return Counter{}
		}

		cacheKey := CacheKey{a: a, b: b, i: i}
		if _, ok := cache[cacheKey]; ok {
			return cache[cacheKey]
		}

		newC := data.rules[string([]byte{a,b})]

		counter := Counter{newC: 1}
		counter.Add(f(a, newC, limit, i+1))
		counter.Add(f(newC, b, limit, i+1))

		cache[cacheKey] = counter

		return counter
	}

	return f(a, b, limit, 0)
}

func solvePartX(data DataType, limit int) int {
	template := data.template

	counter := make(Counter)
	for i := 0; i < len(template) - 1; i++ {
		counter.Add(f(data, template[i], template[i+1], limit))
		counter[template[i]]++
	}
	counter[template[len(template)-1]]++

	max := counter[template[0]]
	min := counter[template[0]]
	for _, v := range counter {
		min = Min(min, v)
		max = Max(max, v)
	}
	return max - min
}

func solvePart1(data DataType) (rc int) {
	return solvePartX(data, 10)
}

func solvePart2(data DataType) (rc int) {
	return solvePartX(data, 40)
}

func main() {
	data := parseData(FetchInputData(14))
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
