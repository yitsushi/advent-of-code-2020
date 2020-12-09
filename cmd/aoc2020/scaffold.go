package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/yitsushi/aoc"
)

const (
	packageRoot = "github.com/yitsushi/advent-of-code-2020"
)

type templateVariables struct {
	Day  int
	Root string
}

func scaffoldCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "scaffold",
		Short: "Generate a Day from template",
		Run: func(cmd *cobra.Command, args []string) {
			dayNumber, _ := cmd.Flags().GetInt("day")
			templateDir, _ := cmd.Flags().GetString("template-dir")

			if templateDir == "" {
				templateDir = "template/day"
			}

			err := aoc.Scaffold(
				templateDir,
				fmt.Sprintf("days/day%02d", dayNumber),
				templateVariables{
					Day:  dayNumber,
					Root: packageRoot,
				},
			)
			if err != nil {
				logrus.Errorln(err)
			}
		},
	}

	cmd.Flags().Int("day", 1, "Day")
	cmd.Flags().String("template-dir", "", "Template Directory")

	_ = cmd.MarkFlagRequired("day")

	return cmd
}
