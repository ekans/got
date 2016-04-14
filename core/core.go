package core

import (
	"io"
	"fmt"
	"errors"
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
		repos[i] = strings.TrimRight(repo, "/.git")
	}

	return repos, nil
}

func GotMan(w io.Writer) {
	fmt.Fprintf(w, man)
}

const man = `NAME
       got [GLOBAL_COMMAND] [@REPO...]
       got [GLOBAL_COMMAND] [@REPO REPO_COMMAND]...

SYNOPSIS
       Execute commands across some Git repos in parallel.

GLOBAL_COMMAND | REPO_COMMAND
       Just a git command
OPTIONS
   @REPO
       A folder name containing a Git repo

EXAMPLES
       got status
          Run git status on all Git repos in the working dir

       got status @REPO1
          Run git status on all Git repos in the working dir except REPO1

       got @REPO1 status
          Run git status only on REPO1

       got checkout A_branch @REPO1 checkout Another_branch
          Run git checkout A_branch on all Git repos except REPO01 and git checkout Another_branch on REPO1
`

