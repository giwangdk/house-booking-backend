package helper

import (
	"final-project-backend/db"
	"final-project-backend/entity"
	"fmt"
	"time"
)

func GoCronScheduler() {
	db:= db.Get()

	var reservationID []string

	err:= db.Model(entity.Reservation{}).Select("id").Where("status_id = 1 AND expired < ?", time.Now()).Find(&reservationID).Error
	fmt.Println(reservationID)
	if err != nil {
		fmt.Println(err.Error())
	}

	if len(reservationID) <= 0 {
		return
	}

	err= db.Model(entity.Reservation{}).Where("id IN (?)", reservationID).Update("status_id", 3).Error
		if err != nil {
			fmt.Println(err.Error())
		}
}