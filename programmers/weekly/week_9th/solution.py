"""
프로그래머스(programmers.co.kr) 문제 풀이

코딩테스트 연습 > 위클리 챌린지 > 9주차_전략망을_둘로_나누기(https://programmers.co.kr/learn/courses/30/lessons/86971) 문제 풀이
"""
MAX = 100


def dfs(cur, n, is_visit, graph):
    is_visit[cur] = True
    cnt = is_visit.count(True)
    is_rest = False

    for _next in range(n):
        if graph[cur][_next] == 1 and not is_visit[_next]:
            is_rest = True
            break

    if not is_rest:
        return cnt

    for _next in range(n):
        if graph[cur][_next] == 1 and not is_visit[_next]:
            cnt = max(cnt, dfs(_next, n, is_visit, graph))

    return cnt


def solution(n, wires):
    """
    문제 풀이
    입력:
    * n:      int
    * sizes:  [][]int
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
