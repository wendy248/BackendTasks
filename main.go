package main

import (
	"Github/BackendTasks/config"
	"Github/BackendTasks/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := config.SetupModels()

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	v1 := r.Group("/v1")

	v1.GET("/animal", controller.ReadData)
	v1.GET("/animal/:id", controller.ReadDataByID)
	v1.POST("/animal", controller.UploadData)
	v1.PUT("/animal", controller.UpdateData)
	v1.DELETE("/animal/:name", controller.DeleteData)

	r.Run()
}
