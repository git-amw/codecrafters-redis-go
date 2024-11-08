package main

import (
	"errors"
	"fmt"
	"io"
	"net"
	"os"
)

// Ensures gofmt doesn't remove the "net" and "os" imports in stage 1 (feel free to remove this!)
var _ = net.Listen
var _ = os.Exit

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	//Uncomment this block to pass the first stage

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		go func(conn net.Conn) {
			defer conn.Close()

			// Loop to continuously handle messages from the same connection
			buf := make([]byte, 128)
			for {
				// Read incoming data
				_, err := conn.Read(buf)
				if errors.Is(err, io.EOF) {
					break
				}
				if err != nil {
					fmt.Println("Error reading from client:", err)
					return
				}

				// Process the message, respond with PONG if it's a ping
				conn.Write([]byte("+PONG\r\n"))
			}
		}(conn)
	}
}
