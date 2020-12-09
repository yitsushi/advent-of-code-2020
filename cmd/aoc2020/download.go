package main

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/yitsushi/aoc"
)

func downloadCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "download",
		Short: "Download a puzzle input",
		Run: func(cmd *cobra.Command, args []string) {
			dayNumber, _ := cmd.Flags().GetInt("day")

			targetDir := fmt.Sprintf("input/day%02d", dayNumber)
			targetFile := fmt.Sprintf("%s/input", targetDir)

			ensurePath(targetDir)

			client := aoc.NewClient(os.Getenv("AOC_SESSION"))

			err := client.DownloadAndSaveInput(currentYear, dayNumber, targetFile)
			if err != nil {
				logrus.Fatal(err.Error())

				return
			}

			fmt.Printf("Done... %s\n", targetFile)
		},
	}

	cmd.Flags().Int("day", 1, "Day")

	_ = cmd.MarkFlagRequired("day")

	return cmd
}
