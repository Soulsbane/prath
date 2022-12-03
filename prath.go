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
			_, exist := duplicatePaths[path]
			duplicatePath := "No"

			if exist {
				duplicatePath = color.YellowString("Yes")
			} else {
				duplicatePath = "No"
				duplicatePaths[path] = 1
			}

			if _, err := os.Stat(path); err == nil {
				pathDataTable.AppendRow(table.Row{path, "Yes", duplicatePath})
			} else if os.IsNotExist(err) {
				pathDataTable.AppendRow(table.Row{color.RedString(path), color.RedString("No"), duplicatePath})
			} else {
				fmt.Println("Error: ", err)
			}
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
