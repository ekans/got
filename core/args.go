package core

import (
	"fmt"
	"io"
)

// Check if cmd is a allowed command
func CheckCommand(w io.Writer, cmd string) bool {

	switch cmd {
	case "status":
		return true
	case "fetch":
		return true
	case "checkout":
		return true
	case "log":
		return true
	default:
		fmt.Fprintf(w, "The command %v is not supported\n", cmd)
		return false
	}
}
