package controller

import (
	"Github/BackendTasks/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)


// Delete Data
func DeleteData(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var animal models.Animal

	if err := db.Where("name = ?", c.Param("name")).Find(&animal).Error; err != nil {
		errorMessage := fmt.Sprintf("Data %s does not exist in database", c.Param("name"))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessage,
		})
		return
	}

	db.Delete(&animal)
	c.JSON(http.StatusOK, gin.H{
		"Message": "Data successfully deleted",
	})
}
