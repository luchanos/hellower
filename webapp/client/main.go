package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func ErrorHandler(err error, msg string) {
	if err != nil {
		fmt.Printf("%s %s", err, msg)
	}
}

func makeRerquest() {
	conn, err := net.Dial("tcp", "0.0.0.0:8080")
	ErrorHandler(err, "unable to connect to host!")
	reader := bufio.NewReader(os.Stdin)
	for {
		data, err := reader.ReadString('\n')
		ErrorHandler(err, "smth went wrong")
		fmt.Fprintf(conn, data)
	}
}

func main() {
	makeRerquest()
}
