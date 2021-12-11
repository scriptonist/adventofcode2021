package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := parseInput("input.txt")
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func parseInput(file string) []int {
	bs, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	numsAsStr := strings.Split(string(bs), ",")
	nums := []int{}
	for _, ns := range numsAsStr {
		n, err := strconv.Atoi(ns)
		if err != nil {
			panic(err)
		}
		nums = append(nums, n)
	}
	return nums
}

func part1(positions []int) int {
	bestResult := int(^uint(0) >> 1)
	for _, pos := range positions {
		if fuel := getFuel(positions, pos); fuel < bestResult {
			bestResult = fuel
		}
	}
	return bestResult
}

func getFuel(positions []int, bestPos int) int {
	result := 0
	for _, pos := range positions {
		result += abs(bestPos - pos)
	}
	return result
}

func part2(positions []int) int {
	bestResult := int(^uint(0) >> 1)

	for i := 1; i < 2000; i++ {
		if fuel := getFuel2(positions, i); fuel < bestResult {
			bestResult = fuel
		}
	}
	return bestResult
}

func getFuel2(positions []int, bestPos int) int {
	costMap := map[int]int{}
	for _, pos := range positions {
		costMap[abs(bestPos-pos)]++
	}
	costs := []int{}
	for cost := range costMap {
		costs = append(costs, cost)
	}
	sort.Ints(costs)
	result := 0
	for _, cost := range costs {
		result += ((cost + ((cost * (cost - 1)) / 2)) * costMap[cost])
	}
	return result
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
