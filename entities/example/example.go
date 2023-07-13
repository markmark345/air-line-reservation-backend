package example

import "gorm.io/gorm"

type Example struct {
	gorm.Model
	ExampleId   int    `json:"example_id"`
	ExampleName string `json:"example_name"`
}