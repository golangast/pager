package main

import (
	"context"
	"log"
	"net/http"

	"github.com/golangast/pager/Contextor"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/post", User.createPage())
	handler := cors.Default().Handler(mux)
	c := context.Background()
	log.Fatal(http.ListenAndServe(":8081", Contextor.AddContext(c, handler)))

}
