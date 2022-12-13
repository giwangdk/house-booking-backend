package main

import (
	"final-project-backend/db"
	"final-project-backend/server"
	"fmt"
)

func main() {
	dbErr := db.Connect()
	if dbErr != nil {
		fmt.Println("error while connecting to database", dbErr)
	}
	server.Init()
}
