package main

import (
	"reflect"
	"testing"
)

func Test_parseInput(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name string
		args args
		want []lineSegment
	}{
		{
			"t1",
			args{"test.txt"},
			[]lineSegment{
				{point{0, 9}, point{5, 9}},
				{point{8, 0}, point{0, 8}},
				{point{9, 4}, point{3, 4}},
				{point{2, 2}, point{2, 1}},
				{point{7, 0}, point{7, 4}},
				{point{6, 4}, point{2, 0}},
				{point{0, 9}, point{2, 9}},
				{point{3, 4}, point{1, 4}},
				{point{0, 0}, point{8, 8}},
				{point{5, 5}, point{8, 2}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseInput(tt.args.file); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseInput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findIntermediatePoints(t *testing.T) {
	type args struct {
		p1 point
		p2 point
	}
	tests := []struct {
		name string
		args args
		want []point
	}{
		{
			"t1",
			args{p1: point{1, 1}, p2: point{1, 3}},
			[]point{
				{1, 1},
				{1, 2},
				{1, 3},
			},
		},
		{
			"t2",
			args{p1: point{9, 7}, p2: point{7, 7}},
			[]point{
				{7, 7},
				{8, 7},
				{9, 7},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPointsOnVerticalAndHorizontals(tt.args.p1, tt.args.p2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findIntermediatePoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findDiagonalPoints(t *testing.T) {
	type args struct {
		p1 point
		p2 point
	}
	tests := []struct {
		name string
		args args
		want []point
	}{
		{
			"t1",
			args{p1: point{9, 7}, p2: point{7, 9}},
			[]point{
				{9, 7},
				{8, 8},
				{7, 9},
			},
		},
		{
			"t2",
			args{p1: point{1, 1}, p2: point{3, 3}},
			[]point{
				{1, 1},
				{2, 2},
				{3, 3},
			},
		},
		{
			"t3",
			args{p1: point{3, 3}, p2: point{5, 1}},
			[]point{
				{3, 3},
				{4, 2},
				{5, 1},
			},
		},
		{
			"t4",
			args{p1: point{3, 3}, p2: point{1, 5}},
			[]point{
				{3, 3},
				{2, 4},
				{1, 5},
			},
		},
		{
			"t5",
			args{p1: point{5, 1}, p2: point{3, 3}},
			[]point{
				{5, 1},
				{4, 2},
				{3, 3},
			},
		},
		{
			"t6",
			args{p1: point{5, 1}, p2: point{3, 3}},
			[]point{
				{5, 1},
				{4, 2},
				{3, 3},
			},
		},
		{
			"t7",
			args{p1: point{0, 0}, p2: point{8, 8}},
			[]point{
				{0, 0},
				{1, 1},
				{2, 2},
				{3, 3},
				{4, 4},
				{5, 5},
				{6, 6},
				{7, 7},
				{8, 8},
			},
		},
		{
			"t8",
			args{p1: point{6, 4}, p2: point{2, 0}},
			[]point{
				{2, 0},
				{3, 1},
				{4, 2},
				{5, 3},
				{6, 4},
			},
		},
		{
			"t9",
			args{p1: point{2, 0}, p2: point{6, 4}},
			[]point{
				{2, 0},
				{3, 1},
				{4, 2},
				{5, 3},
				{6, 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPointsOnDiagonals(tt.args.p1, tt.args.p2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDiagonalPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
