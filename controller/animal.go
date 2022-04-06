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

//Read Data
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

func ReadDataByID(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var animal models.Animal

	if err := db.Where("id = ?", c.Param("id")).Find(&animal).Error; err != nil {
		errorMessage := fmt.Sprintf("ID %s does not exist in database", c.Param("id"))
		c.JSON(http.StatusNotFound, gin.H{
			"error": errorMessage,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": animal,
		})
	}
}

//Create Data / Upload Data
func UploadData(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var animal models.Animal
	var dataInput AnimalInput

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

// Update Data
func UpdateData(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var animal models.Animal
	var dataInput AnimalInput

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

func DeleteData(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var animal models.Animal

	a := c.Param("name")
	if err := db.Where("name = ?", a).Find(&animal).Error; err != nil {
		errorMessage := fmt.Sprintf("Data %s does not exist in database", a)
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
