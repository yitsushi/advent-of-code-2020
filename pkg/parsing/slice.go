package parsing

import (
	"strconv"
	"strings"
)

// Int64Slice parses a string as a slice of int64 values.
func Int64Slice(text, separator string) ([]int64, error) {
	list := []int64{}

	for _, value := range strings.Split(text, separator) {
		intValue, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return list, err
		}

		list = append(list, intValue)
	}

	return list, nil
}
