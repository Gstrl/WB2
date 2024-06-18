package main

import "fmt"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

/*
Применимость паттерна "комманда":
1. Организация очереди: Паттерн поддерживает создание и управление очередью запросов.
2. Отмена операций: Позволяет реализовать отмену и повторение операций.
3. Параметризация объектов: Команды могут быть параметризованы и передаваться между объектами, что делает их многоразовыми и гибкими.

Плюсы паттерна "комманда":
1. Принцип единой ответственности: Классы, вызывающие операции, можно отделить от классов, выполняющих эти операции.
2. Принцип открытости/закрытости: Можно вводить в приложение новые команды, не ломая существующий код.
3. Позволяет реализовать отмену и повторение операций.

Минусы паттерна "комманда":
1. Код может стать более сложным, поскольку вводится новый слой между отправителями и получателями.
*/

type Command interface {
	Execute() string
}

type PingCommand struct{}

func (p *PingCommand) Execute() string {
	return "react pings"
}

type StatusCommand struct{}

func (p *StatusCommand) Execute() string {
	return "status command"
}

func execByName(name string) string {

	commands := map[string]Command{
		"ping":   &PingCommand{},
		"status": &StatusCommand{},
	}

	if command := commands[name]; command == nil {
		return "No such command found, throw error?"
	} else {
		return command.Execute()
	}
}

func main() {

	fmt.Println(execByName("status"))
	fmt.Println(execByName("ping"))

	fmt.Println(execByName("unkown"))
}
