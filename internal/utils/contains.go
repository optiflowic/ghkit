package utils

func Contains[T comparable](list []T, target T) bool {
	for _, s := range list {
		if s == target {
			return true
		}
	}
	return false
}
