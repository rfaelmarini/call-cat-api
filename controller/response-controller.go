package controller

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rfaelmarini/call-cat-api/entity"
	"github.com/rfaelmarini/call-cat-api/service"
)

type ResponseController interface {
	FindAll(ctx *gin.Context) (entity.Response, error)
}

type controller struct {
	service service.ResponseService
}

func New(service service.ResponseService) ResponseController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll(ctx *gin.Context) (entity.Response, error) {
	response := entity.Response{}
	response.RequestedURL = ""
	response.Body = "[]"
	response.StatusCode = 200
	name, ok := ctx.GetQuery("name")
	if !ok {
		return response, nil
	}

	apiKey := os.Getenv("API_KEY")
	url := "https://api.thecatapi.com/v1/breeds/search?api_key=" + apiKey + "&q=" + name

	response = c.service.Find(url)
	if response.RequestedURL != "" {
		return response, nil
	}

	resp, err := http.Get(url)
	if err != nil {
		return response, err
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}

	response.RequestedURL = url
	response.Body = string(body)
	response.StatusCode = resp.StatusCode

	c.service.Save(response)
	return response, nil
}
