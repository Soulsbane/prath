package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/alexflint/go-arg"
	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
)

// FIXME: Highlight duplicates
func getPaths() {
	path, variableExists := os.LookupEnv("PATH")

	if variableExists {
		paths := filepath.SplitList(path)
		pathDataTable := table.NewWriter()

		pathDataTable.SetOutputMirror(os.Stdout)
		pathDataTable.AppendHeader(table.Row{"Path", "Exists"})

		for _, path := range paths {
			if _, err := os.Stat(path); err == nil {
				pathDataTable.AppendRow(table.Row{path, "Yes"})
			} else if os.IsNotExist(err) {
				pathDataTable.AppendRow(table.Row{color.RedString(path), color.RedString("No")})
			} else {
				fmt.Println("Error: ", err)
			}
		}
		pathDataTable.SetStyle(table.StyleRounded)
		pathDataTable.Render()
	}
}

func main() {
	var args struct {
		//Ugly bool `arg:"-u" default:"false" help:"Remove colorized output. Yes it's ugly."`
	}

	arg.MustParse(&args)
	getPaths()
}
