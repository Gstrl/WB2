package main

import (
	"testing"
)

func TestAppEnv_RunApp(t *testing.T) {
	testTable := []struct {
		app     AppEnv
		example []string
	}{
		{
			app: AppEnv{
				isNumeric:       false,
				isReverse:       false,
				deleteDuplicate: false,
				column:          0,
				lines:           []string{""},
			},
			example: []string{""},
		},
		{
			app: AppEnv{
				isNumeric:       true,
				isReverse:       true,
				deleteDuplicate: false,
				column:          1,
				lines:           []string{"1 -3 2", "45 0 4", "323 2 4", "1 34 -1"},
			},
			example: []string{"323 2 4", "45 0 4", "1 -3 2", "1 34 -1"},
		},
		{
			app: AppEnv{
				isNumeric:       true,
				isReverse:       true,
				deleteDuplicate: true,
				column:          1,
				lines:           []string{"1 -3 2", "45 0 4", "323 2 4", "323 2 4", "1 34 -1"},
			},
			example: []string{"323 2 4", "45 0 4", "1 -3 2", "1 34 -1"},
		},
	}

	for _, v := range testTable {
		err := v.app.RunApp()
		if !Equal(v.app.lines, v.example) || err != nil {
			t.Errorf("Incorrect test: %v", v.app.lines)
		}

	}
}

func TestNewAppEnv(t *testing.T) {
	_, err := NewAppEnv(false, false, false, 0, "")
	if err == nil {
		t.Error("Incorrect test")
	}

	_, err = NewAppEnv(true, true, true, 1, "_unsorted.txt")
	if err != nil {
		t.Error("Incorrect test")
	}
}

func Equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestAll(t *testing.T) {
	TestNewAppEnv(t)
	TestAppEnv_RunApp(t)
}
