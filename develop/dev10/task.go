package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Telnet struct {
	address string
	timeout time.Duration
	conn    net.Conn
}

func (t *Telnet) FromArgs() error {
	timeout := flag.String("timeout", "10s", "Usage: go run task.go [--timeout=timeout] host port")
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		return errors.New("usage: go run task.go [--timeout=timeout] host port")
	}
	t.address = args[0]

	timeoutParse, err := time.ParseDuration(*timeout)
	if err != nil {
		return errors.New("the timeout flag is specified incorrectly")
	}
	t.timeout = timeoutParse
	return nil
}

func (t *Telnet) Connect() error {
	var err error
	t.conn, err = net.DialTimeout("tcp", t.address, t.timeout)
	if err != nil {
		return fmt.Errorf("error when connecting to %s: %w", t.address, err)
	}
	return nil
}

func (t *Telnet) ReadConn(done chan<- bool) {
	for {
		if _, err := io.Copy(os.Stdout, t.conn); err != nil {
			if errors.Is(err, io.EOF) {
				fmt.Println("Connection closed by remote host")
			} else {
				fmt.Println("Error reading from connection", err)
			}

			done <- true
			break
		}
	}
}

func (t *Telnet) WriteConn(done chan<- bool) {
	for {
		if _, err := io.Copy(t.conn, os.Stdin); err != nil {
			if errors.Is(err, io.EOF) {
				fmt.Println("Connection closed by remote host")
			} else {
				fmt.Println("Error reading from connection", err)
			}
			done <- true
			break
		}
	}
}

func (t *Telnet) Run() {
	err := t.FromArgs()
	if err != nil {
		log.Fatal(err)
	}

	err = t.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer t.conn.Close()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	done := make(chan bool, 1)

	go t.ReadConn(done)
	go t.WriteConn(done)

	for {
		select {
		case <-sigChan:
			fmt.Println("Closing connection... Transferred to syscall.")
			return
		case <-done:
			fmt.Println("Closing connection...")
			return
		}
	}
}

func main() {
	var telnet Telnet
	telnet.Run()
}
