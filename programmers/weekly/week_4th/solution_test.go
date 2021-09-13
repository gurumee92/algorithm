package week_4th

import "testing"

func Test_solution(t *testing.T) {
	type args struct {
		table      []string
		languages  []string
		preference []int
	}

	var tests = []struct {
		name string
		args args
		want string
	}{
		{
			name: "1st",
			args: args{},
			want: "",
		},
		{
			name: "2nd",
			args: args{},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.table, tt.args.languages, tt.args.preference); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
