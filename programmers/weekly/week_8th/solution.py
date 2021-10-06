"""
프로그래머스(programmers.co.kr) 문제 풀이

코딩테스트 연습 > 위클리 챌린지 > 8주차_최소직사각형(https://programmers.co.kr/learn/courses/30/lessons/86491) 문제 풀이
"""


def solution(sizes):
    """
    문제 풀이
    입력:
    * sizes: [][]int
    출력:
    * answer: int
    """
    sizes = sorted([sorted(size, reverse=True) for size in sizes], reverse=True)
    width, height = sizes[0]

    for _, tmp in sizes[1:]:
        height = max(height, tmp)

    answer = width * height
    return answer
