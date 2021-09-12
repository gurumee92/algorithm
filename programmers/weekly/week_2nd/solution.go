package week_2nd

func scorePointToGrade(score int) string {
	if score >= 90 {
		return "A"
	} else if score >= 80 {
		return "B"
	} else if score >= 70 {
		return "C"
	} else if score >= 50 {
		return "D"
	} else {
		return "F"
	}
}

func solution(scores [][]int) string {
	n := len(scores)
	result := ""

	for i := 0; i < n; i++ {
		myScore := scores[i][i]
		max := myScore
		min := myScore
		sum := 0
		sameCnt := 0

		for j := 0; j < n; j++ {
			score := scores[j][i]
			sum += score

			if max < score {
				max = score
			}

			if min > score {
				min = score
			}

			if myScore == score {
				sameCnt += 1
			}
		}

		average := 0

		if (max == myScore || min == myScore) && sameCnt == 1 {
			sum -= myScore
			average = sum / (n - 1)
		} else {
			average = sum / n
		}

		result += scorePointToGrade(average)
	}

	return result
}