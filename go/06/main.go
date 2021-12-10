package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input := parseInput("input.txt")
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(timers []int) int {
	return countFishesAfterRepro2(timers, 80)
}

func part2(timers []int) int {
	return countFishesAfterRepro3(timers, 256)
}

func parseInput(file string) []int {
	bs, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	numsString := strings.Split(string(bs), ",")
	nums := []int{}
	for _, n := range numsString {
		num, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		nums = append(nums, num)
	}
	return nums
}

func countFishesAfterRepro1(timers []int, days int) int {
	for day := 0; day < days; day++ {
		for fish, timer := range timers {
			if timer == 0 {
				timers = append(timers, 8)
				timers[fish] = 6
			} else {
				timers[fish] -= 1
			}
		}
	}
	return len(timers)
}

func countFishesAfterRepro2(timers []int, days int) int {
	for day := days; ; {
		lowest := findLowest(timers)
		day -= lowest
		for idx := range timers {
			if timers[idx] == 0 {
				timers = append(timers, 8-(lowest-1))
				timers[idx] = 6 - (lowest - 1)
			} else {
				timers[idx] -= lowest
			}
		}
		if day <= 0 {
			break
		}

	}
	return len(timers)
}

func findLowest(ns []int) int {
	lowest := int(^uint(1) >> 1)
	for _, n := range ns {
		if n < lowest && n != 0 {
			lowest = n
		}
	}
	return lowest
}

/*
Fishes can life between 0 through 8
so store number of fishes of each life

fish/   d0 d1 d2
0: 	0  1  1
1: 	1  1  2
2: 	1  2  1
3: 	2  1  0
4: 	1  0  0
5: 	0  0  1
6: 	0  1  1
7: 	0  0  1
8: 	0  1  1
*/
func countFishesAfterRepro3(timers []int, days int) int {
	daysCount := make([]int, 9)
	for _, time := range timers {
		daysCount[time]++
	}
	for day := 0; day < days; day++ {
		elapsedCount := daysCount[0]
		for i := 0; i <= 7; i++ {
			if daysCount[i+1] > 0 {
				daysCount[i] = daysCount[i+1]
			} else {
				daysCount[i] = 0
			}
		}
		daysCount[8] = 0
		daysCount[6] += elapsedCount
		daysCount[8] += elapsedCount
	}
	return sum(daysCount)
}

func sum(ns []int) int {
	s := 0
	for _, n := range ns {
		s += n
	}
	return s
}
