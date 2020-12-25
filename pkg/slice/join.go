package slice

import (
	"strconv"
	"strings"
)

// JoinInt produces a string with all the int values in the provided
// slice separated by 'separator'.
func JoinInt(list []int, separator string) string {
	output := []string{}

	for _, item := range list {
		output = append(output, strconv.Itoa(item))
	}

	return strings.Join(output, separator)
}
