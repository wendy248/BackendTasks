package main

import (
	"Github/BackendTasks/controller"
	"Github/BackendTasks/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//Models
	db := models.SetupModels()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.GET("/animal", controller.ReadData)
	r.POST("/animal", controller.UploadData)
	r.PUT("/animal", controller.UpdateData)
	r.DELETE("/animal/:name", controller.DeleteData)

	r.Run()
}
