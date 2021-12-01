package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func findIncreasedReadings(file string) int {
	numbers := []int{}
	bs, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	for _, line := range strings.Split(string(bs), "\n") {
		n, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, n)
	}

	c := 0
	for i := 1; i < len(numbers); i++ {
		if numbers[i] > numbers[i-1] {
			c++
		}
	}

	return c
}

func findIncreasedReadingsFromSlidingWindow(file string) int {
	numbers := []int{}
	bs, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	for _, line := range strings.Split(string(bs), "\n") {
		n, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, n)
	}

	c := -1
	prevSum := -1
	for i := 0; i+2 < len(numbers); i++ {
		s := numbers[i] + numbers[i+1] + numbers[i+2]
		if s > prevSum {
			c++
		}
		prevSum = s
	}

	return c
}

func main() {
	fmt.Println(findIncreasedReadings("input.txt"))
	fmt.Println(findIncreasedReadingsFromSlidingWindow("input.txt"))
}
