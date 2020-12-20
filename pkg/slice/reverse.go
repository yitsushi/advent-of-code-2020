package slice

// ReverseString returns with the reverse of a string.
func ReverseString(str string) string {
	result := []byte{}

	for idx := len(str) - 1; idx >= 0; idx-- {
		result = append(result, str[idx])
	}

	return string(result)
}
