package main

import (
	"testing"
)

func TestArgsAreValid(t *testing.T) {
	cases := []testDataArgsAreValid{
		{[]string{""}, false},
		{[]string{"got", "st"}, true},
	}
	for _, data := range cases {
		result := argsAreValid(data.in)
		if result != data.result {
			t.Errorf("argsAreNotValid(%s) == %q, expected %q", data.in, result, data.result)
		}
	}
}

type testDataArgsAreValid struct {
	in     []string
	result bool
}
