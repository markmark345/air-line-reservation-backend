package examples

import (
	"air-line-reservation-backend/dbs"
	"air-line-reservation-backend/entities/example"

	"sync"

	"gorm.io/gorm"
)

type ExampleService struct {
	db *gorm.DB
}

var exampleService *ExampleService
var initOnce sync.Once

func GetService() *ExampleService {
	initOnce.Do(initService)
	return exampleService
}

func initService() {
	db := dbs.GetDB()
	exampleService = NewExampleService(db)
}

func NewExampleService(db *gorm.DB) *ExampleService {
	return &ExampleService{
		db: db,
	}
}

func (dl *ExampleService) GetAllExamples() ([]example.Example, error) {
	var examples []example.Example
	result := dl.db.Find(&examples)

	if result.Error != nil {
		return nil, result.Error
	}

	return examples, nil
}

func (dl *ExampleService) CreateExample(example *example.Example) error {
	result := dl.db.Create(example)
	if result.Error != nil {
		return result.Error
	}
	return nil
}