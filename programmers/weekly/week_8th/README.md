# \[ALGORITHM\] 최소직사각형

## 문제

이 문서에서 다룰 문제 "최소직사각형"는 "프로그래머스 > 코딩테스트 연습 > 위클리 챌린지 > 8주차" 문제이다. 문제 링크는 다음과 같다.

* [https://programmers.co.kr/learn/courses/30/lessons/86491](https://programmers.co.kr/learn/courses/30/lessons/86491)

## 문제 풀이

## 코드

전체 코드는 다음과 같다.

```python
def solution(sizes):
    sizes = sorted([sorted(size, reverse=True) for size in sizes], reverse=True)
    width, height = sizes[0]

    for _, tmp in sizes[1:]:
        height = max(height, tmp)

    answer = width * height
    return answer
```