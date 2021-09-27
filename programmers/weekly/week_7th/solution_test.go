package week_7th

import (
	"reflect"
	"testing"
)

func Test_solution(t *testing.T) {
	type args struct {
		enter []int
		leave []int
	}

	var tests = []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1st",
			args: args{
				enter: []int{1, 3, 2},
				leave: []int{1, 2, 3},
			},
			want: []int{0, 1, 1},
		},
		{
			name: "2nd",
			args: args{
				enter: []int{1, 4, 2, 3},
				leave: []int{2, 1, 3, 4},
			},
			want: []int{2, 2, 1, 3},
		},
		{
			name: "3rd",
			args: args{
				enter: []int{3, 2, 1},
				leave: []int{2, 1, 3},
			},
			want: []int{1, 1, 2},
		},
		{
			name: "4th",
			args: args{
				enter: []int{3, 2, 1},
				leave: []int{1, 3, 2},
			},
			want: []int{2, 2, 2},
		},
		{
			name: "5th",
			args: args{
				enter: []int{1, 4, 2, 3},
				leave: []int{2, 1, 4, 3},
			},
			want: []int{2, 2, 0, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.enter, tt.args.leave); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
