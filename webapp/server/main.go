package main

import (
	"bufio"
	"fmt"
	"hellower/errorhandler"
	"net"
	"time"
)

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
		errorhandler.ErrorHandler(err, "unable to read received bytes")

		fmt.Printf("%s", bytes)
		fmt.Fprintf(conn, "I've got it!")
	}

}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	errorhandler.ErrorHandler(err, "unable to create server")

	defer func() {
		listener.Close()
		fmt.Println("listener is closed")
	}()

	for {
		conn, err := listener.Accept()
		errorhandler.ErrorHandler(err, "unable to get connection with client")
		go HandleConnection(conn)
	}
}
