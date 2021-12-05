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
			if got := findIntermediatePoints(tt.args.p1, tt.args.p2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findIntermediatePoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
