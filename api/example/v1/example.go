package v1

import (
	"fmt"
	"net/http"

	"air-line-reservation-backend/entities/example"

	"github.com/gin-gonic/gin"

	"air-line-reservation-backend/pkg/examples"
)

type RESTApiCtx struct {
	router *gin.Engine
}

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
	var example example.Example
	if err := c.ShouldBindJSON(&example); err != nil {
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