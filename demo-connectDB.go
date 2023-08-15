//go mod init github.com/TurterDev/go4web
//go get -u github.com/go-sql-driver/mysql
package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func query(db *sql.DB) {
	var (
		id         int
		coursename string
		price      float64
		instructor string
	)

	var inputID int
	fmt.Scan(&inputID)

	query := "SELECT id, coursename, price, instructor FROM table WHERE id = ?"
	if err := db.QueryRow(query, inputID).Scan(&id, &coursename, &price, &instructor); err != nil {
		log.Fatal(err)
	}
	fmt.Println(id, coursename, price, instructor)
}

func main() {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/databasename")
	if err != nil {
		fmt.Println("failed to connect")
	} else {
		fmt.Println("connect successfully")
	}
	// fmt.Println(db)
	query(db)
}
