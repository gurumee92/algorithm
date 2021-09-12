package week_2nd

import "testing"

func TestScorePointToGrade(t *testing.T) {
	var tests = []struct {
		score    int
		expected string
	}{
		{
			90,
			"A",
		},
		{
			95,
			"A",
		},
		{
			80,
			"B",
		},
		{
			70,
			"C",
		},
		{
			60,
			"D",
		},
		{
			50,
			"D",
		},
		{
			49,
			"F",
		},
	}

	for _, test := range tests {
		result := scorePointToGrade(test.score)

		if test.expected != result {
			t.Errorf("Wrong result expected: %v result: %v\n", test.expected, result)
		}
	}
}

func TestSolution(t *testing.T) {
	var tests = []struct {
		scores   [][]int
		expected string
	}{
		{
			[][]int{
				{100, 90, 98, 88, 65},
				{50, 45, 99, 85, 77},
				{47, 88, 95, 80, 67},
				{61, 57, 100, 80, 65},
				{24, 90, 94, 75, 65},
			},
			"FBABD",
		},
		{
			[][]int{
				{50, 90},
				{50, 87},
			},
			"DA",
		},
		{
			[][]int{
				{70, 49, 90},
				{68, 50, 38},
				{73, 31, 100},
			},
			"CFD",
		},
	}

	for _, test := range tests {
		result := solution(test.scores)

		if test.expected != result {
			t.Errorf("Wrong result expected: %v result: %v\n", test.expected, result)
		}
	}
}
