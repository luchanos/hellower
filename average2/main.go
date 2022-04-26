package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Println(os.Args[1:])
	numners := os.Args[1:]
	//var mySlice []int64
	//for _, val := range numners {
	//	int_val, err := strconv.ParseInt(val, 10, 64)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	mySlice = append(mySlice, int_val)
	//}
	//fmt.Println(mySlice)

	var minVal int64 = 0
	for _, val := range numners {
		int_val, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		if int_val < minVal {
			minVal = int_val
		}
	}
	fmt.Println(minVal)
}
