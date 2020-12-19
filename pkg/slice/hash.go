package slice

import (
	"fmt"
	"strings"
)

func Hash(h []int64) string {
	stringValue := []string{}

	for _, v := range h {
		stringValue = append(stringValue, fmt.Sprintf("%d", v))
	}

	return strings.Join(stringValue, ",")
}
