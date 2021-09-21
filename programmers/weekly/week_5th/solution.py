"""
프로그래머스(programmers.co.kr) 문제 풀이

코딩테스트 연습 > 위클리 챌린지 > 5주차_모음사전(https://programmers.co.kr/learn/courses/30/lessons/84512) 문제 풀이
"""
LETTERS = ["A", "E", "I", "O", "U"]


def make_dict(curr, _dict, length):
    """
    AEIOU로 만들 수 있는 단어들의 사전을 만든다.
    현재 단어 curr에 LETTERS에 존재하는 알파벳을 하나씩 붙인 단어들을 _dict에 차례로 추가한다.
    """
    if length == 5:
        return

    for letter in LETTERS:
        _next = curr + letter
        _dict.append(_next)
        make_dict(_next, _dict, length + 1)


def solution(word):
    """
    문제 풀이
    입력: word: str
    출력: idx: int (사전에서 word의 index)
    제한 사항
    - 1 <= len(word) <= 5
    - word는 A, E, I, O, U로만 이루어져 있다.
    """
    _dict = []

    for letter in LETTERS:
        _dict.append(letter)
        make_dict(letter, _dict, 1)

    answer = sorted(list(_dict)).index(word) + 1
    return answer
