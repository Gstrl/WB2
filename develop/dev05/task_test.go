package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"testing"
)

func TestGrep(t *testing.T) {

	var TestTable = []struct {
		env EnvGrep
	}{
		{
			env: EnvGrep{
				after:    2,
				lineNum:  true,
				pattern:  "apple 3",
				fixed:    true,
				fileName: "_income.txt",
			},
		},
		{
			env: EnvGrep{
				after:    2,
				count:    true,
				pattern:  "app",
				fileName: "_income.txt",
			},
		},
		{
			env: EnvGrep{
				before:   2,
				pattern:  "app",
				fileName: "_income.txt",
			},
		},
		{
			env: EnvGrep{
				context:  1,
				pattern:  "1",
				lineNum:  true,
				fileName: "_income.txt",
			},
		},
		{
			env: EnvGrep{
				invert:   true,
				pattern:  "1",
				fileName: "_income.txt",
			},
		},
		{
			env: EnvGrep{
				ignoreCase: true,
				pattern:    "honeydew",
				fileName:   "_income.txt",
			},
		},
	}

	for i, v := range TestTable {
		g, err := NewGrep(v.env)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Номер теста", i+1)
		g.Run()

		// Открытие файла для чтения ожидаемых результатов
		filename := fmt.Sprintf("_testCase/_expected%v.txt", i+1)
		fileExpected, err := os.Open(filename)
		if err != nil {
			t.Errorf("Ошибка открытия файла: %v", err)
		}
		defer fileExpected.Close()

		// Открытие файла для чтения полученных результатов
		fileSorted, err := os.Open("_result.txt")
		if err != nil {
			t.Errorf("Ошибка открытия файла: %v", err)
		}
		defer fileSorted.Close()

		expected, _ := io.ReadAll(fileExpected)
		result, _ := io.ReadAll(fileSorted)

		if !bytes.Equal(result, expected) {
			t.Errorf("%v:Неожиданный результат: фактический результат = %v\nожидаемый результат = %v\nаргументы - %v", i+1, result, expected, v)
		}
	}
}
