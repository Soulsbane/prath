package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/alexflint/go-arg"
	termtables "github.com/brettski/go-termtables"
	"github.com/fatih/color"
)

// FIXME: Highlight duplicates
func getPaths() {
	path, variableExists := os.LookupEnv("PATH")

	if variableExists {
		paths := filepath.SplitList(path)
		table := termtables.CreateTable()

		table.AddHeaders("Path", "Exists")

		for _, path := range paths {
			if _, err := os.Stat(path); err == nil {
				table.AddRow(path, "Yes")
			} else if os.IsNotExist(err) {
				table.AddRow(color.RedString(path), color.RedString("No"))
			} else {
				fmt.Println("WHAT WENT WRONG")
			}
		}

		fmt.Println(table.Render())
	}
}

func main() {
	var args struct {
		//Ugly bool `arg:"-u" default:"false" help:"Remove colorized output. Yes it's ugly."`
	}

	arg.MustParse(&args)
	getPaths()
}
