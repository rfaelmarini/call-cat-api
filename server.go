package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rfaelmarini/call-cat-api/controller"
	"github.com/rfaelmarini/call-cat-api/repository"
	"github.com/rfaelmarini/call-cat-api/service"
)

var (
	responseRepository repository.ResponseRepository = repository.NewResponseRepository()
	responseService    service.ResponseService       = service.New(responseRepository)
	responseController controller.ResponseController = controller.New(responseService)
)

func setEnvVariables() {
	os.Setenv("API_KEY", "1a9c1e22-9dc7-48fa-844c-5d137e80694")
}

func main() {
	setEnvVariables()
	server := gin.Default()
	server.GET("/breeds", func(ctx *gin.Context) {
		response, err := responseController.FindAll(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			ctx.Abort()
		}

		ctx.JSON(response.StatusCode, json.RawMessage(response.Body))
	})

	server.Run(":8080")
}
