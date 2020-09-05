/*things to know about this.
1. behavior is just creating things
2. borrows a lot from DB.go
3. Pages has all other structs
*/
package main

import "fmt"

//User is user but Author is user email
type User struct {
	ID    int
	Name  string
	Email string
	Pass  string
}

//Pages that are created.
type Pages struct {
	ID     int
	Book   string
	Page   string
	Author User.Email
	Cont   Contents
}

//Books that are created.
type Books struct {
	ID   int
	Page Pages
}

//Contents that are created.
type Contents struct {
	ID      int
	PageID  int
	Content string
}

/*
! behavior of User package
* 1. creatPage()
* 2. createBook()
*/
func (u user) createPage(p Pages) string {
	users := getAllUsers()

	//check user
	if p.Page.Author == users.Email {
		//create the page
		DB.InsertPage(u, p)
	} else {
		fmt.Println("not the user")
	}

}

func (u user) createBook(b Books) string {
	users := getAllUsers()

	//check user
	if b.Page.Author == users.Email {
		//create the page
		DB.InsertBook(b)
	} else {
		fmt.Println("not the user")
	}
}

func (u user) createContent(p Pages) string {
	users := getAllUsers()

	//check user
	if p.Page.Author == users.Email {
		//create the page
		DB.InsertContent(p)
	} else {
		fmt.Println("not the user")
	}
}
