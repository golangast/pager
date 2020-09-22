/*things to know about this.
1. behavior is just creating things
2. borrows a lot from DB.go
3. Pages has all other structs
*/
package User

import (
	"fmt"

	DB "github.com/golangast/pager/DB"
)

//User is user but Author is user email
type User struct {
	ID    int
	Name  string
	Email string
	Pass  string
	Library
}

//User is user but Author is user email
type Pages struct {
	ID      int
	Name    string
	URL     string
	Content string
}

//Pages that are created.
type Library struct {
	ID     int
	Book   string
	Page   string
	Author string
	Books
}

//Contents that are created.
type Books struct {
	ID     int
	PageID string
	Name   string
	Author string
	Pages
}

/*
! behavior of User package
* 1. creatPage()
* 2. createBook()
* 3. createContent()
*/
func CreatePage(u User, Name string, URL string, c string) {
	users := DB.GetAllUsers()

	for _, v := range users {
		//check user

		if User.Email == v.Email {
			//create the page
			DB.InsertPage(Name, URL, c)
		} else {
			fmt.Println("not the user")
		}
	}

}

func CreateBook(u User, b Books) {
	users := DB.GetAllUsers()
	for _, v := range users {
		//check user
		if User.Email == v.Email {
			//create the page
			DB.InsertBook(b)
		} else {
			fmt.Println("not the user")
		}
	}
}

func CreateContent(u User, pageid int, content string) {
	users := DB.GetAllUsers()

	for _, v := range users {
		//check user
		if User.Email == v.Email {

			//create the Content
			DB.UpdateContent(pageid, content)
		} else {
			fmt.Println("not the user")
		}
	}
}
