package controller

import (
	"Github/BackendTasks/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
)

// Update Data
func UpdateData(c *gin.Context) {
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

	if err2 := db.Where("name = ?", dataInput.Name).Find(&animal).Error; err2 != nil {
		anm := models.Animal{
			Name:  dataInput.Name,
			Class: dataInput.Class,
			Legs:  dataInput.Legs,
		}

		db.Create(&anm)
		c.JSON(http.StatusOK, gin.H{
			"message": "Data does not exist in database, so the program inserted the data to database",
			"Data":    anm,
		})

	} else {

		db.Model(&animal).Update(&dataInput)
		c.JSON(http.StatusOK, gin.H{
			"Message": "Successfull to Update Data",
			"Data":    animal,
		})
	}
}
