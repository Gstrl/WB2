package main

import (
	"errors"
	"testing"
)

func TestUnpackString(t *testing.T) {
	testTable := []struct {
		str         string
		expectedStr string
		expectedErr error
	}{
		{
			str:         "a4bc2d5e",
			expectedStr: "aaaabccddddde",
			expectedErr: nil,
		},
		{
			str:         "abcd",
			expectedStr: "abcd",
			expectedErr: nil,
		},
		{
			str:         "45",
			expectedStr: "",
			expectedErr: errors.New("некорректная строка"),
		},
		{
			str:         "",
			expectedStr: "",
			expectedErr: nil,
		},
	}

	for _, test := range testTable {
		resault, err := UnpackString(test.str)
		if resault != test.expectedStr && err != test.expectedErr {
			t.Errorf("Incorrect resault: %s", test.expectedStr)
		}
	}
}

func TestUnpackStringWithEscape(t *testing.T) {
	testTable := []struct {
		str         string
		expectedStr string
		expectedErr error
	}{
		{
			str:         "qwe\\4\\5",
			expectedStr: "qwe45",
			expectedErr: nil,
		},
		{
			str:         "qwe\\45",
			expectedStr: "qwe44444",
			expectedErr: nil,
		},
		{
			str:         "qwe\\\\5",
			expectedStr: "qwe\\\\\\\\\\",
			expectedErr: nil,
		},
		{
			str:         "\\\\",
			expectedStr: "\\",
			expectedErr: nil,
		},
	}

	for _, test := range testTable {
		resault, err := UnpackStringWithEscape(test.str)
		if resault != test.expectedStr && err != test.expectedErr {
			t.Errorf("Incorrect test:%s, resault: %s", test.str, resault)
		}
	}
}
