package utils

func Sum(data []int64) int64 {
	var sum int64
	for _, num := range data {
		sum += num
	}
	return sum
}
