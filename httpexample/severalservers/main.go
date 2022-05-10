package main

import (
	"fmt"
	"net/http"
)

func runServer(addr string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Addr", addr, "URL", r.URL.String())
	})

	server := http.Server{
		Addr:    addr,
		Handler: mux,
	}
	fmt.Println("starting a new server at", addr)
	server.ListenAndServe()
}

func main() {
	go runServer(":8081")
	go runServer(":8080")
	fmt.Scanln()
}
