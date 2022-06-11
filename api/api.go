package api

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func RunServer(ch chan func() error) {
	listener, err := net.Listen("tcp", "localhost:6666")
	if err != nil {
		fmt.Printf("Failed to start IPC server: %v\n", err)
		ch <- func() error {
			os.Exit(1)
			return nil
		}
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Failed to accept a connection, error: %v\n", err)
		}
		go handelConnection(conn, ch)
	}
}

func handelConnection(conn net.Conn, ch chan func() error) {
	defer conn.Close()
	data, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Printf("Failed to read from a connection, error: %v\n", err)
		return
	}
	args := strings.Split(strings.Trim(data, "\n"), " ")
	if len(args) <= 0 {
		conn.Write([]byte("Expected a command\n"))
		return
	}
	cmd, ok := commands[args[0]]
	if !ok {
		conn.Write([]byte("Command not found: " + args[0] + "\n"))
		return
	}
	if cmd.argCount != uint8(len(args)-1) {
		conn.Write([]byte(fmt.Sprintf("Command %s expects %d arguments\n", args[0], cmd.argCount)))
	}
	conn.Write([]byte("Executing\n"))
	ch <- cmd.function(args[1:])

}
