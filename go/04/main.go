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
		id: 0,
	}
	for i := 2; i < len(lines); i++ {
		if len(lines[i]) == 0 {
			boards = append(boards, curBoard)
			curBoard = board{
				items: func() *[][]int {
					s := [][]int{}
					return &s
				}(),
				id: curBoard.id + 1,
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
		if score, _ := checkBoards(boards, nil); score != -1 {
			return score * pick
		}
	}
	return -1
}

func part2(picks []int, boards []board) int {
	lastScore := -1
	lastNum := -1
	var doneBoards = map[int]bool{}

	for _, pick := range picks {
		markBoards(boards, pick)
		if score, id := checkBoards(boards, doneBoards); score > 0 && !doneBoards[id] {
			doneBoards[id] = true
			lastScore = score
			lastNum = pick

			fmt.Println(pick)
			printBoards(boards)
			fmt.Println()
		}
		// fmt.Println(doneBoards)
		if len(doneBoards) == len(boards) {
			break
		}
	}
	fmt.Println(lastScore, lastNum)
	return lastScore * lastNum
}

type board struct {
	items *[][]int
	id    int
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

func checkBoards(boards []board, doneBoards map[int]bool) (int, int) {
	for _, board := range boards {
		if doneBoards != nil && doneBoards[board.id] {
			continue
		}
		if found := checkBoard(board); found {
			score := findScore(board)
			return score, board.id
		}
	}
	return -1, -1
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
	for row := range *board.items {
		for col := range (*board.items)[row] {
			if (*board.items)[row][col] == n {
				(*board.items)[row][col] = -1
			}
		}
	}
}

func checkBoard(board board) bool {
	if len(*board.items) <= 0 {
		return false
	}
	for row := range *board.items {
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
	for col := range (*board.items)[0] {
		found := true
		for row := range *board.items {
			elem := (*board.items)[row][col]
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
	for row := range *board.items {
		for col := range (*board.items)[row] {
			if (*board.items)[row][col] != -1 {
				s += (*board.items)[row][col]
			}
		}
	}
	return s
}

func main() {
	picks, boards := parseInput("input.txt")
	fmt.Println(part1(picks, boards))
	fmt.Println(part2(picks, boards))
}
