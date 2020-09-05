package main

import (
	"database/sql"
	"fmt"
)

//User is user but Author is user email
type User struct {
	ID    int
	Name  string
	Email string
	Pass  string
}

//Pages that are created.
type Pages struct {
	ID     string
	Book   string
	Page   string
	Author User.Email
	Cont   Contents
}

//Contents that are created.
type Books struct {
	ID   string
	Page Pages
}

//Contents that are created.
type Contents struct {
	ID      string
	Content string
}

func GetAllUsers() User{
	db:=createConn()
	defer db.Close()
	var (
		id    int
		email string
		name  string
		pass  string
		author string
		user []User
	)

	i := 0

	rows, err := db.Query("select * from userpage")
	for rows.Next() {
		err := rows.Scan(&id, &name, &email, &pass)
		if err != nil {
			fmt.Println(err)
		} else {
			i++
			fmt.Println("scan ", i)
		}
		user = append(user, User{Name: name, Email: email, Pass: pass})

	}
	defer rows.Close()
	return user
}

func InsertPage(u User, p Pages){
	fmt.Println("creating page")
	db:=createConn()
	stmt, err := db.Prepare("INSERT INTO pages(book, page, author) VALUES(?, ?, ?)")
		if err != nil {
			log.Fatal(err)
		}

		userstemp := Data{P: Page{Book: p.Book, Page: p.Page, Author: u.Email, Cont: p.Contents}}
		fmt.Println(userstemp)

		p := userstemp
		s, err := Save(p)
		if err != nil {
			log.Fatal(err)
		}

		res, err := stmt.Exec(s.P.Book, s.P.Page, s.P.Author, s.P.Cont)
		if err != nil {
			log.Fatal(err)
		}
		lastId, err := res.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}
		rowCnt, err := res.RowsAffected()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
		fmt.Println("reached query")

}
func InsertBook(b Books){
	fmt.Println("creating book")
	db:=createConn()
	stmt, err := db.Prepare("INSERT INTO Books(book, page, author) VALUES(?, ?, ?)")
		if err != nil {
			log.Fatal(err)
		}

		userstemp := Data{P: Books{Book: b.Page.Book, Page: b.Page.Page, Author: b.Page.Author, Cont: b.Page.Cont}}
		fmt.Println(userstemp)

		p := userstemp
		s, err := Save(p)
		if err != nil {
			log.Fatal(err)
		}

		res, err := stmt.Exec(s.P.Book, s.P.Page, s.P.Author, s.P.Cont)
		if err != nil {
			log.Fatal(err)
		}
		lastId, err := res.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}
		rowCnt, err := res.RowsAffected()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
		fmt.Println("reached query")

}
func InsertContent(p Pages){
	fmt.Println("creating book")
	db:=createConn()
	stmt, err := db.Prepare("INSERT INTO Content(pageid, content) VALUES(?, ?, ?)")
		if err != nil {
			log.Fatal(err)
		}

		userstemp := Data{C: Contents{PageID: p.ID, Content: p.Contents}}
		fmt.Println(userstemp)

		p := userstemp
		s, err := Save(p)
		if err != nil {
			log.Fatal(err)
		}

		res, err := stmt.Exec(s.P.Cont)
		if err != nil {
			log.Fatal(err)
		}
		lastId, err := res.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}
		rowCnt, err := res.RowsAffected()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
		fmt.Println("reached query")

}
func createConn()return *DB{
	//opening database
	fmt.Println("db begin")
	db, err := sql.Open("mysql", "root:@/user")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("open ")
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("ping ")
	}
}
