package examples

import (
	"air-line-reservation-backend/models/example"
	"sync"
)

type ExampleService struct {
	// Database instance goes here
}

var exampleService *ExampleService
var initOnce sync.Once

func GetService() *ExampleService {
	initOnce.Do(initService)
	return exampleService
}

func initService() {
	exampleService = NewExampleService()
}

func NewExampleService() *ExampleService {
	return &ExampleService{}
}

func (dl *ExampleService) GetAllExamples() ([]example.ExampleModel, error) {
	// var examples []example.ExampleModel
	result := []example.ExampleModel{
		{ExampleId: 1, ExampleName: "Example 1"},
		{ExampleId: 2, ExampleName: "Example 2"},
		{ExampleId: 3, ExampleName: "Example 3"},
	}

	if result == nil {
		return nil, nil
	}

	return result, nil
}