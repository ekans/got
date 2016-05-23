package core

import (
	"os"
	"testing"
)

func TestCheckCommand(t *testing.T) {
	cases := []testCheckCommand{
		{"", false},
		{"unTrucBidon", false},
		{"status", true},
	}
	for _, data := range cases {
		result := CheckCommand(os.Stdout, data.in)
		if result != data.result {
			t.Errorf("argsAreNotValid(%s) == %t, expected %t", data.in, result, data.result)
		}
	}
}

type testCheckCommand struct {
	in     string
	result bool
}
