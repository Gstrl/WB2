package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

func Shell(input string) {
	// Для выхода из программы используем 'quit'
	if input == "quit" {
		syscall.Exit(0)
	}

	// Обработка конвейера на пайпах
	pipeline := strings.Split(input, " | ")

	for _, pipe := range pipeline {
		// Разделяем команду на аргументы
		args := strings.Fields(pipe)
		if len(args) == 0 {
			continue
		}

		// Обрабатываем различные команды:
		switch args[0] {
		case "cd":
			// Смена директории
			if len(args) < 2 {
				// Если аргументы отсутствуют, то переходим в домашнуюю директорию пользователя
				home, err := os.UserHomeDir()
				if err != nil {
					fmt.Fprintf(os.Stderr, "cd: %v\n", err)
				}

				err = os.Chdir(home)
				if err != nil {
					fmt.Fprintf(os.Stderr, "cd: %v\n", err)
				}
			} else {
				// Если указан путь, то переходим по указанному пути
				err := os.Chdir(args[1])
				if err != nil {
					fmt.Fprintf(os.Stderr, "cd: %v\n", err)
				}
			}
		case "pwd":
			// Показывает путь до текущего каталога
			dir, err := os.Getwd()
			if err != nil {
				fmt.Fprintf(os.Stderr, "pwd: %v\n", err)
			}

			fmt.Println(dir)
		case "echo":
			// Вывод аргумента в STDOUT
			fmt.Println(strings.Join(args[1:], " "))
		case "kill":
			// Завершение процесса, переданного в качестве аргумента
			if len(args) < 2 {
				fmt.Println("kill: missing argument")
			} else {
				// Преобразуем аргумент в ProcessID
				pid, err := strconv.Atoi(args[1])
				if err != nil {
					fmt.Println(err)
				}

				// Поиск процесса по указанному ProcessID
				proc, err := os.FindProcess(pid)
				if err != nil {
					fmt.Fprintf(os.Stderr, "kill: %v\n", err)
				}

				// Завершаем процесс с указанным ProcessID
				err = proc.Kill()
				if err != nil {
					fmt.Fprintf(os.Stderr, "kill: %v\n", err)
				}
			}
		case "ps":
			// Вывод информации по запущенным процессам
			cmd := exec.Command("ps")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			err := cmd.Run()
			if err != nil {
				fmt.Fprintf(cmd.Stderr, "ps: %v\n", err)
			}
		default:
			// Выполнение других команд
			cmd := exec.Command(args[0], args[1:]...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			err := cmd.Run()
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v: %v\n", args[0], err)
			}
		}
	}
}

func main() {
	fmt.Println("Shell on Go. Type 'quit' to quit.")

	for {
		fmt.Print("> ")
		// Считываем данные из STDIN
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()
		Shell(input)
	}
}
