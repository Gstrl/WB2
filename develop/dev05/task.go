package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type EnvGrep struct {
	after      int
	before     int
	context    int
	count      bool
	ignoreCase bool
	invert     bool
	fixed      bool
	lineNum    bool
	pattern    string
	fileName   string
}

func (env *EnvGrep) FromArgs() {
	// Определение флагов
	flagA := flag.Int("A", 0, "печатать +N строк после совпадения")
	flagB := flag.Int("B", 0, "печатать +N строк до совпадения")
	flagC := flag.Int("C", 0, "печатать ±N строк вокруг совпадения")
	flagc := flag.Bool("c", false, "количество строк")
	flagI := flag.Bool("i", false, "игнорировать регистр")
	flagV := flag.Bool("v", false, "вместо совпадения, исключать")
	flagF := flag.Bool("F", false, "точное совпадение со строкой, не паттерн")
	flagN := flag.Bool("n", false, "печатать номер строки")

	flag.Parse()

	args := flag.Args()

	if len(args) < 2 {
		fmt.Println("Необходимо указать паттерн для поиска.")
		os.Exit(1)
	}

	env.pattern = args[0]
	env.fileName = args[1]

	env.after = *flagA
	env.before = *flagB
	env.context = *flagC
	env.count = *flagc
	env.ignoreCase = *flagI
	env.invert = *flagV
	env.fixed = *flagF
	env.lineNum = *flagN
}

type Grep struct {
	env          EnvGrep
	lines        []string
	satisfyIndex []int
	res          []int
}

func (g *Grep) search() {
	if !g.env.fixed {
		if g.env.ignoreCase {
			for i := 0; i <= len(g.lines)-1; i++ {
				contain := strings.Contains(strings.ToLower(g.lines[i]), strings.ToLower(g.env.pattern))
				if contain {
					g.satisfyIndex = append(g.satisfyIndex, i)
				}
			}
		} else {
			for i := 0; i <= len(g.lines)-1; i++ {
				contain := strings.Contains(g.lines[i], g.env.pattern)
				if contain {
					g.satisfyIndex = append(g.satisfyIndex, i)
				}
			}
		}
	} else {
		for i := 0; i <= len(g.lines)-1; i++ {
			if g.lines[i] == g.env.pattern {
				g.satisfyIndex = append(g.satisfyIndex, i)
			}
		}
	}

	g.res = append(g.res, g.satisfyIndex...)
}

func (g *Grep) after() {
	for _, v := range g.satisfyIndex {
		for i := 1; i <= g.env.after; i++ {
			if v+i <= len(g.lines)-1 {
				g.res = append(g.res, v+i)
			}
		}
	}
}

func (g *Grep) before() {
	for _, v := range g.satisfyIndex {
		for i := 1; i <= g.env.before; i++ {
			if v-i >= 0 {
				g.res = append(g.res, v-i)
			}
		}
	}
}

func (g *Grep) context() {
	for _, v := range g.satisfyIndex {
		for i := 1; i <= g.env.context; i++ {
			if v-i >= 0 {
				g.res = append(g.res, v-i)
			}
		}
	}

	for _, v := range g.satisfyIndex {
		for i := 1; i <= g.env.context; i++ {
			if v+i <= len(g.lines)-1 {
				g.res = append(g.res, v+i)
			}
		}
	}
}

func (g *Grep) count() int {
	return len(g.satisfyIndex)
}

func (g *Grep) invert() (resStr []string, resLineNum []int) {
	for i := 0; i <= len(g.lines)-1; i++ {
		_, ok := slices.BinarySearch(g.res, i)
		if !ok {
			resStr = append(resStr, g.lines[i])
			resLineNum = append(resLineNum, i+1)
		}
	}
	return resStr, resLineNum
}

func (g *Grep) createRes() (resStr []string, resLineNum []int) {
	sort.Ints(g.res)
	g.res = unique(g.res)

	if g.env.invert {
		return g.invert()
	}

	for _, v := range g.res {
		resStr = append(resStr, g.lines[v])
		resLineNum = append(resLineNum, v+1)
	}

	return resStr, resLineNum
}

func (g *Grep) Run() {
	g.search()

	var resStr []string
	var resLineNum []int

	if g.env.context != 0 {
		g.context()
	}
	if g.env.after != 0 {
		g.after()
	}
	if g.env.before != 0 {
		g.before()
	}

	// Открытие файла для записи
	file, err := os.Create("_result.txt")
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		os.Exit(1)
	}
	defer file.Close()

	if g.env.count {
		fmt.Print(g.count())
		file.WriteString(strconv.Itoa(g.count()))
	} else {
		resStr, resLineNum = g.createRes()
	}

	if resStr != nil {
		if g.env.lineNum {
			for i := 0; i <= len(resStr)-1; i++ {
				fmt.Printf("%d:%s\n", resLineNum[i], resStr[i])
				if i != len(resStr)-1 {
					file.WriteString(fmt.Sprintf("%d:%s\n", resLineNum[i], resStr[i]))
				} else {
					file.WriteString(fmt.Sprintf("%d:%s", resLineNum[i], resStr[i]))
				}
			}
		} else {
			for i := 0; i <= len(resStr)-1; i++ {
				fmt.Println(resStr[i])
				if i != len(resStr)-1 {
					file.WriteString(fmt.Sprint(resStr[i], "\n"))
				} else {
					file.WriteString(fmt.Sprint(resStr[i]))
				}
			}
		}
	}
}

func NewGrep(env EnvGrep) (*Grep, error) {

	file, err := os.Open(env.fileName)
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

	satisfyIndex := make([]int, 0)

	return &Grep{
		env:          env,
		lines:        lines,
		satisfyIndex: satisfyIndex,
	}, nil
}

func unique(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func main() {
	var env EnvGrep
	env.FromArgs()
	g, err := NewGrep(env)
	if err != nil {
		log.Fatal(err)
	}
	g.Run()
}
