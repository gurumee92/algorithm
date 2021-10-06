"""
프로그래머스(programmers.co.kr) 문제 풀이

코딩테스트 연습 > 위클리 챌린지 > 8주차_최소직사각형(https://programmers.co.kr/learn/courses/30/lessons/86491) 문제 풀이
"""

MAX_LEN = 1000


def get_min_area(sizes):
    """
    입력:
    * sizes: [][]int
    출력:
    * area:  int (현재 sizes에 대한 최소 사각형 넓이)
    sizes를 순회하여 최대 너비, 높이를 구하여 넓이를 구한다.
    """
    width, height = 0, 0

    for size in sizes:
        width = max(width, size[0])
        height = max(height, size[1])

    return width * height


def dfs(current, end, sizes):
    """
    입력:
    * current:  int
    * end:      int
    * sizes:    [][]int
    출력:
    * area:     int
    current에 대해서 sizes[current]를 회전 한 것과 안 한 것을 비교. 즉 모든 원소를 순회하면서, 만들 수 있는 최소 직사각형 너비를 구한다.
    """
    if current == end:
        area = get_min_area(sizes)
        return area

    area = MAX_LEN * MAX_LEN
    area = min(area, dfs(current + 1, end, sizes))
    sizes = (
        sizes[:current]
        + [[sizes[current][1], sizes[current][0]]]
        + sizes[current + 1 :]
    )
    area = min(area, dfs(current + 1, end, sizes))
    return area


def solution(sizes):
    """
    문제 풀이
    입력:
    * sizes: [][]int
    출력:
    * answer: int
    """
    answer = dfs(0, len(sizes), sizes)
    return answer
