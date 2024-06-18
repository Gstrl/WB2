package main

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
- "a4bc2d5e" => "aaaabccddddde"
- "abcd" => "abcd"
- "45" => "" (некорректная строка)
- "" => ""

Дополнительное задание: поддержка escape - последовательностей
- qwe\4\5 => qwe45 (*)
- qwe\45 => qwe44444 (*)
- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.
Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// UnpackString распаковывает строку согласно заданному формату.
func UnpackString(s string) (string, error) {
	var result strings.Builder
	var prevRune rune
	for i, r := range s {
		if unicode.IsDigit(r) {
			if i == 0 || unicode.IsDigit(prevRune) {
				// Если строка начинается с цифры или содержит две подряд идущие цифры,
				// то это некорректная строка.
				return "", errors.New("некорректная строка")
			}
			count, _ := strconv.Atoi(string(r))
			result.WriteString(strings.Repeat(string(prevRune), count-1))
		} else {
			result.WriteRune(r)
		}
		prevRune = r
	}
	return result.String(), nil
}

// UnpackStringWithEscape распаковывает строку реализовывая поддержку escape-последовательностей.
func UnpackStringWithEscape(s string) (string, error) {
	var result strings.Builder
	var prevRune rune
	var escaping = true
	const slashRune rune = 92

	for _, r := range s {
		// елси символ == \ и до этого символа были отличные от \
		// попдаем в escaping зону (ингорируем последний сивол предыдущей зоны)
		if r == slashRune && escaping {
			escaping = false
			prevRune = -1
			continue
		}
		if r != slashRune {
			escaping = true
		}

		if unicode.IsDigit(r) {
			if prevRune == -1 {
				result.WriteRune(r)
			} else {
				count, err := strconv.Atoi(string(r))
				if err != nil {
					return "", err
				}
				result.WriteString(strings.Repeat(string(prevRune), count-1))
			}
		} else {
			result.WriteRune(r)
		}
		prevRune = r
	}

	return result.String(), nil
}
