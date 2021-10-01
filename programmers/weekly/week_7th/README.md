# \[ALGORITHM\] 입실 퇴실

## 문제

이 문서에서 다룰 문제 "입실 퇴실"는 "프로그래머스 > 코딩테스트 연습 > 위클리 챌린지 > 7주차" 문제이다. 문제 링크는 다음과 같다.

* [https://programmers.co.kr/learn/courses/30/lessons/86048](https://programmers.co.kr/learn/courses/30/lessons/86048)

## 문제 풀이

먼저 문제의 입력은 다음과 같다.

* enter []int 
* leave []int

제한 사항은 다음과 같다. 

* 1 ≤ enter의 길이 ≤ 1,000
* 1 ≤ enter의 원소 ≤ enter의 길이
  * 모든 사람의 번호가 중복없이 하나씩 들어있습니다.
* leave의 길이 = enter의 길이
* 1 ≤ leave의 원소 ≤ leave의 길이
  * 모든 사람의 번호가 중복없이 하나씩 들어있습니다.

이 문제의 포인트는 바로 "떠날 때까지 남아 있는다"이다. 가령 예를 들어보자. 입실 순서가 `[1, 4, 2, 3]`, 퇴실 순서가 `[2, 1, 3, 4]`라고 해보자. 먼저 입실 순서에 따라 1번이 들어간다.

```
enter: [4, 2, 3] # 1번이 들어갔다.
leave: [2, 1, 3, 4]
room: [1]
```

퇴실 순서에 2번이 제일 앞에 있다. 따라서 나가는 사람은 없다. 즉 2번이 들어올 때까지는 들어오기만 한다. 4번도 들어간다.

```
enter: [2, 3] # 4번이 들어갔다.
leave: [2, 1, 3, 4]
room: [1, 4]
```

이제 2번이 들어갈 차례이다. 근데 2번이 나갈 차례이기도 하다. 따라서 2번은 들어갔다 나온다.

```
enter: [3] # 2번이 들어갔다.
leave: [1, 3, 4] # 2번이 나갔다.
room: [1, 4]
result : [1, 2, 0, 1] # 1번 - 2번, 2번 - 4번 만남, 1번 1명, 4번 1명, 2번 2명 만났다.
```

이 때, 1번과 2번, 4번과 2번은 만나게 된 것으로 처리할 수 있다. 그 다음 1번이 퇴실 순서에 있으므로 나갈 수 있다.

```
enter: [3] 
leave: [3, 4] # 1번이 나갔다.
room: [4] # 1번이 나갔다.
result : [2, 2, 0, 2] # 1번 - 4번 만남, 1번 2명, 2번 2명, 4번 2명을 만났다.
```

이제 3번이 들어가고 나간다.

```
enter: [] # 3번이 들어갔다. 
leave: [4] # 3번이 나갔다.
room: [4] # 3번이 나갔다.
result : [2, 2, 1, 3] # 3번 - 4번 만남, 1번 2명, 2번 2명, 3번 1명, 4번 3명을 만났다.
```

이제 4번이 나간다.

```
enter: [] 
leave: [] # 4번이 나갔다. 
room: [] # 4번이 나갔다.
result : [2, 2, 1, 3] 
```

즉 `[2, 2, 1, 3]`이 반환되어야 한다. 이 문제의 핵심은 크게 2가지이다.

1. leave[0]가 enter[idx] 될 때까지 사람이 들어간다.
2. leave[0]가 room에 존재하지 않은 사람일 때까지 사람이 나간다.

`room`은 방에 상태, 즉 사람 번호에 따라 들어있는지 안들어와있는지 여부를 나타내기만 하면 된다.
## 코드

이 문제는 단순 구현이기 때문에 코드만 기술한다. 전체 코드는 다음과 같다. 

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