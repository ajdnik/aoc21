package utils

func IsIncluded(data []int, itm int) bool {
	for _, i := range data {
		if i == itm {
			return true
		}
	}
	return false
}
