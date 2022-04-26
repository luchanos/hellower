package main

import (
	"fmt"
	"hellower/datafile"
	"log"
)

func main() {
	numbers, err := datafile.GetFloats("/Users/nnsviridov/go/src/hellower/average/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	sampleCount := float64(len(numbers))
	fmt.Printf("Average: %0.2f\n", sum/sampleCount)
	//var notes []string
	//notes = []string{"do", "re", "mi", "fa", "so", "la", "ti"}
	//someSlice := notes[1:3]
	//fmt.Println(someSlice)
	//anotherSlice := notes[2:5]
	//fmt.Println(anotherSlice)
	//notes[2] = "XXXX"
	//fmt.Println(someSlice)
	//fmt.Println(anotherSlice)
	//anotherSlice[0] = "AAAA"
	//fmt.Println(notes)
	//notes = append(notes, "RRRRR")
	//anotherSlice = append(anotherSlice, "TEST")
	//fmt.Println(notes)
	//fmt.Println(anotherSlice)
}
