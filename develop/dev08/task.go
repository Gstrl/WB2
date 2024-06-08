package main

import (
	"fmt"
	"log"
	"os/exec"
	"syscall"
)

func CmdRunner(comand string) string {
	msg := exec.Command("powershell", comand)
	msg.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	out, err := msg.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}

func main() {

	// Print the output
	fmt.Println(CmdRunner("pwd"))
}
