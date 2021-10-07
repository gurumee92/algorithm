# \[ALGORITHM\] 최소직사각형

## 문제

이 문서에서 다룰 문제 "최소직사각형"는 "프로그래머스 > 코딩테스트 연습 > 위클리 챌린지 > 8주차" 문제이다. 문제 링크는 다음과 같다.

* [https://programmers.co.kr/learn/courses/30/lessons/86491](https://programmers.co.kr/learn/courses/30/lessons/86491)

## 문제 풀이

먼저 문제의 입력은 다음과 같다.

* sizes [][]int 

이 문제는 매우 간단하다. 첫 번째 입력을 한 번 살펴보자.

```
[60, 50]
[30, 70]
[60, 30]
[80, 40]
```

이 때 첫 번째 원소가 가로, 두 번째 원소가 세로일 때 가로가 최대가 되게끔 재 배열한다.

```
[60, 50]
[70, 30]
[60 ,30]
[80, 40]
```

그럼 최대 가로 길이는 `80`이다. 그 후 최대 세로 길이를 구하면 된다. 최대 세로 길이는 `50`이다. 따라서 답은 `80 * 50 = 4000`이 된다. 단순히 재 정렬해서 풀 수도 있지만, 조금만 더 생각해보자. 이 문제는 `Max Heap` 2개를 사용하면 매우 깔끔하게 풀 수 있다. 

가로를 저장하는 힙을 `heap_width` 세로를 저장하는 힙을 `heap_height`라 하자

```
heap_width = []
heap_height = []
```

먼저 `sizes`를 순회한다. 순회할 때 첫번째 원소와 두번째 원소 중 큰 것이 가로로, 작은 것이 세로로 가면 된다. 각 순회 시, 힙 상황은 다음과 같다.

순회 첫 번째
```
sizes[0] = [60, 50]
heap_width = [60]
heap_height = [50]
```

순회 두 번째
```
sizes[1] = [30, 70]
heap_width = [70, 60]
heap_height = [50, 30]
```

순회 세 번째
```
sizes[2] = [60, 30]
heap_width = [70, 60, 60]
heap_height = [50, 30, 30]

```

순회 네 번째
```
sizes[3] = [80, 40]
heap_width = [80, 70, 60, 60]
heap_height = [50, 40, 30, 30]
```

그 후 힙의 첫 원소를 꺼내서 곱하면 된다.

```
# heappop은 힙의 첫 원소를 꺼낸다.
width = heappop(heap_width)     # 80
height = heappop(heap_height)   # 50
answer = width * height         # 4000
```
## 코드

위의 내용을 구현만 하면 된다. 전체 코드는 다음과 같다. 다만 `Python`의 `heapq`는 최소 힙이기 때문에 -를 붙여 저장하면 최대 힙이 된다.

```python
import heapq


def solution(sizes):
    heap_width = []
    heap_height = []

    for width, height in sizes:
        heapq.heappush(heap_width, -max(width, height))
        heapq.heappush(heap_height, -min(width, height))

    answer = heapq.heappop(heap_width) * heapq.heappop(heap_height)
    return answer
```