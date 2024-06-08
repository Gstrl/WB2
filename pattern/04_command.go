package main

import "fmt"

//Шаблон команды, как следует из названия, используется, когда мы хотим создавать и выполнять «команды».
//Разные команды имеют свою реализацию, но этапы выполнения одинаковы.

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
