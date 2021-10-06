"""
프로그래머스(programmers.co.kr) 문제 풀이 (테스트 코드)

코딩테스트 연습 > 위클리 챌린지 > 8주차_최소직사각형(https://programmers.co.kr/learn/courses/30/lessons/86491) 문제 풀이
"""
import pytest

from programmers.weekly.week_8th.solution import solution


TEST_DATA = [
    ([[60, 50], [30, 70], [60, 30], [80, 40]], 4000),
    ([[10, 7], [12, 3], [8, 15], [14, 7], [5, 15]], 120),
    ([[14, 4], [19, 6], [6, 16], [18, 7], [7, 11]], 133),
]


@pytest.mark.parametrize("sizes,expected", TEST_DATA)
def test_solution(sizes, expected):
    """
    solution에 대한 입/출력을 테스트한다.
    """
    result = solution(sizes)
    assert result == expected
