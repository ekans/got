package server

import (
	"fmt"
	"github.com/ekans/got/core"
	"log"
	"net/http"
	"path/filepath"
	"strings"
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

	if ok := core.CommandIsAllowed(w, P_CMD); !ok {
		return
	}

	if P_HOME_PATH != "" && P_CMD != "" {

		cmdByRepo := make(map[string]core.GitCmd)
		var repos []string

		if P_REPOS == "" {
			reposGlob, err := filepath.Glob(P_HOME_PATH + "/*/.git")
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			repos, err = core.CheckGitRepos(reposGlob)
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}

		} else {
			repos = strings.Split(P_REPOS, ",")
			for i, repo := range repos {
				repos[i] = P_HOME_PATH + "/" + repo
			}
		}

		for _, repo := range repos {
			cmdByRepo[repo] = core.GitCmd{Repo: repo, Cmd: []string{P_CMD}}
		}
		core.LaunchCmdAndWriteResult(w, cmdByRepo)

	} else {
		fmt.Fprintf(w, "Hum, what did you expect?")
	}
}
