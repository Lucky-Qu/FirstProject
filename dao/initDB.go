package dao

import (
	"FirstProject/model"
	"gorm.io/gorm"
)

func initDB(db *gorm.DB) {
	if !db.Migrator().HasTable(&model.Student{}) {
		err := db.Migrator().CreateTable(&model.Student{})
		if err != nil {
			panic(err)
		}
	}
}
