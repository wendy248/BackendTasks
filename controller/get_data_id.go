package controller

import (
	"Github/BackendTasks/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//Read Data by ID from database
func ReadDataByID(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var animal models.Animal

	if err := db.Where("id = ?", c.Param("id")).Find(&animal).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Data not found",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": animal,
		})
	}
}
