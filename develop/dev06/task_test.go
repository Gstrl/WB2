package main

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCut_Run(t *testing.T) {
	testTable := []struct {
		cut      Cut
		line     string
		expected string
	}{
		{
			cut: Cut{
				fields:    []int{1, 2},
				delimiter: " ",
				separated: false,
			},
			line:     "fhf dsj jsd ds",
			expected: "fhf dsj",
		},
		{
			cut: Cut{
				fields:    []int{1, 2},
				delimiter: ":",
				separated: false,
			},
			line:     "fhf dsj:jsd:ds",
			expected: "fhf dsj:jsd",
		},
		{
			cut: Cut{
				fields:    []int{1},
				delimiter: "_",
				separated: true,
			},
			line:     "fhf dsj:jsd:ds",
			expected: "",
		},
		{
			cut: Cut{
				fields:    []int{1},
				delimiter: "_",
				separated: false,
			},
			line:     "fhf dsj:jsd:ds",
			expected: "fhf dsj:jsd:ds",
		},
	}

	for _, v := range testTable {
		res := v.cut.Run(v.line)
		t.Logf("Calling cut.Run(%v), resault %s\n", v.line, res)
		assert.Equal(t, v.expected, res, fmt.Sprintf("Incorrect resault. Expect %s, got %s", v.expected, res))
	}
}

func TestCut_FromArgs(t *testing.T) {
	var cut Cut
	err := cut.FromArgs()
	assert.Equal(t, errors.New("flag F is not set or set incorrectly"), err, fmt.Sprint("Incorrect err"))
}
