package main

import "testing"

func Test_checkBoard(t *testing.T) {
	type args struct {
		board board
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"t1",
			args{
				board: board{
					&[][]int{
						{-1, 2, 3, 4},
						{-1, 3, 4, 4},
						{-1, 3, 5, 5},
					},
					1,
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkBoard(tt.args.board); got != tt.want {
				t.Errorf("checkBoard() = %v, want %v", got, tt.want)
			}
		})
	}
}
