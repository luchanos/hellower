package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func ErrorHandler(err error, msg string) {
	if err != nil {
		fmt.Printf("%s %s", err, msg)
	}
}

func HandleConnection(conn net.Conn) {

	defer func() {
		fmt.Println("closing connection...")
		conn.Close()
		fmt.Println("connection has been closed")
	}()

	timeoutDuration := 5 * time.Second
	bufReader := bufio.NewReader(conn)

	for {
		conn.SetReadDeadline(time.Now().Add(timeoutDuration))

		bytes, err := bufReader.ReadBytes('\n')
		ErrorHandler(err, "unable to read received bytes")

		fmt.Printf("%s", bytes)
		fmt.Fprintf(conn, "I've got it!")
	}

}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	ErrorHandler(err, "unable to create server")

	defer func() {
		listener.Close()
		fmt.Println("listener is closed")
	}()

	for {
		conn, err := listener.Accept()
		ErrorHandler(err, "unable to get connection with client")
		go HandleConnection(conn)
	}
}
