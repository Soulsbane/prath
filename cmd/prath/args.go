package main

import (
	"fmt"

	"github.com/carlmjohnson/versioninfo"
)

type ProgramArgs struct {
}

func (args ProgramArgs) Description() string {
	return "Pretty prints a table of all directories found in $PATH environmental variable."
}

func (ProgramArgs) Version() string {
	return fmt.Sprintln("Version: ", versioninfo.Short())
}
