package core

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

func LaunchCmdAndWriteResult(w io.Writer, cmdByRepo map[string]GitCmd) {

	output := make(chan GitCmd, len(cmdByRepo))

	for _, gitCmd := range cmdByRepo {
		gitCmd.Exec(output)
	}

	for range cmdByRepo {
		fmt.Fprintln(w, <-output)
	}
}

func CheckGitRepos(repos []string, err error) ([]string, error) {
	if err != nil {
		return nil, err
	}

	if len(repos) == 0 {
		return nil, errors.New("No Git repos found in subfolders :-(")
	}

	for i, repo := range repos {
		repos[i] = strings.TrimSuffix(repo, "/.git")
	}

	return repos, nil
}
