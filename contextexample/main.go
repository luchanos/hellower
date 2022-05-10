package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func worker(ctx context.Context, workerNum int, out chan<- int) {
	waitTime := time.Duration(rand.Intn(100)+10) * time.Millisecond
	fmt.Println(workerNum, "sleep", waitTime)
	select {
	case <-ctx.Done():
		return
	case <-time.After(waitTime):
		fmt.Println("worker", workerNum, "done")
		out <- workerNum
	}
}

func Example1() {
	ctx, finish := context.WithCancel(context.Background()) // тут вернулся сам контекст и функция отмены
	result := make(chan int, 1)
	for i := 0; i <= 10; i++ {
		go worker(ctx, i, result) // контекст передается первым аргументом
	}

	foundBy := <-result // дожидаемся первого результата
	fmt.Println("result found by", foundBy)
	finish() // поскольку мы его дождались - дальше вызываем функцию отмены
	time.Sleep(time.Second)
}

func Example2() {
	workTime := 50 * time.Millisecond
	ctx, _ := context.WithTimeout(context.Background(), workTime)
	result := make(chan int, 1)

	for i := 0; i <= 10; i++ {
		go worker(ctx, i, result)
	}

	totalFound := 0
LOOP:
	for {
		select {
		case <-ctx.Done():
			break LOOP
		case foundBy := <-result:
			totalFound++
			fmt.Println("result found by", foundBy)
		}
	}
}

func main() {
	Example2()
}
