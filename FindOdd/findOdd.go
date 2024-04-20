package kata

func FindOdd(seq []int) int {
	odd := 0
	for _, v := range seq {
		odd ^= v
	}
	return odd
}
