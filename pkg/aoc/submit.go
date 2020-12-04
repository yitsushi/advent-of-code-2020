package aoc

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
)

const (
	// CorrectAnswer is the goal.
	CorrectAnswer = "That's the right answer!"
	// AlreadySolved means our solution is already there.
	AlreadySolved = "Did you already complete it?"
	// WrongAnswer means the solution was incorrect.
	WrongAnswer = "That's not the right answer;"
	// WaitMore means you have to wait more.
	WaitMore = "You gave an answer too recently;"
)

func correctOrNot(content string) (bool, error) {
	valid := []string{
		CorrectAnswer,
		AlreadySolved,
	}

	for _, check := range valid {
		if strings.Contains(content, check) {
			return true, nil
		}
	}

	if strings.Contains(content, WrongAnswer) {
		re := regexp.MustCompile(
			`That's not the right answer; (?P<Hint>[^\.]+)..*(?P<Wait>Please wait [^.]+).`,
		)

		result := re.FindStringSubmatch(content)

		return false, IncorrectAnswer{
			Hint: result[1],
			Wait: result[2],
		}
	}

	if strings.Contains(content, WaitMore) {
		re := regexp.MustCompile(
			`You gave an answer too recently; you .* again. (?P<Wait>[^\.]+).`,
		)

		result := re.FindStringSubmatch(content)

		return false, IncorrectAnswer{
			Hint: "You gave an answer too recently",
			Wait: result[1],
		}
	}

	re := regexp.MustCompile(`<article><p>(?P<Content>.*)</p></article>`)
	result := re.FindStringSubmatch(content)

	return false, IncorrectAnswer{
		Hint: result[1],
		Wait: "",
	}
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

	return correctOrNot(string(content))
}
