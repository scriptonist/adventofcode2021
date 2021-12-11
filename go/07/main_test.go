package main

import "testing"

func Test_getFuel2(t *testing.T) {
	type args struct {
		positions []int
		bestPos   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"t1",
			args{[]int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}, 5},
			168,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFuel2(tt.args.positions, tt.args.bestPos); got != tt.want {
				t.Errorf("getFuel2() = %v, want %v", got, tt.want)
			}
		})
	}
}
