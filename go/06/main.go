package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

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
func part1(timers []int) int {
	return countFishesAfterRepro1(timers, 80)
}

func main() {
	input := parseInput("input.txt")
	fmt.Println(part1(input))
}
