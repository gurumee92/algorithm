# \[ALGORITHM\] 전력망을 둘로 나누기

## 문제

이 문서에서 다룰 문제 "전력망을 둘로 나누기"는 "프로그래머스 > 코딩테스트 연습 > 위클리 챌린지 > 9주차" 문제이다. 문제 링크는 다음과 같다.

* [https://programmers.co.kr/learn/courses/30/lessons/86971](https://programmers.co.kr/learn/courses/30/lessons/86971)

## 문제 풀이

먼저 문제의 입력은 다음과 같다.

* n:      int
* wires:  [][]int

제한 사항은 다음과 같다.

* 2 <= n <= 100 (즉 답은 100을 넘지 못한다.)
* len(wires) = n-1
* wires의 모든 원소는 하나의 트리를 이룬다. (모든 노드가 연결되어 있다.)

이 문제는 그래프로 해결할 수 있다.(시간은 오래 걸리지만) 먼저 그래프를 초기화한다.

```
graph = [[0 for _ in range(n)] for _ in range(n)]

# 노드 연결
for wire in wires:
    to = wire[0]
    from = wire[1]
    graph[to - 1][from - 1] = 1
    graph[from - 1][to - 1] = 1
```

그 다음 간선을 하나씩 제거해보면서 한 노드에서 도달할 수 있는 총 노드 개수를 구한다. 그 후 반대쪽 노드 개수를 구한 뒤 차이를 구하면 된다. 이 차이의 최소가 바로 정답이다. **이 때 중요한 것은 DFS가 끝난 후 제거한 간선을 복구하는 것이다.**

```
for wire in wires:
    to = wire[0]
    from = wire[1]

    # 간선 제거
    graph[to - 1][from - 1] = 0
    graph[from - 1][to - 1] = 0
    
    # 모든 노드에서 DFS를 돌린 후 최대 노드 개수 파악, 반대편 노드와의 차이가 최소인 것을 정답으로 구한다.
    for cur in range(n):
        is_visit = [False for _ in range(n)]    # DFS의 방문 여부를 확인하는 배열
        dfs(cur, n, is_visit, graph)            # DFS로 현재 노드에서 방문할 수 있는 모든 노드 조회
        a = is_visit.count(True)                # 방문한 노드 개수 
        b = n - a                               # 반대편 노드 개수 
        diff = abs(a - b)                       # a, b 차이
        answer = min(answer, diff)

    # 간선 복구
    graph[_to - 1][_from - 1] = 1  
    graph[_from - 1][_to - 1] = 1
```

DFS는 매우 쉽다. 모든 노드를 순회하게끔 하면 된다. 종료 조건은 더 이상 다른 노드를 방문할 수 없을 때까 종료 조건이다.

```
def dfs(cur, num, is_visit, graph):
    is_visit[cur] = True        # 방문 표시
    is_rest = False             # 종료 조건

    for _next in range(num):    # 순회할 수 있는 노드가 있는지 확인
        if graph[cur][_next] == 1 and not is_visit[_next]:
            is_rest = True
            break

    if not is_rest:             # 순회할 노드가 없으면 끝
        return 

    for _next in range(num):    # 순회할 노드가 있다면 순회
        if graph[cur][_next] == 1 and not is_visit[_next]:
            dfs(_next, num, is_visit, graph)
```

## 코드

전체 코드는 다음과 같다. 

```python
MAX = 100


def dfs(cur, num, is_visit, graph):
    is_visit[cur] = True
    is_rest = False

    for _next in range(num):
        if graph[cur][_next] == 1 and not is_visit[_next]:
            is_rest = True
            break

    if not is_rest:
        return

    for _next in range(num):
        if graph[cur][_next] == 1 and not is_visit[_next]:
            dfs(_next, num, is_visit, graph)


def solution(n, wires):
    graph = [[0 for _ in range(n)] for _ in range(n)]

    for (_to, _from) in wires:
        graph[_to - 1][_from - 1] = 1
        graph[_from - 1][_to - 1] = 1

    answer = MAX

    for (_to, _from) in wires:
        graph[_to - 1][_from - 1] = 0
        graph[_from - 1][_to - 1] = 0

        for cur in range(n):
            is_visit = [False for _ in range(n)]
            dfs(cur, n, is_visit, graph)
            a = is_visit.count(True)
            b = n - a
            diff = abs(a - b)
            answer = min(answer, diff)

        graph[_to - 1][_from - 1] = 1
        graph[_from - 1][_to - 1] = 1

    return answer
```

하지만 위의 코드의 경우 실행 속도가 매우 오래 걸린다. 정확도만 체크해서 그런지 통과는 했는데 만약 효율성 체크가 들어갔다면 통과 못했을 것 같다. 그래서 어떻게 할까 고민 좀 해봤는데 잘 모르겠더라. 결국 다른 사람이 푼 "유니온 파인드" 방식이 가장 간단해 보여서 추후 이걸 분석하는 방식으로 한 번 더 문제를 살펴 봐야겠다. 

```python
MAX = 100

def find(uf, a):
    if uf[a] < 0: 
        return a
    
    uf[a] = find(uf, uf[a])
    return uf[a]

def merge(uf, a, b):
    pa = find(uf, a)
    pb = find(uf, b)
    
    if pa == pb: 
        return
    
    uf[pa] += uf[pb]
    uf[pb] = pa

def solution(n, wires):
    answer = MAX
    k = len(wires)
    
    for i in range(k):
        uf = [-1 for _ in range(n+1)]
        tmp = [wires[x] for x in range(k) if x != i]
        
        for a, b in tmp: 
            merge(uf, a, b)
        
        v = [x for x in uf[1:] if x < 0]
        answer = min(answer, abs(v[0]-v[1]))

    return answer
```