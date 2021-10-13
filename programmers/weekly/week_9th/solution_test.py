"""
프로그래머스(programmers.co.kr) 문제 풀이 (테스트 코드)

코딩테스트 연습 > 위클리 챌린지 > 9주차_전략망을_둘로_나누기(https://programmers.co.kr/learn/courses/30/lessons/86971) 문제 풀이
"""
import pytest

from programmers.weekly.week_9th.solution import solution


TEST_DATA = [
    (9, [[1, 3], [2, 3], [3, 4], [4, 5], [4, 6], [4, 7], [7, 8], [7, 9]], 3),
    (4, [[1, 2], [2, 3], [3, 4]], 0),
    (7, [[1, 2], [2, 7], [3, 7], [3, 4], [4, 5], [6, 7]], 1),
]


@pytest.mark.parametrize("n,wires,expected", TEST_DATA)
def test_solution(n, wires, expected):
    """
    solution에 대한 입/출력을 테스트한다.
    """
    result = solution(n, wires)
    assert result == expected
