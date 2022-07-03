package util

func RepeatString(s string, time int) string {
	current := ""

	for i := 0; i < time; i++ {
		current += s
	}

	return current
}