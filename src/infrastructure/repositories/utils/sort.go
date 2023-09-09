package utils

import (
	"github.com/javohir01/eater-service/scr/domain/address/models"
	"gorm.io/gorm"
)

func Sort(sort string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		orderBy := "created_at DESC"
		if sort == models.SortByCreatedAtAsc {
			orderBy = "created_at ASC"
		}
		return db.Order(orderBy)
	}
}
