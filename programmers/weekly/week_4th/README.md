# \[ALGORITHM\] 직업군 추천하기

## 문제

이 문서에서 다룰 문제 "직업군 추천하기"는 "프로그래머스 > 코딩테스트 연습 > 위클리 챌린지 > 4주차" 문제이다. 문제 링크는 다음과 같다.

* [https://programmers.co.kr/learn/courses/30/lessons/84325](https://programmers.co.kr/learn/courses/30/lessons/84325)

## 문제 풀이

먼저 문제의 입력은 다음과 같다.

* table []string
* languages []string
* preference []int

`table`은 문자열 배열로써 하나의 원소는 다음과 같다.

```
SI JAVA JAVASCRIPT SQL PYTHON C#
```

이는 다음과 같은 형태로 분해할 수 있다.

```
[직업군] [5점 언어] [4점 언어] [3점 언어] [2점 언어] [1점 언어]
```

문제를 푸려면 이 문자열을 통해서 직업 군마다 각 언어에 대한 점수를 가진 데이터로 변환해주어야 한다. 어떻게 할 수 있을까? 먼저 `table`의 원소를 `" "`로 쪼갠다. 

```
SI          # [직업군] 
JAVA        # [5점 언어] 
JAVASCRIPT  # [4점 언어] 
SQL         # [3점 언어] 
PYTHON      # [2점 언어] 
C#          # [1점 언어]
```

그 후 "언어"를 키로 값으로 "점수"를 갖는 맵을 생성한 후, 키로 "직업군"을 값으로 방금 생성한 맵을 갖는 맵을 생성하면 된다. 즉 다음과 같은 형태로 데이터를 초기화해주면 된다.

```
{ 직업군: { 언어: 점수 } }
```

예를 들면 이런 식으로 말이다.

```
SI : JAVA: 5
     JAVASCRIPT: 4
     SQL: 3
     PYTHON: 2
     C#: 1

...
```

그 후 이렇게 생성된 맵을 순회하고, `languages`와 `preference`을 이용하여 각 직업군의 점수를 뽑아내면 된다. 문제에서 최고 점수를 가진 직업군을 뽑아내는 것이 목표이므로 **최대 점수도 이 때 갱신**해주어야 한다. 

```
languages  = ["PYTHON", "C++, "SQL"]
preference = [7,        5,    5]

# [직업군]: [점수] (preference[i] * m[직업군][languages[i]])
SI: 29 (=7 x 2(PYTHON) + 5 x 0(C++) + 5 x 3(SQL))
CONTENTS: 36
HARDWARE: 41
PORTAL: 21
GAME: 25

maxScore: 41
```

그 후 방금 생성한 키를 직업군으로, 값을 점수로 갖는 이 맵을 토대로 최고 점수만 가진 직업군들을 추려내면 된다. 

```
maxScore: 41

# maxScore와 동일한 점수를 지닌 직업군
["HARDWARE"]
```

이 최고 점수를 가진 직업군은 여러 개가 될 수 있으므로, 위에 생성한 리스트를 사전순으로 정렬한 후 첫 원소를 반환해주면 된다.

## 코드

코드는 크게 4가지 부분으로 나뉜다.

1. 직업 별로 언어 별 점수가 저장된 맵을 생성한다.
2. (1)을 통해 생성한 맵을 순회해서 직업별 총 점수를 저장하는 맵을 생성한다.
3. (2)를 통해 생성한 맵을 순회하여 최대 점수를 가진 직업군들을 필터링하여 리스트를 생성한다.
4. (3)을 통해 필터링된 직업군 리스트를 사전순으로 정렬한다. 그 후 첫 원소를 반환한다.

먼저 `{직업: {언어: 점수}}` 형태의 맵을 생성하기 위해서 다음과 같이 작성한다.

```go
const POINT = 5
languagePreferenceByJob := make(map[string]map[string]int)

for _, str := range table {
    row := strings.Split(str, " ")                          // " "로 table의 한 원소를 쪼갠다.
    job := row[0]                                           // 이렇게 쪼개진 row의 첫 원소는 직업군 job을 나타낸다..
    preferredLanguages := row[1:]                           // 나머지는 각 언어에 대한 점수이다.
    languagePreferenceByJob[job] = make(map[string]int)

    for idx, language := range preferredLanguages {
        languagePreferenceByJob[job][language] = (POINT - idx) // 역순으로 점수를 매겨야 한다. 첫 원소는 인덱스는 0, 점수는 5점, 두번째 원소는 인덱스는 1 점수는 4점 ...
    }
}
```

이제 `{직업: 총 점수}` 형태의 맵을 생성하기 위해서 다음과 같이 코드를 작성한다. 이 때 위에 생성한 `languagePreferenceByJob`를 순회해야 하는데, 최고 점수도 갱신해준다.

```go
maxScore := 0
scoreByJob := make(map[string]int)

for job, languagePreference := range languagePreferenceByJob {
    // 각 직업군 별 점수 합산
    // score = score + (languagePreference[job][languages[i]] * preferences[i])
    score := 0

    for idx, language := range languages {
        preferScore := preference[idx]
        score += (languagePreference[language] * preferScore)
    }

    // {직업: 총 점수} 맵의 키-값 할당
    scoreByJob[job] = score
    
    // 최대 점수 갱신
    if maxScore <= score {
        maxScore = score
    }
}
```

그 후 `maxScore`를 갖는 직업군을 추려낸다.

```go
results := make([]string, 0)

for job, score := range scoreByJob {
    if maxScore == score {
        results = append(results, job)
    }
}
```

그리고 이렇게 얻은 리스트 `results`를 사전 순으로 정렬한 후, 첫 원소를 반환하면 된다.

```go
sort.Slice(results, func(i, j int) bool {
    return results[i] < results[j]
})

return results[0]
```

끝이다. 전체 코드는 다음과 같다. 

```go
import (
	"sort"
	"strings"
)

const POINT = 5

func solution(table []string, languages []string, preference []int) string {
	// 1. 직업 별, 언어 별 점수 생성
	languagePreferenceByJob := make(map[string]map[string]int)

	for _, str := range table {
		row := strings.Split(str, " ")
		job := row[0]
		preferredLanguages := row[1:]
		languagePreferenceByJob[job] = make(map[string]int)

		for idx, language := range preferredLanguages {
			languagePreferenceByJob[job][language] = (POINT - idx)
		}
	}

	// 2. 직업 별 총 점수 생성
	maxScore := 0
	scoreByJob := make(map[string]int)

	for job, languagePreference := range languagePreferenceByJob {
		score := 0

		for idx, language := range languages {
			preferScore := preference[idx]
			score += (languagePreference[language] * preferScore)
		}

		scoreByJob[job] = score

		if maxScore <= score {
			maxScore = score
		}
	}

	// 3. 최대 점수를 가진 직업군 판별
	results := make([]string, 0)

	for job, score := range scoreByJob {
		if maxScore == score {
			results = append(results, job)
		}
	}

	// 4. 사전순으로 정렬
	sort.Slice(results, func(i, j int) bool {
		return results[i] < results[j]
	})

	return results[0]
}
```