package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/alexflint/go-arg"
	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
)

func getPaths() {
	path, variableExists := os.LookupEnv("PATH")

	if variableExists {
		paths := filepath.SplitList(path)
		pathDataTable := table.NewWriter()
		duplicatePaths := make(map[string]int)

		pathDataTable.SetOutputMirror(os.Stdout)
		pathDataTable.AppendHeader(table.Row{"Path", "Exists", "Duplicate"})

		for _, path := range paths {
			var duplicateYesNo, existsYesNo string
			_, exist := duplicatePaths[path]

			if exist {
				duplicateYesNo = color.YellowString("Yes")
			} else {
				duplicateYesNo = "No"
				duplicatePaths[path] = 1
			}

			if _, err := os.Stat(path); err == nil {
				existsYesNo = "Yes"
			} else if os.IsNotExist(err) {
				path = color.RedString(path)
				existsYesNo = color.RedString("No")
			} else {
				existsYesNo = "Unknown"
			}

			pathDataTable.AppendRow(table.Row{path, existsYesNo, duplicateYesNo})
		}

		pathDataTable.SetStyle(table.StyleRounded)
		pathDataTable.Render()
	} else {
		fmt.Println("Error: PATH variable not found")
	}
}

func main() {
	var args struct {
		//Ugly bool `arg:"-u" default:"false" help:"Remove colorized output. Yes it's ugly."`
	}

	arg.MustParse(&args)
	getPaths()
}
