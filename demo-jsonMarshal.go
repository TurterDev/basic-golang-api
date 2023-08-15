//ข้อควรระวัง json.Marshal จะอ่านค่าเฉพาะตัวแปลที่ขึ้นต้นด้วยตัวพิมพ์ใหญ่เท่านั้น
package main

import (
	"encoding/json"
	"fmt"
)

type employee struct {
	Id int
	EmployeeName string
	Tel string
	Email string
}

func main() {
	data, _ := json.Marshal(&employee{101,"Turter","1234567890","turterdev@mail.com"})
	fmt.Println(string(data))
}