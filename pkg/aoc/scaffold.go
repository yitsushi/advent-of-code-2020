package aoc

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/sirupsen/logrus"
)

type dayTemplate struct {
	Day  int
	Root string
}

// Scaffold generates a Solution file directory structure for a day.
func Scaffold(day int, templateDir, packageRoot string) error {
	targetDir := fmt.Sprintf("days/day%02d", day)

	ensureDirectory(targetDir)

	files, err := ioutil.ReadDir(templateDir)
	if err != nil {
		return err
	}

	tmpl, err := template.ParseGlob(filepath.Join(templateDir, "*.tmpl"))
	if err != nil {
		return nil
	}

	templateVars := dayTemplate{
		Day:  day,
		Root: packageRoot,
	}

	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".tmpl") {
			continue
		}

		targetFile := strings.TrimSuffix(filepath.Join(targetDir, file.Name()), ".tmpl")

		f, err := os.Create(targetFile)
		if err != nil {
			logrus.Errorf("Skip %s because: %s", targetFile, err.Error())

			continue
		}

		defer f.Close()

		err = tmpl.ExecuteTemplate(f, file.Name(), templateVars)
		if err != nil {
			return err
		}
	}

	return nil
}
