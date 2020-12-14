package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/yitsushi/advent-of-code-2020/pkg/extra"
	"github.com/yitsushi/advent-of-code-2020/pkg/puzzle"
)

func solveCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "solve",
		Short: "Solve a puzzle",
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

			fmt.Println(solution)
		},
	}

	cmd.Flags().Int("day", 1, "Day")
	cmd.Flags().Int("part", 1, "Part (0 => both)")
	cmd.Flags().String("input", "", "Input as value")
	cmd.Flags().String("input-file", "", "Input as file (path)")

	_ = cmd.MarkFlagRequired("day")

	return cmd
}
