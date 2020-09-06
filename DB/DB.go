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
//User is user but Author is user email
type Pages struct {
	ID    int
	Name  string
	URL string
	ContentID int
}

//Pages that are created.
type Library struct {
	ID     string
	Book   string
	Page   string
	Author User.Email
	Cont   Contents
}

//Contents that are created.
type Books struct {
	ID   string
	PageID string
	Name string
	Author string
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

	rows, err := db.Query("select * from user")
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

func InsertPage(p Pages){
	fmt.Println("creating page")
	db:=createConn()
	stmt, err := db.Prepare("INSERT INTO pages(Name, URL, ContentID) VALUES(?, ?, ?)")
		if err != nil {
			log.Fatal(err)
		}

		userstemp := Data{P: Pages{Name: p.Name, URL: p.URL, ContentID:p.ContentID}}
		fmt.Println(userstemp)

		l := userstemp
		s, err := Save(l)
		if err != nil {
			log.Fatal(err)
		}

		res, err := stmt.Exec(s.P.Name, s.P.URL, s.P.ContentID)
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
	stmt, err := db.Prepare("INSERT INTO Books(PageID, Name, Author) VALUES(?, ?, ?)")
		if err != nil {
			log.Fatal(err)
		}

		userstemp := Data{B: Books{PageID: b.PageID, Name: b.Name, Author: b.Author}}
		fmt.Println(userstemp)

		p := userstemp
		s, err := Save(p)
		if err != nil {
			log.Fatal(err)
		}

		res, err := stmt.Exec(s.B.PageID, s.B.Name, s.B.Author)
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
func InsertContent(pageid int, content string){
	fmt.Println("creating book")
	db:=createConn()
	stmt, err := db.Prepare("INSERT INTO Content(Content) VALUES(?)")
		if err != nil {
			log.Fatal(err)
		}

		userstemp := Data{C: Contents{PageID: pageid, Content: content}}
		fmt.Println(userstemp)

		P := userstemp
		s, err := Save(P)
		if err != nil {
			log.Fatal(err)
		}

		res, err := stmt.Exec(s.P.ID, s.P.Cont)
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
	db, err := sql.Open("mysql", "root:@/userpage")
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
