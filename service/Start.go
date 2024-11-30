package service

import (
	"FirstProject/config"
	"FirstProject/dao"
	"FirstProject/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func StartService() {
	var err error
	db := dao.ConnectDB()
	g := gin.Default()
	g.POST("/FirstProject", func(c *gin.Context) {
		student := model.Student{}
		if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
			err = v.RegisterValidation("AgeCheck", model.AgeCheck)
			if err != nil {
				panic(err)
			}
		}
		err = c.ShouldBindJSON(&student)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "年龄低于零或者大于一百",
			})
		} else {
			db.Create(&student)
			c.JSON(200, gin.H{
				"msg":     "增加成功",
				"student": student,
			})
		}
	})
	g.GET("/FirstProject", func(c *gin.Context) {
		name := c.Query("name")
		student := model.Student{}
		db.Model(&model.Student{}).Where("name = ?", name).Find(&student)
		c.JSON(200, gin.H{
			"msg":     "查询成功",
			"student": student,
		})
	})
	g.DELETE("/FirstProject/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "删除成功",
		})
	})
	g.PUT("/FirstProject", func(c *gin.Context) {
		student := model.Student{}
		num := c.Query("num")
		if db.Model(&model.Student{}).Where("num = ?", num).First(&student); student.Name == "" {
			panic("No")
		}
		if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
			err = v.RegisterValidation("AgeCheck", model.AgeCheck)
			if err != nil {
				panic(err)
			}
		}
		student2 := model.Student{}
		err = c.ShouldBindJSON(&student2)
		if err != nil {
			panic(err)
		}
		db.Model(&model.Student{}).Where(&student).Updates(&student2)
		c.JSON(200, gin.H{
			"msg":     "更改成功",
			"student": student2,
		})
	})
	err = g.Run(config.URL)
	if err != nil {
		panic(err)
	}
}
