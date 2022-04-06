package controller

import (
	"Github/BackendTasks/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)


//Read Data from database
func ReadData(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var animal []models.Animal

	db.Find(&animal)

	if len(animal) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Error": "Data not found",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": animal,
		})
	}
}
