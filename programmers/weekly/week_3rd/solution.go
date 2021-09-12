package week_3rd

import "fmt"

func solution(game_board [][]int, table [][]int) int {
	n := len(game_board)
	initVal := 0

	for i:=0; i<n; i++ {
		for j:=0; j<n; j++{
			initVal += table[i][j]
		}
	}

	fmt.Println(initVal)
	return -1
}