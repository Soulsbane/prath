package main

type ProgramArgs struct {
}

func (args ProgramArgs) Description() string {
	return "Pretty prints a table of all directories found in $PATH environmental variable."
}
