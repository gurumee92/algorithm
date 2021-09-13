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
			args: args{
				table:      []string{"SI JAVA JAVASCRIPT SQL PYTHON C#", "CONTENTS JAVASCRIPT JAVA PYTHON SQL C++", "HARDWARE C C++ PYTHON JAVA JAVASCRIPT", "PORTAL JAVA JAVASCRIPT PYTHON KOTLIN PHP", "GAME C++ C# JAVASCRIPT C JAVA"},
				languages:  []string{"PYTHON", "C++", "SQL"},
				preference: []int{7, 5, 5},
			},
			want: "HARDWARE",
		},
		{
			name: "2nd",
			args: args{
				table:      []string{"SI JAVA JAVASCRIPT SQL PYTHON C#", "CONTENTS JAVASCRIPT JAVA PYTHON SQL C++", "HARDWARE C C++ PYTHON JAVA JAVASCRIPT", "PORTAL JAVA JAVASCRIPT PYTHON KOTLIN PHP", "GAME C++ C# JAVASCRIPT C JAVA"},
				languages:  []string{"JAVA", "JAVASCRIPT"},
				preference: []int{7, 5},
			},
			want: "PORTAL",
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
