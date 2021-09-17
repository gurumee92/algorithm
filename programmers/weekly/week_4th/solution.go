package week_4th

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
			languagePreferenceByJob[job][language] = POINT - idx
		}
	}

	// 2. 직업 별 총 점수 생성
	maxScore := 0
	scoreByJob := make(map[string]int)

	for job, languagePreference := range languagePreferenceByJob {
		score := 0

		for idx, language := range languages {
			preferScore := preference[idx]
			score += languagePreference[language] * preferScore
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
