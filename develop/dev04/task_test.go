package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMakeSet(t *testing.T) {

	var TestTable = []struct {
		arrStr   []string
		expected string
	}{
		{
			arrStr:   []string{"пятак", "лиСток", "пЯтка", "слиток", "тяпка"},
			expected: "map[листок:[листок слиток] пятак:[пятак пятка тяпка]]",
		},
		{
			arrStr:   []string{},
			expected: "map[]",
		},
		{
			arrStr:   []string{"сос", "кек", "ссо", "осс", "лол", "со", "a"},
			expected: "map[сос:[сос ссо осс]]",
		},
	}

	for _, v := range TestTable {
		res := fmt.Sprint(MakeSet(v.arrStr))

		t.Logf("Calling MakeSet(%v), resault %s\n", v.arrStr, res)
		assert.Equal(t, v.expected, res, fmt.Sprintf("Incorrect resault. Expect %s, got %s", v.expected, res))
	}
}
