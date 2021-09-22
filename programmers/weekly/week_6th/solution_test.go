package week_6th

import (
	"reflect"
	"testing"
)

func Test_solution(t *testing.T) {
	type args struct {
		weights   []int
		head2head []string
	}

	var tests = []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1st",
			args: args{
				weights:   []int{50, 82, 75, 120},
				head2head: []string{"NLWL", "WNLL", "LWNW", "WWLN"},
			},
			want: []int{3, 4, 1, 2},
		},
		{
			name: "2nd",
			args: args{
				weights:   []int{145, 92, 86},
				head2head: []string{"NLW", "WNL", "LWN"},
			},
			want: []int{2, 3, 1},
		},
		{
			name: "3rd",
			args: args{
				weights:   []int{60, 70, 60},
				head2head: []string{"NNN", "NNN", "NNN"},
			},
			want: []int{2, 1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.weights, tt.args.head2head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
