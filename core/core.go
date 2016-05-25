package core

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

// Launch some GitCmd and write result on the io.Writer in parameter
func LaunchCmdAndWriteResult(w io.Writer, cmdByRepo map[string]GitCmd) {

	output := make(chan GitCmd, len(cmdByRepo))

	for _, gitCmd := range cmdByRepo {
		gitCmd.Exec(output)
	}

	for range cmdByRepo {
		fmt.Fprintln(w, <-output)
	}
}

// Retrieve paths of git repos from a pattern given in parameter
func CheckGitRepos(repos []string) ([]string, error) {

	if len(repos) == 0 {
		return nil, errors.New("No Git repos found in subfolders :-(")
	}

	for i, repo := range repos {
		repos[i] = strings.TrimSuffix(repo, "/.git")
	}

	return repos, nil
}

// Check if cmd is a allowed command
func CommandIsAllowed(w io.Writer, cmd string) bool {

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
