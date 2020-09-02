package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	Contextor "github.com/golangast/pager/Contexter"
	"github.com/rs/cors"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/post", run)
	handler := cors.Default().Handler(mux)
	c := context.Background()
	log.Fatal(http.ListenAndServe(":8081", Contextor.AddContext(c, handler)))

}
func run(w http.ResponseWriter, r *http.Request) {
	fmt.Println("works")
}
