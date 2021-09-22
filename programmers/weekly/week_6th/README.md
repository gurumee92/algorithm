# \[ALGORITHM\] 복서 정렬하기

## 문제

이 문서에서 다룰 문제 "복서 정렬하기"는 "프로그래머스 > 코딩테스트 연습 > 위클리 챌린지 > 6주차" 문제이다. 문제 링크는 다음과 같다.

* [https://programmers.co.kr/learn/courses/30/lessons/85002](https://programmers.co.kr/learn/courses/30/lessons/85002)

## 문제 풀이

먼저 문제의 입력은 다음과 같다.

* weights []int 
* head2head []string

제한 사항은 다음과 같다. 

* 2 <= len(weights) <= 1000
  * 45 <= weights[i] <= 150
  * weights[i] = i+1번 선수의 무게
* len(head2head) == len(weights)
  * head2head의 모든 문자열은 길이가 weights의 길이와 동일하며, 'N', 'W', 'L'로 이루어진 문자열입니다.
  * head2head[i] 는 i+1번 복서의 전적을 의미하며, head2head[i][j]는 i+1번 복서와 j+1번 복서의 매치 결과를 의미합니다
    * 'N' (None)은 두 복서가 아직 붙어본 적이 없음을 의미합니다.
    * 'W' (Win)는 i+1번 복서가 j+1번 복서를 이겼음을 의미합니다.
    * 'L' (Lose)는 i+1번 복사가 j+1번 복서에게 졌음을 의미합니다.
  * 임의의 i에 대해서 head2head[i][i] 는 항상 'N'입니다. 자기 자신과 싸울 수는 없기 때문입니다.
  * 임의의 i, j에 대해서 head2head[i][j] = 'W' 이면, head2head[j][i] = 'L'입니다.
  * 임의의 i, j에 대해서 head2head[i][j] = 'L' 이면, head2head[j][i] = 'W'입니다.
  * 임의의 i, j에 대해서 head2head[i][j] = 'N' 이면, head2head[j][i] = 'N'입니다.

우선 순위 정렬 조건은 다음과 같다.

1. 승률 내림차순
2. 승률 동일 시, 무게 내림차순
3. 무게 동일 시, 자신보다 무거운 선수를 이긴 횟수 내림차순
4. 자신보다 무거운 선수를 이긴 횟수 동일 시, 번호 오름차순

이 문제는 단순 구현이다. 각 선수의 번호, 무게, 승률, 자신보다 무거운 선수를 이긴 횟수에 대한 정보를 가진 목록을 만들어 문제의 요구사항에 맞게 정렬하면 된다. 그 후 정렬된 목록에서 선수들의 번호를 추출하여 리스트로 만들고 결과로 반환하면 끝이다.

주의할 점은 선수의 승률을 구할 때, 총 경기의 수는 `이긴 횟수(W) + 진 횟수(L)`로 구해야 한다. 아직 경기를 안한 횟수 (N)이 있기 때문이다.

```
승률 = 이긴 횟수 * 100 / (이긴 횟수 + 진 횟수)
```

## 코드

코드는 크게 다음으로 나눌 수 있다.

1. 선수 목록을 만든다. 각 선수는 번호, 무게, 승률, 자신보다 무거운 선수를 이긴 횟수에 대한 정보를 가지고 있다.
2. 선수 목록을 정렬한다. 다음 우선 순위 별로 정렬한다.
   1. 승률 내림차순
   2. 승률 동일 시, 무게 내림차순
   3. 무게 동일 시, 자신보다 무거운 선수를 이긴 횟수 내림차순
   4. 자신보다 무거운 선수를 이긴 횟수 동일 시, 번호 오름차순
3. 정렬된 선수 목록에서 번호만 뽑아 결과로 반환한다.

전체 코드는 다음과 같다.

```go
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
```