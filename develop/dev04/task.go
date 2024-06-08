package main

import (
	"fmt"
	"strings"
)

// linter ругается на Magic number. Тк Входные данные для функции: ссылка на массив,
// каждый элемент которого - слово на русском языке в кодировке utf8.
// в некоторый местах для получения количества символов, я делю
// количество байт на 2, отсюда берется Magic number

// Set  setSymbols - символы множества и их количество, setElem - элементы множества
type Set struct {
	setSymbols map[rune]int
	setElem    []string
}

// CheckSet проверяет входит ли слово в множество
func (s *Set) CheckSet(str string) bool {
	setStr := make(map[rune]int, len(str)/2)
	for _, v := range str {
		_, ok := setStr[v]
		if ok {
			setStr[v]++
		} else {
			setStr[v] = 1
		}
	}

	// если символы и их количество в слове совпадает с символами множества возвращаем true
	for _, v := range str {
		_, ok := s.setSymbols[v]
		if ok {
			if setStr[v] != s.setSymbols[v] {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

// AppendSet добавляет слово к элементам множества
func (s *Set) AppendSet(str string) {
	s.setElem = append(s.setElem, str)
}

// newSet создает новое множество (символы множества и первый элемент)
func newSet(str string) Set {
	setElem := make([]string, 1, 2)
	setElem[0] = str
	setSymbols := make(map[rune]int, len(str)/2)
	for _, v := range str {
		_, ok := setSymbols[v]
		if ok {
			setSymbols[v]++
		} else {
			setSymbols[v] = 1
		}
	}
	return Set{
		setElem:    setElem,
		setSymbols: setSymbols,
	}
}

// MakeSet группирует слайс строк в множества
func MakeSet(arrStr []string) map[string][]string {
	if len(arrStr) == 0 {
		return nil
	}
	// приведение к нижнему регистру
	for i := 0; i <= len(arrStr)-1; i++ {
		arrStr[i] = strings.ToLower(arrStr[i])
	}

	// создаем слайс множеств и добавляем первое множество
	arrSet := make([]Set, 0)
	arrSet = append(arrSet, newSet(arrStr[0]))

	// если не попали ни в одно множество создаем новое
	for i := 1; i <= len(arrStr)-1; i++ {
		// если в слове <=1 оно не может быть множеством
		if len(arrStr[i]) <= 1 {
			continue
		}

		hitSet := false

		for ii := 0; ii <= len(arrSet)-1; ii++ {
			if arrSet[ii].CheckSet(arrStr[i]) {
				hitSet = true
				arrSet[ii].AppendSet(arrStr[i])
				break
			}
		}
		if !hitSet {
			arrSet = append(arrSet, newSet(arrStr[i]))
		}
	}
	return resMap(arrSet)
}

// resMap проверяет количество элементов множества и возвращает результат группировки
func resMap(s []Set) map[string][]string {
	res := make(map[string][]string)
	for i := 0; i <= len(s)-1; i++ {
		if len(s[i].setElem) == 1 {
			continue
		}
		res[s[i].setElem[0]] = s[i].setElem
	}
	return res
}

func main() {
	arrStr := []string{"сос", "кек", "ссо", "осс", "лол"}
	fmt.Println(MakeSet(arrStr))
}
