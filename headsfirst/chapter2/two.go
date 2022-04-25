package chapter2

import (
	"fmt"
	"strings"
)

func main() {
	broken := "test## # #s#tring"
	replacer := strings.NewReplacer("#", "")
	fixed := replacer.Replace(broken)
	fmt.Println(fixed)
}
