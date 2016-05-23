package core

import (
	"fmt"
	"io"
)

func CheckCommand(w io.Writer, cmd string) bool {

	switch cmd {
	case "status":
		return true
	default:
		fmt.Fprintf(w, "The command %v is not supported\n", cmd)
		return false
	}
}
