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
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type Cut struct {
	fields    []int
	delimiter string
	separated bool
}

func (c *Cut) FromArgs() error {
	flagF := flag.String("f", "", "выбрать поля (колонки)")
	flagD := flag.String("d", "	", "использовать другой разделитель")
	flagS := flag.Bool("s", false, "только строки с разделителем")

	flag.Parse()

	var fields []string
	if strings.Contains(*flagF, ",") {
		fields = strings.Split(*flagF, ",")
	} else {
		return errors.New("flag F is not set or set incorrectly")
	}

	fieldsInt := make([]int, 0, len(fields))

	for _, v := range fields {
		index, err := strconv.Atoi(v)
		if err != nil {
			return errors.New("flag F set incorrectly")
		}
		fieldsInt = append(fieldsInt, index)
	}
	sort.Ints(fieldsInt)

	c.fields = fieldsInt
	c.delimiter = *flagD
	c.separated = *flagS
	return nil
}

func (c *Cut) Run(line string) string {
	if c.separated && !strings.Contains(line, c.delimiter) {
		return ""
	}

	splitLine := strings.Split(line, c.delimiter)
	resArr := make([]string, 0, len(c.fields))
	for _, v := range c.fields {
		if v-1 <= len(splitLine)-1 {
			resArr = append(resArr, splitLine[v-1])
		}
	}
	resStr := strings.Join(resArr, c.delimiter)
	return resStr
}

func main() {
	var cut Cut
	err := cut.FromArgs()
	if err != nil {
		log.Fatal(err)
	}
	// Считываем данные из STDIN
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(cut.Run(line))
	}
}
