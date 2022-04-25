package chapter2

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	year := now.Year()
	fmt.Println("Current datetime:", now)
	fmt.Println("Currnet year", year)
}
