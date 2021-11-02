package core

func sum(array []int) int {
	var result int
	for _, v := range array {
		result += v
	}
	return result
}

func mean(array []int) float64 {
	return float64(sum(array)) / float64(len(array))
}
