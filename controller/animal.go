package controller

import (
	"Github/BackendTasks/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
)

type AnimalInput struct {
	Name  string `json:"name" binding:"required"`
	Class string `json:"class" binding:"required"`
	Legs  int16  `json:"legs" binding:"required"`
}

type AnimalUpdate struct {
	Name  string `json:"name" binding:"required"`
	Class string `json:"class"`
	Legs  int16  `json:"legs"`
}

//Read Data
func ReadData(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var animal []models.Animal

	db.Find(&animal)
	c.JSON(http.StatusOK, gin.H{
		"data": animal,
	})
}

//Create Data / Upload Data
func UploadData(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//Validate Input
	var animal models.Animal
	var dataInput AnimalInput

	err := c.ShouldBindJSON(&dataInput)
	err2 := db.Where("name = ?", dataInput.Name).Find(&animal).Error

	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error %s, message: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})
		return
	}

	if err2 == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Data already exist in database",
		})
		return
	}

	//Process Input
	mhs := models.Animal{
		Name:  dataInput.Name,
		Class: dataInput.Class,
		Legs:  dataInput.Legs,
	}

	db.Create(&mhs) //Create DB MySQL
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully insert data",
		"Data":    mhs,
	})
}

// Update Data
func UpdateData(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//Validate Data
	var animal models.Animal
	var inputAnimal AnimalUpdate

	err := c.ShouldBindJSON(&inputAnimal)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Update Data
	db.Model(&animal).Update(&inputAnimal)
	c.JSON(http.StatusOK, gin.H{
		"Message": "Successfull to Update Data",
		"Data":    animal,
	})
}
