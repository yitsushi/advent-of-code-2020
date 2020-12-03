package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/yitsushi/advent-of-code-2020/pkg/aoc"
	"github.com/yitsushi/advent-of-code-2020/pkg/extra"
	"github.com/yitsushi/advent-of-code-2020/pkg/puzzle"
)

func submitCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submit",
		Short: "Submit the solution",
		Run: func(cmd *cobra.Command, args []string) {
			dayNumber, _ := cmd.Flags().GetInt("day")
			partNumber, _ := cmd.Flags().GetInt("part")
			inputValue, _ := cmd.Flags().GetString("input")
			inputFile, _ := cmd.Flags().GetString("input-file")

			day, err := extra.DaySelector(dayNumber)
			if err != nil {
				logrus.Fatal(err.Error())

				return
			}

			solver := puzzle.NewSolver(day, partNumber)

			solution, err := solver.Solve(inputValue, inputFile)
			if err != nil {
				logrus.Fatalf("Unable to solve the puzzle: %s\n\n", err)

				return
			}

			valid, err := aoc.SubmitSolution(currentYear, dayNumber, partNumber, solution)
			if err != nil {
				fmt.Printf("Error :( => %s", err.Error())

				return
			}

			if valid {
				fmt.Println("Done \\o/")
			} else {
				fmt.Println("Something is wrong :(")
			}
		},
	}

	cmd.Flags().Int("day", 1, "Day")
	cmd.Flags().Int("part", 1, "Part")
	cmd.Flags().String("input", "", "Input as value")
	cmd.Flags().String("input-file", "", "Input as file (path)")

	_ = cmd.MarkFlagRequired("day")

	return cmd
}
