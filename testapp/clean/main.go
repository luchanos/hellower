package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

type MyServer struct {
	clientsMap map[uuid.UUID]MyClientStruct
}

func CreateMyServer() MyServer {
	return MyServer{
		clientsMap: make(map[uuid.UUID]MyClientStruct), // надо обязательно инитить
	}
}

func (s *MyServer) Unsubscribe() http.HandlerFunc {
	f := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Stop subscription for id {if}", r.URL.String())
	}
	return http.HandlerFunc(f)
}

func (s *MyServer) Subscribe(param uuid.UUID) http.HandlerFunc {
	f := func(w http.ResponseWriter, r *http.Request) {
		client := MyClientStruct{"channel"}
		s.clientsMap[param] = client
		fmt.Fprintln(w, "Start subscription!", r.URL.String())
		return
	}
	return http.HandlerFunc(f)
}

// MyClientStruct какая-то структура, обозначающая клиента
type MyClientStruct struct {
	subCh string // указатель
}

func (c *MyClientStruct) Stop() error {
	if err := c.subCh.Close(); err != nil {
		return err
	}
	return nil
}

func setMap(ctx context.Context) context.Context {
	context.WithValue(ctx, "subscribed", map[uuid.UUID]MyClientStruct{})
	return ctx
}

func main() {
	id := uuid.New()
	myClients := map[uuid.UUID]MyClientStruct{}
	someClient := MyClientStruct{}
	myClients[id] = someClient
	ctx := context.Background()
	ctx = setMap(ctx)

	s := CreateMyServer()
	mux := http.NewServeMux()

	mux.HandleFunc("/a", s.Subscribe(id))
	mux.HandleFunc("/b", s.Unsubscribe())

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	fmt.Println("starting a new server at", server.Addr)
	server.ListenAndServe()
}
