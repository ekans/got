package terminal

import (
	"github.com/ekans/got/core"
	"reflect"
	"testing"
)

func TestParseArgs(t *testing.T) {
	cases := []testDataParseArgs{{
		in:        []string{"st"},
		globalCmd: core.GitCmd{Cmd: []string{"st"}},
	}}

	for _, data := range cases {
		globalCmd, specificCmds := parseArgs(data.in)
		if globalCmd.Repo != data.globalCmd.Repo || !reflect.DeepEqual(globalCmd.Cmd, data.globalCmd.Cmd) {
			t.Errorf("parseArgs(%s) == (%s, %s), expected (%s, %s)", data.in, globalCmd, specificCmds, data.globalCmd, data.specificCmds)
		}
		if len(specificCmds) != 0 {
			t.Errorf("0 specificCmds expected but %s received %s", len(specificCmds), specificCmds)
		}
	}

	cases = []testDataParseArgs{{
		in:           []string{"@REPO1", "st"},
		specificCmds: map[string]core.GitCmd{"@REPO1": core.GitCmd{Cmd: []string{"st"}, Repo: "@REPO1"}},
	}}

	for _, data := range cases {
		globalCmd, specificCmds := parseArgs(data.in)
		if globalCmd.Repo != "" || len(globalCmd.Cmd) != 0 {
			t.Errorf("parseArgs(%s) == (%s, %s), expected (%s, %s)", data.in, globalCmd, specificCmds, data.globalCmd, data.specificCmds)
		}
		if len(specificCmds) != 1 {
			t.Errorf("1 specificCms expected but %s received", len(specificCmds))
		}
	}
	cases = []testDataParseArgs{{
		in:        []string{"br -a", "@REPO1", "st -s", "@REPO2", "ll ..origin/master"},
		globalCmd: core.GitCmd{Cmd: []string{"br -a"}},
		specificCmds: map[string]core.GitCmd{
			"REPO1": core.GitCmd{Cmd: []string{"st -s"}, Repo: "REPO1"},
			"REPO2": core.GitCmd{Cmd: []string{"ll ..origin/master"}, Repo: "REPO2"},
		},
	}}

	for _, data := range cases {
		globalCmd, specificCmds := parseArgs(data.in)
		if globalCmd.Repo != data.globalCmd.Repo || !reflect.DeepEqual(globalCmd.Cmd, data.globalCmd.Cmd) {
			t.Errorf("parseArgs(%s) == (%s, %s), expected (%s, %s)", data.in, globalCmd, specificCmds, data.globalCmd, data.specificCmds)
		}
		if len(specificCmds) != 2 {
			t.Errorf("2 specificCmds expected but %s received", len(specificCmds))
		}
		for repo, gitCmd := range specificCmds {
			if gitCmd.Repo != data.specificCmds[repo].Repo || !reflect.DeepEqual(gitCmd.Cmd, data.specificCmds[repo].Cmd) {
				t.Errorf("parseArgs(%s) == (%s, %s), expected (%s, %s)", data.in, globalCmd, specificCmds, data.globalCmd, data.specificCmds)
			}
		}
	}
}

type testDataParseArgs struct {
	in           []string
	globalCmd    core.GitCmd
	specificCmds map[string]core.GitCmd
}
