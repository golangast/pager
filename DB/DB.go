package main

import (
	"database/sql"
	"fmt"
)
type User struct {
	Id int
	Name  string
	Email string
	Pass string
	Author string
}
//Page that are created.
type Page struct {
	ID    string
	Book  string
	Page  string
	Users User
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
		err := rows.Scan(&id, &name, &email, &pass, &author)
		if err != nil {
			fmt.Println(err)
		} else {
			i++
			fmt.Println("scan ", i)
		}
		user = append(user, User{Name: name, Email: email, Pass: pass, Author: author})

	}
	defer rows.Close()
	return user
}

func InsertPage(u User, p Page){
	fmt.Println("creating page")
	db:=createConn()
	stmt, err := db.Prepare("INSERT INTO pages(book, page, author) VALUES(?, ?, ?)")
		if err != nil {
			log.Fatal(err)
		}

		userstemp := Data{P: Page{Book: p.Book, Page: p.Page, Author: p.Users.Author}}
		fmt.Println(userstemp)

		u := userstemp
		s, err := Save(u)
		if err != nil {
			log.Fatal(err)
		}

		res, err := stmt.Exec(s.P.Name, s.P.Email, s.P.Pass)
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
