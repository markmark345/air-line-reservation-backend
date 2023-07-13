package example

import "gorm.io/gorm"

type Example struct {
	gorm.Model
	ExampleId   int    `json:"ExampleId" validate:"required"`
	ExampleName string `json:"ExampleName" validate:"min=1,max=12"`
}