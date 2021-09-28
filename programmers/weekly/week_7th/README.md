# \[ALGORITHM\] 입실 퇴실

## 문제

이 문서에서 다룰 문제 "입실 퇴실"는 "프로그래머스 > 코딩테스트 연습 > 위클리 챌린지 > 7주차" 문제이다. 문제 링크는 다음과 같다.

* [https://programmers.co.kr/learn/courses/30/lessons/86048](https://programmers.co.kr/learn/courses/30/lessons/86048)

## 문제 풀이



## 코드

전체 코드는 다음과 같다.

```go
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
```