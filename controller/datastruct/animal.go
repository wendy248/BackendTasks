package controller

type AnimalInput struct {
	Name  string `json:"name" binding:"required"`
	Class string `json:"class" binding:"required"`
	Legs  int16  `json:"legs" binding:"required"`
}