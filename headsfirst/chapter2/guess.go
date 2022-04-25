package chapter2

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("I've chosen some number between 1 and 100")
	fmt.Println("Can you guess it?")

	reader := bufio.NewReader(os.Stdin)
	success := false
	fmt.Println("Make a guess!")

	for x := 0; x < 10; x++ {
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatalln(err)
		}

		if guess < target {
			fmt.Println("Oops! Too LOW!")
		} else if guess > target {
			fmt.Println("Oops! Too HIGH!")
		} else {
			fmt.Println("Well done!!")
			success = true
			break
		}
	}
	if !success {
		fmt.Println("Sorry, your attempts has been ended. The target was:", target)
	}
}
