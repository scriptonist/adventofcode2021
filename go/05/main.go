package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

type lineSegment struct {
	p1, p2 point
}

func parseInput(file string) []lineSegment {
	bs, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(bs), "\n")
	segments := []lineSegment{}
	for _, line := range lines {
		parts := strings.Split(line, " ")
		segmentString := parts[0]
		segmentNumsString := strings.Split(segmentString, ",")

		p1 := point{}
		p1.x, err = strconv.Atoi(segmentNumsString[0])
		if err != nil {
			panic(err)
		}
		p1.y, err = strconv.Atoi(segmentNumsString[1])
		if err != nil {
			panic(err)
		}

		segmentString = parts[2]
		segmentNumsString = strings.Split(segmentString, ",")

		p2 := point{}
		p2.x, err = strconv.Atoi(segmentNumsString[0])
		if err != nil {
			panic(err)
		}
		p2.y, err = strconv.Atoi(segmentNumsString[1])
		if err != nil {
			panic(err)
		}
		segments = append(segments, lineSegment{p1, p2})
	}

	return segments
}

/*

x1,y1: 2
x2,y2: 1
- Given two points find list of all points in that line segment
- store these points in map to it's count
x1,y1: 1
x2,y2: 2
*/
func part1(segments []lineSegment) int {
	pointsMap := map[point]int{}
	for _, segment := range segments {
		points := findPointsOnVerticalAndHorizontals(segment.p1, segment.p2)
		for _, p := range points {
			pointsMap[p]++
		}
	}
	result := 0
	for _, count := range pointsMap {
		if count > 1 {
			result++
		}
	}
	// fmt.Println(pointsMap)
	return result
}

func part2(segments []lineSegment) int {
	pointsMap := map[point]int{}
	for _, segment := range segments {
		points := findPointsOnVerticalAndHorizontals(segment.p1, segment.p2)
		points = append(points, findPointsOnDiagonals(segment.p1, segment.p2)...)
		// fmt.Println(segment, points)
		for _, p := range points {
			pointsMap[p]++
		}
	}
	result := 0
	for _, count := range pointsMap {
		if count > 1 {
			result++
		}
	}
	// for k, v := range pointsMap {
	// 	fmt.Println(k, v)
	// }
	return result
}

func findPointsOnVerticalAndHorizontals(p1, p2 point) []point {
	var from, to int
	if p1.x == p2.x {
		from = p1.y
		to = p2.y
	} else if p1.y == p2.y {
		from = p1.x
		to = p2.x
	}
	points := []point{}
	if from > to {
		to, from = from, to
	}
	for i := from; i <= to; i++ {
		if p1.x == p2.x {
			points = append(points, point{p1.x, i})
		} else if p1.y == p2.y {
			points = append(points, point{i, p1.y})
		}
	}
	return points
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func findPointsOnDiagonals(p1, p2 point) []point {
	if !(abs(p1.x-p1.y) == abs(p2.x-p2.y)) && !(p1.x+p1.y == p2.x+p2.y) && !(p1.x == p1.y && p2.x == p2.y) {
		return []point{}
	}
	points := []point{}

	if (p1.x == p1.y) && (p2.x == p2.y) {
		from, to := p1.x, p2.x
		if from > to {
			from, to = to, from
		}
		for i := from; i <= to; i++ {
			points = append(points, point{i, i})
		}
	} else if p1.x+p1.y == p2.x+p2.y {
		if p1.x > p2.x {
			for i, j := p1.x, p1.y; i >= p2.x; i, j = i-1, j+1 {
				points = append(points, point{i, j})
			}
		} else if p1.x < p2.x {
			for i, j := p1.x, p1.y; i <= p2.x; i, j = i+1, j-1 {
				points = append(points, point{i, j})
			}
		}
	} else if abs(p1.x-p1.y) == abs(p2.x-p2.y) {
		if p1.x < p2.x {
			for i, j := p1.x, p1.y; i <= p2.x; i, j = i+1, j+1 {
				points = append(points, point{i, j})
			}
		} else if p1.x > p2.x {
			for i, j := p2.x, p2.y; i <= p1.x; i, j = i+1, j+1 {
				points = append(points, point{i, j})
			}
		}
	}
	return points
}

func main() {
	segments := parseInput("input.txt")
	fmt.Println(part1(segments))
	fmt.Println(part2(segments))
}
