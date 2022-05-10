package main

import (
	"fmt"
	"log"
	"net/http"
)

func ErrorHandler(err error, msg string) {
	if err != nil {
		log.Fatalf("%s %s", err, msg)
	}

}

func handler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "Привет, мир!")
	ErrorHandler(err, "smth went wrong!")
	_, err = w.Write([]byte("!!!"))
	ErrorHandler(err, "smth went wrong!")
}

func Example1() {
	http.HandleFunc("/", handler) // привязываем роут

	http.HandleFunc("/page",
		func(w http.ResponseWriter, r *http.Request) {
			_, err := fmt.Fprintln(w, "Single Pages: ", r.URL.String())
			ErrorHandler(err, "Smth went wrong /page handler")
		})

	http.HandleFunc("/pages/",
		func(w http.ResponseWriter, r *http.Request) {
			_, err := fmt.Fprintln(w, "Multiple Pages: ", r.URL.String())
			ErrorHandler(err, "Smth went wrong /pages/ handler")
		})

	fmt.Println("starting server at :8080")
	err := http.ListenAndServe(":8080", nil)
	ErrorHandler(err, "smth went wrong!")
}

func main() {
	Example1()
}
