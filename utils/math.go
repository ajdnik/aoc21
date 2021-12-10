package utils

func Sum(data []int64) int64 {
	var sum int64
	for _, num := range data {
		sum += num
	}
	return sum
}

func SumUpTo(n int64) int64 {
	var sum int64
	for i := int64(1); i <= n; i++ {
		sum += i
	}
	return sum
}

func Mul(data []int64) int64 {
	mul := int64(1)
	for _, num := range data {
		mul *= num
	}
	return mul
}

func Add(data []int64, num int64) []int64 {
	res := []int64{}
	for _, n := range data {
		res = append(res, n+num)
	}
	return res
}
