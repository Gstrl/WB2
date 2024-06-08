package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)

Основное.
Поддержать ключ:
-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное.
Поддержать ключи:
-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type AppEnv struct {
	isNumeric       bool
	isReverse       bool
	deleteDuplicate bool
	column          int
	lines           []string
}

func (app *AppEnv) getColumn(line string) string {
	fields := strings.Fields(line)
	if app.column > 0 && app.column <= len(fields) {
		return fields[app.column-1]
	}

	return line
}

func (app *AppEnv) removeDuplicates() {
	unique := make(map[string]struct{})
	var result []string

	for _, v := range app.lines {
		if _, exist := unique[v]; !exist {
			unique[v] = struct{}{}
			result = append(result, v)
		}
	}

	app.lines = result
}

func (app *AppEnv) sort() {
	// Функция сравнения строк для сортировки
	compare := func(i, j int) bool {
		s1 := app.getColumn(app.lines[i])
		s2 := app.getColumn(app.lines[j])

		// Преобразование в числа, если указан флаг -n
		if app.isNumeric {
			num1, err1 := strconv.Atoi(s1)
			num2, err2 := strconv.Atoi(s2)
			if err1 == nil && err2 == nil {
				switch {
				case app.isReverse && num1 < num2:
					return false
				case app.isReverse && num1 > num2:
					return true
				case !app.isReverse && num1 < num2:
					return true
				case !app.isReverse && num1 > num2:
					return false
				}
			}
		}

		// Сравнение строк
		result := strings.Compare(s1, s2)

		// Применение флага -r
		if app.isReverse {
			return result > 0
		}

		return result < 0
	}

	// Применение флага -u
	if app.deleteDuplicate {
		app.removeDuplicates()
	}

	sort.SliceStable(app.lines, compare)
}

func (app *AppEnv) createSortFile() error {
	// Открытие файла для записи
	file, err2 := os.Create("_sorted.txt")
	if err2 != nil {
		fmt.Println("Ошибка открытия файла:", err2)
		os.Exit(1)
	}
	defer func() {
		file.Close()
	}()

	// Записываем строки в файл
	for idx, value := range app.lines {
		_, err := file.WriteString(value)
		if err != nil {
			return err
		}
		if idx != len(app.lines)-1 {
			_, err := file.WriteString("\n")
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (app *AppEnv) RunApp() error {
	app.sort()
	err := app.createSortFile()
	if err != nil {
		return err
	}
	return nil
}

func NewAppEnv(isNumeric, isReverse, deleteDuplicate bool, column int, filePath string) (*AppEnv, error) {
	// Открытие файла для чтения
	if filePath == "" {
		return nil, errors.New("недопустимое имя для файла")
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	// Считывание строк из файла
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return &AppEnv{
		isNumeric:       isNumeric,
		isReverse:       isReverse,
		deleteDuplicate: deleteDuplicate,
		column:          column,
		lines:           lines,
	}, nil
}

func main() {
	// Определение флагов
	flagF := flag.String("file", "", "Указание колонки для сортировки (по умолчанию 0 - вся строка)")
	flagK := flag.Int("k", 0, "Указание колонки для сортировки (по умолчанию 0 - вся строка)")
	flagN := flag.Bool("n", false, "Сортировать по числовому значению")
	flagR := flag.Bool("r", false, "Сортировать в обратном порядке")
	flagU := flag.Bool("u", false, "Не выводить повторяющиеся строки")

	flag.Parse()

	app, err := NewAppEnv(*flagN, *flagR, *flagU, *flagK, *flagF)
	if err != nil {
		log.Fatal(err)
	}

	err = app.RunApp()
	if err != nil {
		log.Fatal(err)
	}
}
