package errorhandler

import "fmt"

func ErrorHandler(err error, msg string) {
	if err != nil {
		fmt.Printf("%s %s", err, msg)
	}
}
