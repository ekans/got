package core

import (
	"os/exec"
	"fmt"
)

type GitCmd struct {
	Cmd    []string
	Repo   string
	output []byte
	err    error
}

func (gitCmd GitCmd) Exec(output chan GitCmd) {
	go func() {
		cmd := exec.Command("git", gitCmd.Cmd...)
		cmd.Dir = gitCmd.Repo
		gitCmd.output, gitCmd.err = cmd.CombinedOutput()
		output <- gitCmd
	}()
}

func (gitCmd GitCmd) String() string {
	var result string

	result += fmt.Sprintf(">>> %s in %s", gitCmd.Cmd, gitCmd.Repo)
	if gitCmd.err != nil {
		result += fmt.Sprintf("\nCommand failed with %s", gitCmd.err.Error())
	}

	if len(gitCmd.output) > 0 {
		result += fmt.Sprintf("\n%s", gitCmd.output)
	}

	return result
}
