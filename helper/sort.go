package helper

import "gorm.io/gorm"

func SortBy(sortBy string, sort string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if sort == "desc" {
			return db.Order(sortBy + " " + sort)
		} else {
			return db.Order(sortBy + " " + sort)
		}
	}
}