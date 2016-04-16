package core

import (
	"reflect"
	"testing"
)

func TestCheckGitRepos(t *testing.T) {
	cases := []testReposAreValid{
		{[]string{"repo1"}, []string{"repo1"}},
		{[]string{"bash-git-prompt"}, []string{"bash-git-prompt"}},
	}
	for _, data := range cases {
		result, err := CheckGitRepos(data.repos, nil)
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
