package model

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Num  string `json:"num"`
	Name string `json:"name"`
	Age  int    `json:"age" binding:"required,AgeCheck"`
	Sex  string `json:"sex"`
}

func AgeCheck(fl validator.FieldLevel) bool {
	if age, ok := fl.Field().Interface().(int); ok {
		if age < 0 || age > 100 {
			return false
		}
		return true
	}
	return false
}
