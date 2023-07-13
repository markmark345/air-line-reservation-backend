package v1

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"air-line-reservation-backend/entities/example"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"air-line-reservation-backend/pkg/examples"
)

type RESTApiCtx struct {
	router *gin.Engine
}

// Caching Validator instance
var validate *validator.Validate

func (api *RESTApiCtx) Serve(addr string) error {
	return api.router.Run(addr)
}

func path(basePath, endpoint string) string {
	return fmt.Sprintf("%s/%s", basePath, endpoint)
}

func NewExampleRESTApi() *RESTApiCtx {
	basePath := "/v1/"
	router := gin.Default()
	api := &RESTApiCtx {
		router,
	}

	router.GET(path(basePath, "examples"), api.GetExamples)
	router.POST(path(basePath, "examples"), api.AddExample)

	return api
}


func (api *RESTApiCtx) GetExamples(c *gin.Context) {
	examples, err := examples.GetService().GetAllExamples()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch projects"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": examples,
	})
}

func (api *RESTApiCtx) AddExample(c *gin.Context) {
	// Validator
	validate = validator.New()

	// Register Reflection for TagName
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	
	var example example.Example
	if err := c.ShouldBindJSON(&example); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Using validator to valdiate the sturcted entities from request body
	if err := validate.Struct(example); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := examples.GetService().CreateExample(&example); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add new example"})
	}

	c.JSON(http.StatusOK, gin.H{
		"id": example.ID,
	})
}