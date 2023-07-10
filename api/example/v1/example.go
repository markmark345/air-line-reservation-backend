package v1

import (
	"fmt"
	"net/http"

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