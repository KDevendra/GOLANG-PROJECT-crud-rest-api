package main

import (
	"GOLANG-PROJECT-crud-rest-api/database"
	"fmt"
)

func main() {
	database.Connect()
	fmt.Println("Hello")
}
