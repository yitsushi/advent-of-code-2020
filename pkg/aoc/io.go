package aoc

import (
	"fmt"
	"os"
	"strings"
)

func ensurePath(path string) {
	parts := strings.Split(path, "/")

	path = "."

	for _, part := range parts {
		path = fmt.Sprintf("%s/%s", path, part)
		ensureDirectory(path)
	}
}

func ensureDirectory(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		_ = os.Mkdir(path, os.ModePerm)
	}
}
