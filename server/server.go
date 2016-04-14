package server

import (
	"net/http"
	"log"
	"github.com/ekans/got/core"
	"path/filepath"
	"strings"
	"fmt"
)

const ServerOption string = "--server"

func ServerMode() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {

	var params = r.URL.Query()
	P_HOME_PATH := params.Get("home")
	P_CMD := params.Get("cmd")
	P_REPOS := params.Get("repos")

	if P_HOME_PATH != "" && P_CMD != "" {

		cmdByRepo := make(map[string]core.GitCmd)
		var repos []string

		if P_REPOS == "" {
			repos, _ = core.CheckGitRepos(filepath.Glob(P_HOME_PATH + "/*/.git"))

		} else {
			repos = strings.Split(P_REPOS, ",")
			for i, repo := range repos {
				repos[i] = P_HOME_PATH + "/" + repo
			}
		}

		for _, repo := range repos {
			cmdByRepo[repo] = core.GitCmd{Repo:repo, Cmd:[]string{P_CMD}}
		}
		core.LaunchCmdAndWriteResult(w, cmdByRepo)

	} else {
		fmt.Fprintf(w, "Hum, what did you expect?")
	}
}
