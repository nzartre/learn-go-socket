package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const ADDR = ":8080"

func main() {
	// Connect
	conn, err := net.Dial("tcp", ADDR)
	if err != nil {
		os.Exit(1)
		return
	}
	defer conn.Close()
	fmt.Println("Connected to server")

	// Send a message
	conn.Write([]byte("test"))
	fmt.Println("Sent a message")

	// Read a response
	status, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Errorf("%s\n", err)
		os.Exit(1)
	}
	fmt.Println("Result is:", status)
}
