package day18

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

type calculatorFunction func(string) int64

func removeParentheses(str string, calculator calculatorFunction) string {
	findBlock := regexp.MustCompile(`\([^\(\)]+\)`)

	for {
		matches := findBlock.FindAllString(str, -1)
		if len(matches) == 0 {
			break
		}

		for _, match := range matches {
			str = strings.ReplaceAll(
				str,
				match,
				fmt.Sprintf("%d", calculator(match[1:len(match)-1])))
		}
	}

	return str
}

func removeAddition(str string) string {
	findBlock := regexp.MustCompile(`\d+( \+ \d+)+`)

	for {
		matches := findBlock.FindAllString(str, -1)
		if len(matches) == 0 {
			break
		}

		for _, match := range matches {
			str = strings.ReplaceAll(
				str,
				match,
				fmt.Sprintf("%d", calculateSimple(match)))
		}
	}

	return str
}

func calculateSimple(str string) int64 {
	parts := strings.Split(str, " ")

	var (
		result int64
		op     string = "+"
	)

	for _, part := range parts {
		value, err := strconv.ParseInt(part, 10, 54)
		if err != nil {
			op = part

			continue
		}

		switch op {
		case "+":
			result += value
		case "*":
			result *= value
		default:
			logrus.Info(value, err, op)

			panic("no op")
		}
	}

	return result
}

func calculateAdvanced(str string) int64 {
	parts := strings.Split(removeAddition(str), " ")

	var (
		result int64
		op     string = "+"
	)

	for _, part := range parts {
		value, err := strconv.ParseInt(part, 10, 54)
		if err != nil {
			op = part

			continue
		}

		switch op {
		case "+":
			result += value
		case "*":
			result *= value
		default:
			logrus.Info(value, err, op)

			panic("no op")
		}
	}

	return result
}
