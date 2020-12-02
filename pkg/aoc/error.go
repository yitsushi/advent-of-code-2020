package aoc

import "fmt"

// DownloadError occurs when the AoC server returns with a status code
// other than 200.
type DownloadError struct {
	StatusCode int
}

func (e DownloadError) Error() string {
	return fmt.Sprintf("Download error: %d", e.StatusCode)
}
