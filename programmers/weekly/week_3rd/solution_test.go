package week_3rd

import "testing"

func Test_solution(t *testing.T) {
	type args struct {
		game_board [][]int
		table      [][]int
	}

	var tests = []struct {
		name string
		args args
		want int
	}{
		{
			name: "1st",
			args: args{
				game_board: [][]int{
					{1,1,0,0,1,0},
					{0,0,1,0,1,0},
					{0,1,1,0,0,1},
					{1,1,0,1,1,1},
					{1,0,0,0,1,0},
					{0,1,1,1,0,0},
				},
				table: [][]int{
					{1,0,0,1,1,0},
					{1,0,1,0,1,0},
					{0,1,1,0,1,1},
					{0,0,1,0,0,0},
					{1,1,0,1,1,0},
					{0,1,0,0,0,0},
				},
			},
			want: 14,
		},
		{
			name: "2nd",
			args: args{
				game_board: [][]int{
					{0,0,0},
					{1,1,0},
					{1,1,1},
				},
				table: [][]int{
					{1,1,1},
					{1,0,0},
					{0,0,0},
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.game_board, tt.args.table); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
