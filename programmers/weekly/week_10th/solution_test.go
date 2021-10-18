package week_10th

import (
	"reflect"
	"testing"
)

func Test_solution(t *testing.T) {
	type args struct {
		line [][]int
	}

	var tests = []struct {
		name string
		args args
		want []string
	}{
		{
			name: "1st",
			args: args{
				line: [][]int{
					{2, -1, 4},
					{-2, -1, 4},
					{0, -1, 1},
					{5, -8, -12},
					{5, 8, 12},
				},
			},
			want: []string{
				"....*....",
				".........",
				".........",
				"*.......*",
				".........",
				".........",
				".........",
				".........",
				"*.......*",
			},
		},
		{
			name: "2nd",
			args: args{
				line: [][]int{
					{0, 1, -1},
					{1, 0, -1},
					{1, 0, 1},
				},
			},
			want: []string{
				"*.*",
			},
		},
		{
			name: "3rd",
			args: args{
				line: [][]int{
					{1, -1, 0},
					{2, -1, 0},
				},
			},
			want: []string{
				"*",
			},
		},
		{
			name: "4th",
			args: args{
				line: [][]int{
					{1, -1, 0},
					{2, -1, 0},
					{4, -1, 0},
				},
			},
			want: []string{
				"*",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
