package main

//User who creates pages
type User struct {
	Id     int
	Name   string
	Email  string
	Pass   string
	Author string
}

//Page that are created.
type Page struct {
	ID    string
	Book  string
	Page  string
	Users User
}

func (u user) createPage(u User) string {
	users := getAllUsers()
	//dateabase stuff to add a page

	//check user
	if User.Email == users.Email {

	}
	//create the page

}
func (u user) createBook() string {
	//dateabase stuff to add a book
	//id,user,book
}
