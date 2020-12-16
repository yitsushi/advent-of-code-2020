package parsing

import (
	"strconv"
	"strings"
)

// Int64Range parses a string as an int64 range.
func Int64Range(text, separator string) (int64, int64, error) {
	parts := strings.SplitN(text, "-", 2)

	from, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return 0, 0, err
	}

	to, err := strconv.ParseInt(parts[1], 10, 64)

	return from, to, err
}
