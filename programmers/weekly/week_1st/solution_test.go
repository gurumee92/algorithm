package week_1st

import "testing"

func TestArithmeticSequence(t *testing.T) {
	var tests = []struct {
		n        int
		expected int
	}{
		{1, 1},
		{2, 3},
		{3, 6},
		{4, 10},
		{5, 15},
	}

	for _, test := range tests {
		result := arithmeticSequence(test.n)

		if test.expected != result {
			t.Errorf("Wrong result sum :%v result: %v\n", test.expected, result)
		}
	}
}

func TestSolution(t *testing.T) {
	var tests = []struct {
		price    int
		money    int
		count    int
		expected int64
	}{
		{3, 20, 4, 10},
		{3, 20, 3, 0},
	}

	for _, test := range tests {
		result := solution(test.price, test.money, test.count)

		if test.expected != result {
			t.Errorf("Wrong result sum :%v result: %v\n", test.expected, result)
		}
	}
}
