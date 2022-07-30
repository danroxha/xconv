package util

func RemoveContains[T comparable](l []T, item T) []T {
	for i, other := range l {
		if other == item {
			return append(l[:i], l[i+1:]...)
		}
	}
	return l
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
