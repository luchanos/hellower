package chapter1

import "fmt"

func main() {
	originalCount, eatenCount := 10, 4
	fmt.Println("There are", originalCount, "apples")
	fmt.Println("I started with", eatenCount, "apples")
	fmt.Println("some jerk ate", originalCount-eatenCount, "apples")
}
