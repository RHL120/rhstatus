package api

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func runServer(ch chan func() error) {
	listener, err := net.Listen("tcp", "localhost:6666")
	if err != nil {
		fmt.Println("Failed to start IPC server: %v", err)
		ch <- func() error {
			os.Exit(1)
			return nil
		}
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Failed to accept a connection, error: %v", err)
			go handelConnection(conn, ch)
		}
	}
}

func handelConnection(conn net.Conn, ch chan func() error) {
	defer conn.Close()
	data := make([]byte, 0)
	_, err := conn.Read(data)
	if err != nil {
		fmt.Println("Failed to read from a connection, error: %v", err)
		return
	}
	args := strings.Split(string(data), " ")
	if len(args) <= 0 {
		conn.Write([]byte("Expected a command"))
		return
	}
	cmd, ok := commands[args[0]]
	if !ok {
		conn.Write([]byte("Command not found: " + args[0]))
		return
	}
	if cmd.argCount != uint8(len(args)) {
		conn.Write([]byte(fmt.Sprintf("Command %s expects %d arguments", args[0], cmd.argCount)))
	}
	ch <- cmd.function(args[1:])

}
