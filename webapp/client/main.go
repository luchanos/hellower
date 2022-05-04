package main

import (
	"bufio"
	"fmt"
	"hellower/errorhandler"
	"net"
	"os"
)

func makeRerquest() {
	conn, err := net.Dial("tcp", "0.0.0.0:8080")
	errorhandler.ErrorHandler(err, "unable to connect to host!")
	reader := bufio.NewReader(os.Stdin)
	for {
		data, err := reader.ReadString('\n')
		errorhandler.ErrorHandler(err, "smth went wrong")
		fmt.Fprintf(conn, data)
	}
}

func main() {
	makeRerquest()
}
