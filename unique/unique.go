package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	myStrings := []string{}
	marker := false
	for {
		fmt.Println("Введите новую строку: ")
		s, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("panic %s", err)
		}

		for _, el := range myStrings {
			if el == s {
				marker = true
				break
			}
		}
		if !marker {
			myStrings = append(myStrings, s)
			fmt.Println(myStrings)
		}
		fmt.Print("Продолжить ввод? [y/n] ")
		s, err = reader.ReadString('\n')
		if err != nil {
			log.Fatalf("panic %s", err)
		}
		if s != "y\n" {
			break
		}
		marker = false
	}
}
