package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func setEnvVariables() {
	os.Setenv("API_KEY", "1a9c1e22-9dc7-48fa-844c-5d137e80694")
}

func main() {
	setEnvVariables()
	server := gin.Default()
	server.GET("/breeds", func(ctx *gin.Context) {
		name, ok := ctx.GetQuery("name")
		if !ok {
			ctx.JSON(http.StatusOK, []string{})
			ctx.Abort()
			return
		}

		apiKey := os.Getenv("API_KEY")
		url := "https://api.thecatapi.com/v1/breeds/search?api_key=" + apiKey + "&q=" + name
		resp, err := http.Get(url)
		if err != nil {
			log.Fatalln(err)
		}

		if resp.Body != nil {
			defer resp.Body.Close()
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		ctx.JSON(http.StatusOK, json.RawMessage(string(body)))
	})

	server.Run(":8080")
}
