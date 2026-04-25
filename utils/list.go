package utils

func Unique[T comparable](data []T) []T {
	keys := map[T]bool{}
	uniq := []T{}
	for _, itm := range data {
		if _, ok := keys[itm]; !ok {
			keys[itm] = true
			uniq = append(uniq, itm)
		}
	}
	return uniq
}
