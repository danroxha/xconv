package util

func RemoveContains[T comparable](l []T, item T) []T {
	if len(l) == 0 {
		return l
	}

	for index, other := range l {
		if other == item {
			return append(l[:index], l[index+1:]...)
		}
	}
	return l
}

func RemoveIndex[T comparable](s []T, index int) []T {
	if len(s) == 0 {
		return s
	}
	return append(s[:index], s[index+1:]...)
}

func ContainsSlice[T comparable](slice []T, item T) bool {
	for _, currentItem := range slice {
		if currentItem == item {
			return true
		}
	}
	return false
}


func ReverseSlice[T comparable](s []T) []T {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
    s[i], s[j] = s[j], s[i]
  }
	return s
}
