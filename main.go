package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"time"
)

func simpleFunc(ctx context.Context) {
	for {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("New message")
		}
	}
}

func main() {
	ctx, finish := context.WithCancel(context.Background())

	go simpleFunc(ctx)
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Введите останов ")
		data, _ := reader.ReadString('\n')
		if data == "n\n" {
			finish()
			break
		}
	}
	fmt.Println("Программа завершена")
}
