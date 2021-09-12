package pkg

import "log"

func Sum(array []int) int {
	var result int
	for _, v := range array {
		result += v
	}
	return result
}

func Mean(array []int) float64 {
	return float64(Sum(array)) / float64(len(array))
}

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
