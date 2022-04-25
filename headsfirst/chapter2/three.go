package chapter2

import "fmt"

func main() {
	a, b := "a", "b"
	var c, d string
	if true {
		c = "c"
		if true {
			d = "d"
			fmt.Println(a)
			fmt.Println(b)
			fmt.Println(c)
			fmt.Println(d)
		}
		fmt.Println(a)
		fmt.Println(b)
		fmt.Println(c)
		fmt.Println(d)
	}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
