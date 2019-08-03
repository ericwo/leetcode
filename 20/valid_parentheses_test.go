package main

import (
	"testing"

	"github.com/theplant/testingutils/errorassert"
)

func TestIsValid(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		output bool
	}{
		{
			"input empty string",
			"",
			true,
		},
		{
			"input ()",
			"()",
			true,
		},
		{
			"input ()[]{}",
			"()[]{}",
			true,
		},
		{
			"input ()]",
			"()]",
			false,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			result := IsValid(test.input)
			errorassert.Equal(t, test.output, result, "result is wrong!")
		})
	}
}
