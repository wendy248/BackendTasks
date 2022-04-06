package models

type Animal struct {
	ID    int    `json:"id" gorm:"primary_key AUTO_INCREMENT NOT_NULL"`
	Name  string `json:"name" gorm:"NOT_NULL"`
	Class string `json:"class" gorm:"NOT_NULL"`
	Legs  int16  `json:"legs" gorm:"NOT_NULL"`
}
