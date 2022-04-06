package controller

import (
	"Github/BackendTasks/controller/datastruct"
	"Github/BackendTasks/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
)


//Insert / Upload Data to database
func UploadData(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var animal models.Animal
	var dataInput datastruct.AnimalInput

	if err := c.ShouldBindJSON(&dataInput); err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("%s field is %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})
		return
	}

	if err2 := db.Where("name = ?", dataInput.Name).Find(&animal).Error; err2 == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Denied. Animal name is already exist in database",
		})
		return
	}

	anm := models.Animal{
		Name:  dataInput.Name,
		Class: dataInput.Class,
		Legs:  dataInput.Legs,
	}

	db.Create(&anm)
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully insert data",
		"Data":    anm,
	})
}
