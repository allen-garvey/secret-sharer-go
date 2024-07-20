package main

import (
	"net/http"
	"fmt"
	"log"
	"github.com/gin-gonic/gin"
	"math/rand/v2"
)

type Secret struct {
	Title   string `form:"title" binding:"required"`
	Content string `form:"content" binding:"required"`
}

func main() {
	secrets := map[int]Secret{}
	
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/assets", "./assets")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Secret Sharer",
		})
	})

	router.POST("/items", func(c *gin.Context) {
		var createParams Secret
		err := c.Bind(&createParams)
		if err != nil {
			log.Fatal(err)
		}

		key := rand.IntN(999) + 1
		secrets[key] = createParams

		c.HTML(http.StatusOK, "created.tmpl", gin.H{
			"title": "Secret Sharer",
			"itemUrl": fmt.Sprintf("/items/%d", key),
		})
	})
	
	router.Run(":3000")
}
