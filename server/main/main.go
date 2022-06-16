package main

import (
	"fmt"
	"net"
	"os"
)

const ADDR = ":8080"

func main() {
	listener, err := net.Listen("tcp", ADDR)
	if err != nil {
		fmt.Errorf("%s\n", err)
		os.Exit(1)
		return
	}
	defer listener.Close()
	fmt.Printf("Listening for TCP at %s\n", ADDR)

	for {
		fmt.Println("Waiting for connection")
		conn, err := listener.Accept()
		fmt.Println("Connected")
		if err != nil {
			os.Exit(1)
			return
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	fmt.Println("  Reading data")

	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	if err != nil {
		fmt.Errorf("%s\n", err)
		return
	}
	status := string(buffer)

	fmt.Println("  Message is:", status)
	conn.Write([]byte("ok"))
	fmt.Println("  Written response")
	conn.Close()
	fmt.Println("Closed connection")
}
