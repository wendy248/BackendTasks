package main

import (
	"Github/BackendTasks/controller"
	"Github/BackendTasks/models"
	"net/http"

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
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "golang web api",
		})
	})
	r.GET("/animal", controller.ReadData)
	r.POST("/animal", controller.UploadData)
	r.PUT("/animal", controller.UpdateData)

	r.Run()
}
