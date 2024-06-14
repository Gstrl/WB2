package main

import (
	"bytes"
	"flag"
	"net"
	"os"
	"strings"
	"testing"
	"time"
)

func TestFromArgs(t *testing.T) {
	tests := []struct {
		args         []string
		expectedAddr string
		expectedErr  bool
	}{
		{[]string{"task.go", "--timeout=5s", "localhost:8080"}, "localhost:8080", false},
		{[]string{"task.go", "localhost:8080"}, "localhost:8080", false},
		{[]string{"task.go"}, "", true},
	}

	for _, tt := range tests {
		t.Run(strings.Join(tt.args, " "), func(t *testing.T) {
			originalArgs := os.Args
			defer func() { os.Args = originalArgs }()
			os.Args = tt.args

			flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
			var telnet Telnet
			err := telnet.FromArgs()
			if (err != nil) != tt.expectedErr {
				t.Errorf("expected error: %v, got: %v", tt.expectedErr, err)
			}
			if telnet.address != tt.expectedAddr {
				t.Errorf("expected address: %s, got: %s", tt.expectedAddr, telnet.address)
			}
		})
	}
}

func TestConnect(t *testing.T) {
	l, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	defer l.Close()

	go func() {
		conn, _ := l.Accept()
		defer conn.Close()
	}()

	var telnet Telnet
	telnet.address = l.Addr().String()
	telnet.timeout = 1 * time.Second

	err = telnet.Connect()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestReadWriteConn(t *testing.T) {
	server, client := net.Pipe()
	defer server.Close()
	defer client.Close()

	var telnet Telnet
	telnet.conn = client

	done := make(chan bool, 1)

	go func() {
		time.Sleep(500 * time.Millisecond)
		server.Write([]byte("hello from server"))
	}()

	go telnet.ReadConn(done)

	select {
	case <-done:
		t.Fatal("ReadConn terminated unexpectedly")
	case <-time.After(1 * time.Second):
		// Pass
	}

	go func() {
		buf := make([]byte, 1024)
		n, err := server.Read(buf)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if !bytes.Equal(buf[:n], []byte("hello from client")) {
			t.Fatalf("expected 'hello from client', got %s", buf[:n])
		}
		done <- true
	}()

	go func() {
		telnet.WriteConn(done)
	}()

	client.Write([]byte("hello from client"))

	select {
	case <-done:
		// Pass
	case <-time.After(1 * time.Second):
		t.Fatal("WriteConn did not receive expected data")
	}
}
