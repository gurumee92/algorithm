package week_6th

import (
	"sort"
)

type Player struct {
	number           int
	weight           int
	winPercent       float32
	winsBiggerThanMe int
}

func solution(weights []int, head2head []string) []int {
	n := len(weights)
	players := make([]Player, 0)

	for i := 0; i < n; i++ {
		num := i + 1
		weight := weights[i]
		wins := 0
		loses := 0
		winsBiggerThanMe := 0

		for j, res := range head2head[i] {
			if res == 'W' {
				wins += 1

				if weight < weights[j] {
					winsBiggerThanMe += 1
				}
			} else if res == 'L' {
				loses += 1
			}
		}

		player := Player{
			number:           num,
			weight:           weight,
			winPercent:       float32(wins) * 100 / float32(wins+loses),
			winsBiggerThanMe: winsBiggerThanMe,
		}
		players = append(players, player)
	}

	sort.Slice(players, func(i, j int) bool {
		if players[i].winPercent > players[j].winPercent {
			return true
		} else if players[i].winPercent < players[j].winPercent {
			return false
		}

		if players[i].winsBiggerThanMe > players[j].winsBiggerThanMe {
			return true
		} else if players[i].winsBiggerThanMe < players[j].winsBiggerThanMe {
			return false
		}

		if players[i].weight > players[j].weight {
			return true
		} else if players[i].weight < players[j].weight {
			return false
		}

		return players[i].number < players[j].number
	})

	result := make([]int, 0)

	for _, player := range players {
		result = append(result, player.number)
	}

	return result
}
