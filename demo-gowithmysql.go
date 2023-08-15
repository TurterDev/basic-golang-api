// go mod init github.com/TurterDev/go4web
// go get -u github.com/go-sql-driver/mysql
package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func creatingTable(db *sql.DB) {
	query := `CREATE TABLE users(
		id INT AUTO_INCREMENT,
		username TEXT NOT NULL,
		password TEXT NOT NULL,
		created_at DATETIME,
		PRIMARY KEY (id)
	)`
	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}
}
func Insert(db *sql.DB) {
	var username string
	var password string
	fmt.Scan(&username)
	fmt.Scan(&password)
	createdAt := time.Now()

	result, err := db.Exec(`INSERT INTO users (username,password, created_at) VALUES (?,?,?)`, username, password, createdAt)
	if err != nil {
		log.Fatal(err)
	}

	id, err := result.LastInsertId()
	fmt.Println(id)
}

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

func delete(db *sql.DB) {
	var deletedid int
	fmt.Scan(&deletedid)
	_, err := db.Exec(`DELETE FROM users WHERE id = ?`, deletedid)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/databasename")
	if err != nil {
		fmt.Println("failed to connect")
	} else {
		fmt.Println("connect successfully")
	}
	// fmt.Println(db)
	// query(db)
	// creatingTable(db)
	// Insert(db)
	delete(db)
}
