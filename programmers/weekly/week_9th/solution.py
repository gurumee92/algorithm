"""
프로그래머스(programmers.co.kr) 문제 풀이

코딩테스트 연습 > 위클리 챌린지 > 9주차_전략망을_둘로_나누기(https://programmers.co.kr/learn/courses/30/lessons/86971) 문제 풀이
"""
MAX = 100


def dfs(cur, num, is_visit, graph):
    """
    DFS로 시작 노드에서 연결된 모든 노드의 최대 개수를 반환한다.
    입력:
    * cur:     int
    * num:     int
    * is_visit []boolean
    * graph:   [][]int
    출력:
    * cnt: int
    """
    is_visit[cur] = True
    cnt = is_visit.count(True)
    is_rest = False

    for _next in range(num):
        if graph[cur][_next] == 1 and not is_visit[_next]:
            is_rest = True
            break

    if not is_rest:
        return cnt

    for _next in range(num):
        if graph[cur][_next] == 1 and not is_visit[_next]:
            cnt = max(cnt, dfs(_next, num, is_visit, graph))

    return cnt


def solution(n, wires):
    """
    문제 풀이
    1. 그래프 초기화 
    2. wires를 순회하여 1개씩 끊은 후 DFS로 연결된 개수 파악, 분할된 네트워크 차이의 최솟값을 구함.
    입력:
    * n:      int
    * wires:  [][]int
    출력:
    * answer: int
    """
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
            path_a = dfs(cur, n, is_visit, graph)
            path_b = n - path_a
            diff = abs(path_a - path_b)
            answer = min(answer, diff)

        graph[_to - 1][_from - 1] = 1
        graph[_from - 1][_to - 1] = 1

    return answer
