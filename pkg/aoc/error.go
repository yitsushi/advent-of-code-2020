package aoc

import "fmt"

type DownloadError struct {
	StatusCode int
}

func (e DownloadError) Error() string {
	return fmt.Sprintf("Download error: %d", e.StatusCode)
}
