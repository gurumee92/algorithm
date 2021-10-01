package week_7th

func solution(enter []int, leave []int) []int {
	n := len(enter)
	result := make([]int, n)
	room := make(map[int]bool)

	for i := 1; i <= n; i++ {
		room[i] = false
	}

	for {
		if len(enter) == 0 {
			break
		}

		enteredPerson := enter[0]
		room[enteredPerson] = true

		for {
			if len(leave) == 0 || !room[leave[0]] {
				break
			}

			leavedPerson := leave[0]
			room[leavedPerson] = false

			for person, isEnter := range room {
				if isEnter {
					result[leavedPerson-1] += 1
					result[person-1] += 1
				}
			}

			leave = leave[1:]
		}

		enter = enter[1:]
	}

	return result
}
