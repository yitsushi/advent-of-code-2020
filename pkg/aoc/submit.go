package aoc

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

const (
	// CorrectAnswer is the goal.
	CorrectAnswer = "That's the right answer!"
	// AlreadySolved means our solution is already there.
	AlreadySolved = "Did you already complete it?"
)

func correctOrNot(content string) (bool, string) {
	valid := []string{
		CorrectAnswer,
		AlreadySolved,
	}

	for _, check := range valid {
		if strings.Contains(content, check) {
			return true, check
		}
	}

	return false, ""
}

// SubmitSolution downloads and saves the requested input file.
func SubmitSolution(year, day, part int, solution string) (bool, error) {
	sessionID := os.Getenv("AOC_SESSION")
	client := &http.Client{}

	form := url.Values{}
	form.Add("level", fmt.Sprintf("%d", part))
	form.Add("answer", solution)

	req, err := http.NewRequestWithContext(
		context.Background(),
		"POST",
		fmt.Sprintf("https://adventofcode.com/%d/day/%d/answer", year, day),
		strings.NewReader(form.Encode()),
	)
	if err != nil {
		return false, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.AddCookie(&http.Cookie{Name: "session", Value: sessionID})

	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, DownloadError{StatusCode: resp.StatusCode}
	}

	content, _ := ioutil.ReadAll(resp.Body)

	logrus.Debugf("%s", content)

	valid, text := correctOrNot(string(content))
	if !valid {
		err = IncorrectAnswer{Hint: "N/A"}
	}

	logrus.Info(text)

	return valid, err
}
