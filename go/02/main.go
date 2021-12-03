package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func parseInput(file string) string {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func part1(input string) int {
	forward, depth := 0, 0
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " ")
		direction, valueS := parts[0], parts[1]
		value, err := strconv.Atoi(valueS)
		if err != nil {
			panic(err)
		}
		switch direction {
		case "forward":
			forward += value
		case "down":
			depth += value
		case "up":
			depth -= value
		default:
			panic(fmt.Sprintf("invalid: %s", direction))
		}
	}

	return forward * depth
}

func part2(input string) int {
	forward, depth, aim := 0, 0, 0
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " ")
		direction, valueS := parts[0], parts[1]
		value, err := strconv.Atoi(valueS)
		if err != nil {
			panic(err)
		}
		switch direction {
		case "forward":
			forward += value
			depth += aim * value
		case "down":
			aim += value
		case "up":
			aim -= value
		default:
			panic(fmt.Sprintf("invalid: %s", direction))
		}
	}

	return forward * depth
}

func main() {
	input := parseInput("input.txt")
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
