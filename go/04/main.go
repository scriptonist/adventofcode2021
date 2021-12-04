package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func parseInput(file string) ([]int, []board) {
	bs, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(bs), "\n")
	picks := genNumArrayFromString(lines[0], ",")
	var boards []board
	curBoard := board{
		items: func() *[][]int {
			s := [][]int{}
			return &s
		}(),
	}
	for i := 2; i < len(lines); i++ {
		if len(lines[i]) == 0 {
			boards = append(boards, curBoard)
			curBoard = board{
				items: func() *[][]int {
					s := [][]int{}
					return &s
				}(),
			}
			continue
		}
		*curBoard.items = append(*curBoard.items, genNumArrayFromString(lines[i], " "))
	}
	boards = append(boards, curBoard)

	return picks, boards
}

// 22 13 17 11 0 -> []int{22,13,17,11,0}
func genNumArrayFromString(s string, sep string) []int {
	ns := strings.Split(s, sep)
	// fmt.Println(ns)
	nums := []int{}
	for _, item := range ns {
		item = strings.Trim(item, " ")
		if len(item) == 0 {
			continue
		}
		// fmt.Println([]byte(item))
		n, err := strconv.Atoi(item)
		if err != nil {
			panic(err)
		}
		nums = append(nums, n)
	}
	return nums
}

/*
- draw each number and mark  change number at position to -1
- check boards

*/
func part1(picks []int, boards []board) int {
	for _, pick := range picks {
		markBoards(boards, pick)
		fmt.Println(pick)
		printBoards(boards)
		fmt.Println()
		if score := checkBoards(boards); score != -1 {
			return score * pick
		}
	}
	return -1
}

func part2(picks []int, boards []board) int {
	lastScore := -1
	lastNum := -1
	for _, pick := range picks {
		markBoards(boards, pick)
		fmt.Println(pick)
		printBoards(boards)
		fmt.Println()
		if score := checkBoards(boards); score != -1 {
			lastScore = score
			lastNum = pick
		}
	}
	return lastScore * lastNum
}

type board struct {
	items *[][]int
}

func (b *board) print() {
	for row := range *b.items {
		for col := range (*b.items)[row] {
			fmt.Printf("%d ", (*b.items)[row][col])
		}
		fmt.Println()
	}
	fmt.Println()
}

func checkBoards(boards []board) int {
	for _, board := range boards {
		if found := checkBoard(board); found {
			score := findScore(board)
			return score
		}
	}
	return -1
}

func printBoards(boards []board) {
	for _, b := range boards {
		b.print()
	}
}

func markBoards(boards []board, n int) {
	for _, b := range boards {
		markBoard(b, n)
	}
}

func markBoard(board board, n int) {
	for row, _ := range *board.items {
		for col, _ := range (*board.items)[row] {
			if (*board.items)[row][col] == n {
				(*board.items)[row][col] = -1
			}
		}
	}
}

func checkBoard(board board) bool {
	for row, _ := range *board.items {
		found := true
		for _, elem := range (*board.items)[row] {
			if elem != -1 {
				found = false
			}
		}
		if found {
			return true
		}
	}

	return false
}

func findScore(board board) int {
	s := 0
	for row, _ := range *board.items {
		for col, _ := range (*board.items)[row] {
			if (*board.items)[row][col] != -1 {
				s += (*board.items)[row][col]
			}
		}
	}
	return s
}

func main() {
	picks, boards := parseInput("test.txt")
	fmt.Println(part1(picks, boards))
	fmt.Println(part2(picks, boards))
}
