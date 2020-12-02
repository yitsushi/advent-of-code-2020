package aoc

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func DownloadInput(year, day, part int) (string, error) {
	sessionID := os.Getenv("AOC_SESSION")
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.AddCookie(&http.Cookie{Name: "session", Value: sessionID})

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", DownloadError{StatusCode: resp.StatusCode}
	}

	targetDir := fmt.Sprintf("input/day%02d", day)
	targetFile := fmt.Sprintf("%s/part%d", targetDir, part)

	ensurePath(targetDir)

	f, err := os.Create(targetFile)
	if err != nil {
		return "", err
	}

	defer f.Close()

	_, err = io.Copy(f, resp.Body)

	return targetFile, err
}

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
		os.Mkdir(path, os.ModePerm)
	}
}
