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
	Lib   Library
}

//User is user but Author is user email
type Pages struct {
	ID        int
	Name      string
	URL       string
	ContentID int
	Cont      Contents
}

//Pages that are created.
type Library struct {
	ID     string
	Book   string
	Page   string
	Author User.Email
	Book   Books
}

//Contents that are created.
type Books struct {
	ID     string
	PageID string
	Name   string
	Author string
	Page   Pages
}

//Contents that are created.
type Contents struct {
	ID      string
	Content string
}

/*
! behavior of User package
* 1. creatPage()
* 2. createBook()
*/
func (u user) createPage() {
	users := getAllUsers()

	//check user
	if p.Page.Author == users.Email {
		//create the page
		DB.InsertPage(u.Library.Books.Pages)
	} else {
		fmt.Println("not the user")
	}

}

func (u user) createBook() {
	users := getAllUsers()

	//check user
	if b.Page.Author == users.Email {
		//create the page
		DB.InsertBook(u.Library.Books)
	} else {
		fmt.Println("not the user")
	}
}

func (u user) createContent(pageid int, content string) {
	users := getAllUsers()

	//check user
	if p.Page.Author == users.Email {
		//create the Content

		DB.InsertContent(pageid, content)
	} else {
		fmt.Println("not the user")
	}
}
