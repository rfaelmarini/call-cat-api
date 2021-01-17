package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.GET("/breeds", func(ctx *gin.Context) {
		name, ok := ctx.GetQuery("name")
		if !ok {
			ctx.JSON(http.StatusOK, []string{})
			ctx.Abort()
			return
		}

		url := "https://api.thecatapi.com/v1/breeds/search?api_key=1a9c1e22-9dc7-48fa-844c-5d137e80694&q=" + name
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
