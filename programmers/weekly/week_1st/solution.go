package week_1st

func arithmeticSequence(n int) int {
	return n * (n+1) / 2
}

func solution(price int, money int, count int) int64 {
	result := int64(price * arithmeticSequence(count) - money)

	if result > 0 {
		return result
	} else {
		return 0
	}
}
