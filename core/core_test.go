package core

import (
	"os"
	"reflect"
	"testing"
)

func TestCheckGitRepos(t *testing.T) {
	cases := []testReposAreValid{
		{[]string{"repo1/.git"}, []string{"repo1"}},
		{[]string{"bash-git-prompt/.git"}, []string{"bash-git-prompt"}},
	}
	for _, data := range cases {
		result, err := CheckGitRepos(data.repos)
		if err != nil {
			t.Errorf("Failed to check Git repos (%q) == %q, expected %q", data.repos, result, data.result)
		}
		if !reflect.DeepEqual(result, data.result) {
			t.Errorf("Failed to check Git repos (%q) == %q, expected %q", data.repos, result, data.result)
		}
	}
}

type testReposAreValid struct {
	repos  []string
	result []string
}

func TestCheckCommand(t *testing.T) {
	cases := []testCheckCommand{
		{"", false},
		{"unTrucBidon", false},
		{"status", true},
	}
	for _, data := range cases {
		result := CommandIsAllowed(os.Stdout, data.in)
		if result != data.result {
			t.Errorf("argsAreNotValid(%s) == %t, expected %t", data.in, result, data.result)
		}
	}
}

type testCheckCommand struct {
	in     string
	result bool
}
