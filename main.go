package main

import (
	"context"
	"log"
	"net/http"
	
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/post", DB.POST)
	handler := cors.Default().Handler(mux)
	c := context.Background()
	log.Fatal(http.ListenAndServe(":8081", Contextor.AddContext(c, handler)))

}
