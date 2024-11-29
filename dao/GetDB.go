package dao

import (
	"FirstProject/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func ConnectDB() *gorm.DB {
	DB, err := gorm.Open(mysql.New(mysql.Config{
		DSN: config.DSN,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			TablePrefix:   "FP_",
		},
	})
	if err != nil {
		panic("failed to connect database")
	}
	initDB(DB)
	return DB
}
