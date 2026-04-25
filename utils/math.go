package utils

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

func Sum[T Number](data []T) T {
	var sum T
	for _, num := range data {
		sum += num
	}
	return sum
}

func SumUpTo[T Number](n T) T {
	var sum, i T
	for i = 1; i <= n; i++ {
		sum += i
	}
	return sum
}

func Mul[T Number](data []T) T {
	mul := T(1)
	for _, num := range data {
		mul *= num
	}
	return mul
}

func Add[T Number](data []T, num T) []T {
	res := make([]T, len(data))
	for i, n := range data {
		res[i] = n + num
	}
	return res
}
