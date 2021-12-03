package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

func parseInput(file string) string {
	bs, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return string(bs)
}

func part1(input string) int {
	report := []string{}
	report = append(report, strings.Split(input, "\n")...)

	noOfBits := len(report[0])
	gammaRate := 0
	epsilonRate := 0
	for i := 0; i < noOfBits; i++ {
		noOfOnes := 0
		noOfZeros := 0
		for _, entry := range report {
			switch rune(entry[i]) {
			case '1':
				noOfOnes++
			case '0':
				noOfZeros++
			}
		}
		if noOfOnes > noOfZeros {
			gammaRate += int(math.Pow(2.0, float64(noOfBits-1-i)))
		} else {
			epsilonRate += int(math.Pow(2.0, float64(noOfBits-1-i)))
		}
	}
	return gammaRate * epsilonRate
}

// 10110
// 01001

func part2(input string) int {
	report := []string{}
	report = append(report, strings.Split(input, "\n")...)
	/*
		[00100, 11110, 10110, 10111,10101,01111,00111,11100,10000,11001, 00010, 01010]
		[00100, 11110, 10110, 10111,10101,01111,00111,11100,10000,11001, 00010, 01010]

		- walk through each position in items
		- at each position populate a "disabled" map which is a map from index of an entry to a bool
		-
	*/

	oxygenRate := 0
	co2Rate := 1
	bitCount := len(report[0])

	disabled := map[int]bool{}
	for idx := range report {
		disabled[idx] = false
	}
	disabledCount := 0
	for i := 0; i < bitCount; i++ {
		onesCount, zerosCount := 0, 0
		for idx, entry := range report {
			if !disabled[idx] {
				switch rune(entry[i]) {
				case '1':
					onesCount++
				case '0':
					zerosCount++
				}
			}
		}
		if onesCount >= zerosCount {
			for idx, entry := range report {
				if !disabled[idx] && rune(entry[i]) == '0' {
					disabled[idx] = true
					disabledCount++
				}
			}
		} else {
			for idx, entry := range report {
				if !disabled[idx] && rune(entry[i]) == '1' {
					disabled[idx] = true
					disabledCount++
				}
			}
		}
		// fmt.Println(i)
		// printMap(disabled, report)
		if disabledCount == len(report)-1 {
			break
		}
	}
	var item string
	for idx, disabled := range disabled {
		if !disabled {
			item = report[idx]
		}
	}
	oxygenRate = makeNumber(item)

	for idx := range report {
		disabled[idx] = false
	}
	disabledCount = 0
	for i := 0; i < bitCount; i++ {
		onesCount, zerosCount := 0, 0
		for idx, entry := range report {
			if !disabled[idx] {
				switch rune(entry[i]) {
				case '1':
					onesCount++
				case '0':
					zerosCount++
				}
			}
		}
		if zerosCount <= onesCount {
			for idx, entry := range report {
				if !disabled[idx] && rune(entry[i]) == '1' {
					disabled[idx] = true
					disabledCount++
				}
			}
		} else {
			for idx, entry := range report {
				if !disabled[idx] && rune(entry[i]) == '0' {
					disabled[idx] = true
					disabledCount++
				}
			}
		}
		// fmt.Println(i)
		// printMap(disabled, report)
		if disabledCount == len(report)-1 {
			break
		}
	}
	for idx, disabled := range disabled {
		if !disabled {
			item = report[idx]
		}
	}
	co2Rate = makeNumber(item)
	// fmt.Println(item, co2Rate)
	return oxygenRate * co2Rate
}

func makeNumber(s string) int {
	n := 0
	for i, c := range s {
		if c == '1' {
			n += int(math.Pow(2.0, float64(len(s)-1-i)))
		}
	}
	return n
}

// func printMap(m map[int]bool, report []string) {
// 	for k, v := range m {
// 		if !v {
// 			fmt.Printf("%s: %v ", report[k], v)
// 		}
// 	}
// 	fmt.Println()
// }
func main() {
	input := parseInput("input.txt")
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
