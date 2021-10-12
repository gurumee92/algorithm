"""
프로그래머스(programmers.co.kr) 문제 풀이

코딩테스트 연습 > 위클리 챌린지 > 9주차_전략망을_둘로_나누기(https://programmers.co.kr/learn/courses/30/lessons/86971) 문제 풀이
"""
MAX = 100

def dfs(cur, n, is_visit, graph):
    is_rest = False

    for i in range(n):
        _next = graph[cur][i]

        if _next == 1 and not is_visit[_next]:
            is_rest = True

    if not is_rest:
        return is_visit.count(True) + 1
    
    is_visit[cur] = True
    cnt = 0

    for _next in graph[cur]:
        if _next == 1 and not is_visit[_next]:
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
    # 1. 그래프 초기화
    graph = [[0 for _ in range(n)] for _ in range(n)]
    
    for (_to, _from) in wires:
        _to -= 1
        _from -= 1
        graph[_to][_from] = 1
        graph[_from][_to] = 1

    
    answer = MAX

    for (_to, _from) in wires:
        _to -= 1
        _from -= 1
        graph[_to][_from] = 0
        graph[_from][_to] = 0
        cnt = 0

        for i in range(n):
            is_visit = [False for _ in range(n)]
            cnt = max(cnt, dfs(i, n, is_visit, graph))
        
        diff = abs((n - cnt) - cnt)
        answer = min(answer, diff)
        graph[_to][_from] = 1
        graph[_from][_to] = 1

    return answer
