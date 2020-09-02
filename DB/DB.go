package main

import (
	"database/sql"
	"fmt"
)
func getAllUsers(){
	db:=createConn()
	defer db.Close()
	var (
		id    int
		email string
		name  string
		pass  string
		login []Login
	)
	i := 0

	rows, err := db.Query("select * from users")
	for rows.Next() {
		err := rows.Scan(&id, &email, &name, &pass)
		if err != nil {
			fmt.Println(err)
		} else {
			i++
			fmt.Println("scan ", i)
		}
		login = append(login, Login{Email: email, Name: name, Pass: pass})

	}
	defer rows.Close()
	return login
}
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
