"""
프로그래머스(programmers.co.kr) 문제 풀이 (테스트 코드)

코딩테스트 연습 > 위클리 챌린지 > 5주차_모음사전(https://programmers.co.kr/learn/courses/30/lessons/84512) 문제 풀이
"""
import pytest

from programmers.weekly.week_5th.solution import solution


TEST_DATA = [
    ("AAAAE", 6),
    ("AAAE", 10),
    ("I", 1563),
    ("EIO", 1189),
]


@pytest.mark.parametrize("word,expected", TEST_DATA)
def test_solution(word, expected):
    """
    solution에 대한 입/출력을 테스트한다.
    """
    result = solution(word)
    assert result == expected
