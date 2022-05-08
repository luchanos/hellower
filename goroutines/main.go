package main

import (
	"fmt"
	"runtime"
	"strconv"
)

func DoSomeWork(in int) {
	for j := 0; j < 7; j++ {
		fmt.Printf(formatWork(in, j))
		runtime.Gosched()
	}
}

func formatWork(in int, j int) string {
	res := ""
	for i := 0; i < j; i++ {
		res += strconv.Itoa(j)
	}
	return res + "\n"
}

func main() {
	for i := 0; i < 5; i++ {
		go DoSomeWork(i)
	}
	fmt.Scanln()
}
