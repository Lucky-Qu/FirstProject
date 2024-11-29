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
		c.JSON(200, gin.H{
			"msg": "查询成功",
		})
	})
	g.DELETE("/FirstProject", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "删除成功",
		})
	})
	g.PUT("/FirstProject", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "更改成功",
		})
	})
	err = g.Run(config.URL)
	if err != nil {
		panic(err)
	}
}
