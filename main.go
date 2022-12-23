package main

import (
	"final-project-backend/db"
	"final-project-backend/helper"
	"final-project-backend/server"
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

func main() {
	dbErr := db.Connect()

	
	if dbErr != nil {
		fmt.Println("error while connecting to database", dbErr)
	}
	s := gocron.NewScheduler(time.Local)
	_, _ = s.Every(10).Seconds().Do(helper.GoCronScheduler)

	defer s.Stop()

	go s.StartAsync()
	server.Init()
}
