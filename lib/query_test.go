package samplify_test

import (
	"testing"

	samplify "github.com/morningconsult/go-samplifyapi-client/lib"
)

func TestInSliceString(t *testing.T) {
	tables := []struct {
		TestCase string
		input    samplify.IntSlice
		expected string
	}{
		{
			"Case 1: happy path",
			samplify.IntSlice{127},
			"127",
		},
		{
			"Case 2: happy path 1",
			samplify.IntSlice{127, 128},
			"127,128",
		},
		{
			"Case 1: happy path",
			samplify.IntSlice{},
			"",
		},
	}

	for _, table := range tables {
		actual := table.input.String()
		if actual != table.expected {
			t.Fail()
		}
	}
}
