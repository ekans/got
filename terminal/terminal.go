package terminal

import (
	"fmt"
	"github.com/ekans/got/core"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func TerminalMode(gotArgs []string) {

	r, w := io.Pipe()

	// -F : Causes less to automatically exit if the entire file can be displayed on the first screen.
	// -R : Like -r, but only ANSI "color" escape sequences are output in "raw" form.
	// -S : Causes  lines  longer  than  the  screen width to be chopped rather than folded.
	// -X : Disables  sending  the  termcap initialization and deinitialization strings to the terminal.
	pager := exec.Command("less", "-FRSX")
	pager.Stdin = r
	pager.Stdout = os.Stdout
	pager.Stderr = os.Stderr

	if ok := core.CheckCommand(os.Stdout, gotArgs[0]); !ok {
		return
	}

	c := make(chan struct{})
	go func() {
		pager.Run()
		close(c)
	}()

	terminalHandler(w, gotArgs)

	w.Close()
	<-c
}

func terminalHandler(w io.Writer, gotArgs []string) {
	globalCmd, localCmdByRepo := parseTerminalArgs(gotArgs)

	repos, err := listGitRepos()
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	cmdByRepo := make(map[string]core.GitCmd)
	if len(globalCmd.Cmd) > 0 {
		for _, repo := range repos {
			cmdByRepo[repo] = core.GitCmd{Repo: repo, Cmd: globalCmd.Cmd}
		}
	}

	for repo, localCmd := range localCmdByRepo {
		cmdByRepo[repo] = localCmd
	}

	for repo, gitCmd := range cmdByRepo {
		if len(gitCmd.Cmd) <= 0 {
			delete(cmdByRepo, repo)
		}
	}

	core.LaunchCmdAndWriteResult(w, cmdByRepo)
}

func parseTerminalArgs(args []string) (globalCmd core.GitCmd, specificGitCmds map[string]core.GitCmd) {

	var posAt = -1
	for i, arg := range args {
		if strings.HasPrefix(arg, "@") {
			posAt = i
			break
		}
	}

	if posAt == -1 {
		globalCmd = core.GitCmd{Cmd: args[:]}
		return
	}

	globalCmd = core.GitCmd{Cmd: args[:posAt]}

	specificGitCmds = make(map[string]core.GitCmd)
	var currentRepo string
	for _, arg := range args[posAt:] {

		switch {
		case strings.HasPrefix(arg, "@"):
			currentRepo = strings.SplitAfter(arg, "@")[1]
			specificGitCmds[currentRepo] = core.GitCmd{Repo: currentRepo}

		default:
			if gitCmd, ok := specificGitCmds[currentRepo]; ok {
				gitCmd.Cmd = append(gitCmd.Cmd, arg)
				specificGitCmds[currentRepo] = gitCmd
			} else {
				//It's logically impossible because a GitCmd, with currentRepo as the key, has been added to the map previously
				panic("Impossible!!!!")
			}
		}
	}

	return globalCmd, specificGitCmds
}

func listGitRepos() ([]string, error) {
	return core.CheckGitRepos(filepath.Glob("*/.git"))
}
