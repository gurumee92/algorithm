"""
프로그래머스(programmers.co.kr) 문제 풀이

코딩테스트 연습 > 위클리 챌린지 > 8주차_최소직사각형(https://programmers.co.kr/learn/courses/30/lessons/86491) 문제 풀이
"""
import heapq


def solution(sizes):
    """
    문제 풀이
    입력:
    * sizes: [][]int
    출력:
    * answer: int
    """
    heap_width = []
    heap_hight = []

    for width, height in sizes:
        heapq.heappush(heap_width, -max(width, height))
        heapq.heappush(heap_hight, -min(width, height))

    answer = heapq.heappop(heap_width) * heapq.heappop(heap_hight)
    return answer
