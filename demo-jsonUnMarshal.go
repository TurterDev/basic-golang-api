package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type employee struct {
	ID int
	EmployeeName string
	Tel string
	Email string
}

func main() {
	e := employee{}
	err := json.Unmarshal([]byte(`{"ID":101,"EmployeeName":"Turter","Tel":"1234567890","Email":"turterdev@mail.com"}`), &e)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(e)
	fmt.Println(e.EmployeeName)
}