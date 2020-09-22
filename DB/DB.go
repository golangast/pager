package DB

import (
	"database/sql"
	"fmt"
	"log"
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
	ID     string
	Book   string
	Page   string
	Author string
	Books
}

//Contents that are created.
type Books struct {
	ID     string
	PageID string
	Name   string
	Author string
	Pages
}

type Data struct {
	P Pages
	B Books
	L Library
	S Sanitizer
}
type Sanitizer interface {
	Sanitize()
}

func Save(l Data) (Data, error) {
	var err error

	// type assertion for Sanitizer (could also use a type switch)
	s, ok := l.S.(Sanitizer)

	if !ok {
		if err != nil {
			log.Fatal(err)
		}
		// ... save without sanitization
		return l, err
	}

	s.Sanitize()
	return l, err
}
func GetAllUsers() []User {
	db := createConn()
	defer db.Close()
	var (
		id    int
		email string
		name  string
		pass  string

		user []User
	)

	i := 0

	rows, err := db.Query("select * from user")
	if err != nil {
		log.Fatal(err)
	}
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

func InsertPage(Name string, URL string, c string) {
	fmt.Println("creating page")
	db := createConn()
	stmt, err := db.Prepare("INSERT INTO pages(Name, URL, Content) VALUES(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	userstemp := Data{P: Pages{Name: Name, URL: URL, Content: c}}

	fmt.Println(userstemp)

	l := userstemp
	s, err := Save(l)
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(s.P.Name, s.P.URL, s.P.Content)
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
	log.Printf("affected = %d\n", lastId, rowCnt)
	fmt.Println("reached query")

}

func InsertBook(b Books) {
	fmt.Println("creating book")
	db := createConn()
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
func UpdateContent(pageid int, content string) {
	fmt.Println("creating book")
	db := createConn()

	stmt, err := db.Prepare("UPDATE pages SET Content=? WHERE ID=?")
	if err != nil {
		log.Fatal(err)
	}

	userstemp := Data{P: Pages{Content: content}}
	fmt.Println(userstemp)

	C := userstemp
	s, err := Save(C)
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(s.P.Content)
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
func createConn() *sql.DB {
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
	return db
}
