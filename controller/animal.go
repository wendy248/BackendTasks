package controller

import (
	"Github/BackendTasks/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type AnimalInput struct {
	Name  string `json:"name" binding:"required"`
	Class string `json:"class" binding:"required"`
	Legs  int16  `json:"legs" binding:"required"`
}

func UploadData(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var Input AnimalInput
	err := c.ShouldBindJSON(&Input)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error %s, message: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})

	} else {
		animal := models.Animal{
			Name:  Input.Name,
			Class: Input.Class,
			Legs:  Input.Legs,
		}

		db.Create(&animal)
		c.JSON(http.StatusOK, gin.H{
			"Info": "Success to insert data",
			"Data": animal,
		})
	}
}
