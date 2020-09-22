package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	Contextor "github.com/golangast/pager/Contexter"
	Users "github.com/golangast/pager/User"

	"github.com/rs/cors"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/me", run)
	handler := cors.Default().Handler(mux)
	c := context.Background()
	log.Fatal(http.ListenAndServe(":8081", Contextor.AddContext(c, handler)))
}
func run(w http.ResponseWriter, r *http.Request) {

	us := User{ID: 1, Name: "jim", Email: "jim@yahoo.com", Pass: "jim", Library: Library{ID: 2, Books: Books{ID: 3, Pages: Pages{ID: 2, Content: "something"}}}}}
	fmt.Println(us)
	Users.CreateContent(us.Library.Books.Pages.ID, us.Library.Books.Pages.Contents.Content)
}
