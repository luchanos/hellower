package main

import (
	"fmt"
	"log"
	"time"
)

func longSQLquery() chan int {
	ch := make(chan int)
	return ch
}

func Example1() {
	timer := time.NewTimer(1 * time.Second)
	select {
	case <-timer.C: // у таймера есть свой канал
		fmt.Println("timer.C timeout happend") // не заблокирует
	case <-time.After(time.Minute): // заблокирует, пока не выполнится
		fmt.Println("time.After timeout happened")
	case result := <-longSQLquery():
		if !timer.Stop() {
			<-timer.C
		}
		fmt.Println("operation result", result)
	}
}

func Example2() {
	ticker := time.NewTicker(time.Second)
	i := 0
	for tickTime := range ticker.C {
		i++
		fmt.Println("step", i, "time", tickTime)
		if i >= 5 {
			ticker.Stop()
			break
		}
	}
}

func HelloWorld() {
	fmt.Println("Hello, World!")
}

func Example3() {
	timer := time.AfterFunc(10*time.Second, HelloWorld)

	_, err := fmt.Scanln()
	if err != nil {
		log.Fatal(err)
	}
	timer.Stop()

	_, err = fmt.Scanln()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	Example3()
}
