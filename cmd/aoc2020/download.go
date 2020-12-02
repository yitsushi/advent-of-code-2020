package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/yitsushi/advent-of-code-2020/pkg/aoc"
)

func downloadCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "download",
		Short: "Download a puzzle input",
		Run: func(cmd *cobra.Command, args []string) {
			dayNumber, _ := cmd.Flags().GetInt("day")
			partNumber, _ := cmd.Flags().GetInt("part")

			output, err := aoc.DownloadInput(2020, dayNumber, partNumber)
			if err != nil {
				logrus.Fatal(err.Error())

				return
			}

			fmt.Printf("Done... %s\n", output)
		},
	}

	cmd.Flags().Int("day", 1, "Day")
	cmd.Flags().Int("part", 1, "Part")

	_ = cmd.MarkFlagRequired("day")

	return cmd
}
