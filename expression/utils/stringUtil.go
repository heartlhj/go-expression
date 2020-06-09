package utils

func IndexOf(expressionString string, ch string, fromIndex int) int {
	max := len(expressionString)
	if fromIndex < 0 {
		fromIndex = 0
	}
	if fromIndex >= max {
		return -1
	}
	for i := fromIndex; i < max; i++ {
		if string(expressionString[i]) == ch {
			return i
		}
	}
	return -1
}

func BinarySearch(a []string, v interface{}) int {
	for i, i2 := range a {
		if i2 == v {
			return i
		}
	}
	return -1
}
