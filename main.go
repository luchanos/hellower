package main

import (
	"fmt"
	"time"
)

func main() {
	var notes [7]string
	notes[0] = "do"
	notes[1] = "re"
	notes[2] = "mi"
	fmt.Println(notes[0])
	fmt.Println(notes[1])

	var dates [3]time.Time
	dates[0] = time.Unix(1257894000, 0)
	dates[1] = time.Unix(1257994000, 0)
	dates[2] = time.Unix(1257494000, 0)
	fmt.Println(dates[1])

	var primes [5]int
	primes[0] = 2
	fmt.Println(primes[0])
	fmt.Println(primes[2])
	fmt.Println(primes[4])

	var counters [3]int
	counters[0]++
	counters[0]++
	counters[2]++
	fmt.Println(counters[0], counters[1], counters[2])

	notes_new := [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	fmt.Println(notes_new[3], notes_new[6], notes_new[0])
	var primes_new [5]int = [5]int{2, 3, 4, 5, 6}
	fmt.Println(primes_new[0], primes_new[2], primes_new[4])

	for i := 0; i < len(notes_new); i++ {
		fmt.Println(notes_new[i])
	}

	for index, value := range notes_new {
		fmt.Println(index, value)
	}
}
