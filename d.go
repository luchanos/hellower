package main

import (
	"fmt"
	"hellower/theirs/github.com/luchanos/hellower/keyboard"
	"log"
)

func d() {
	var status string
	fmt.Print("Enter a grade: ")

	num, err := keyboard.GetFloat()
	if err != nil {
		log.Fatalln(err)
	}
	if num >= 60 {
		status = "Passed!"
	} else {
		status = "Not passed!"
	}
	fmt.Println(status)
}
