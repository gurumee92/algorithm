# \[ALGORITHM\] 부족한 금액 계산하기

## 문제

이 문서에서 다룰 문제 "부족한 금액 계산하기"는 "프로그래머스 > 코딩테스트 연습 > 위클리 챌린지 > 1주차" 문제이다. 문제 링크는 다음과 같다.

* [https://programmers.co.kr/learn/courses/30/lessons/82612?language=go](https://programmers.co.kr/learn/courses/30/lessons/82612?language=go)

## 문제 풀이

문제의 입력은 다음과 같다.

* price (1 <= x <= 2,500)
* money (1 <= x <= 1,000,000,000)
* count (1 <= x <= 2,500)

이용료는 `count`가 올라갈 때마다 `price`가 곱해진다.

```
1회차: 1 * price
2회차: 2 * price
...
n회차: n * price
```

주어진 `count`회차만큼의 총 이용료를 구하고 싶다면 위에 곱해진 모든 이용료를 더해주면 된다.

```
total = 1 * price + 2 * price + .. + count * price
```

이렇게 구한 `total`에서 `money`만큼 뺸 값이 바로 출력이 될 `result`이다.

```
result = total - money
```

즉 `result`의 수식은 다음과 같다.

```
result = total - money
       = 1 * price + 2 * price + .. + count * price - money
       # 이걸 표현하기 위해서 for문을 이용한다.
       = (1 + 2 + .. + count) * price - money
       # 조금 더 성능을 올리고 싶다면, 등차수열 공식을 이용한다.
       = (n * (n + 1) / 2) * price - money   
```

그래서 이 `result`가 0보다 크면 `result`를 아니라면 0을 반환하면 된다.

## 코드

코드는 다음과 같다.

solution.go
```go
// 등차 수열을 golang 으로 표현한 함수
func ArithmeticSequence(n int) int {
    return n * (n+1) / 2
}

// 해결 코드
func solution(price int, money int, count int) int64 {
    result := int64(price * ArithmeticSequence(count) - money)

    if result > 0 {
        return result
    } else {
        return 0
    }
}
```