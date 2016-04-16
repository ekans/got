package main

import (
	"github.com/ekans/got/core"
	"github.com/ekans/got/server"
	"github.com/ekans/got/terminal"
	"os"
)

// Execute commands across some Git repos in parallel.
func main() {

	if ok := argsAreValid(os.Args); !ok {
		core.GotMan(os.Stdout)
		return
	}
	//Remove first argument which is "got"
	gotArgs := os.Args[1:]

	if gotArgs[0] == server.ServerOption {
		server.ServerMode()

	} else {
		terminal.TerminalMode(gotArgs)
	}
}

func argsAreValid(args []string) bool {
	if len(args) <= 1 {
		return false
	}
	return len(args) >= 2 || args[1] == server.ServerOption
}
