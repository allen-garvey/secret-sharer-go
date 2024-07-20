package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand/v2"
	"net/http"
	"strconv"
)

type Secret struct {
	Title   string `form:"title" binding:"required"`
	Content string `form:"content" binding:"required"`
}

func main() {
	secrets := map[int]Secret{}
	const SITE_TITLE = "Secret Sharer"

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/assets", "./assets")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": SITE_TITLE,
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
			"title":   SITE_TITLE,
			"itemUrl": fmt.Sprintf("/items/%d", key),
		})
	})

	router.GET("/items/:key", func(c *gin.Context) {
		keyParam := c.Param("key")

		key, err := strconv.Atoi(keyParam)
		if err != nil {
			c.String(http.StatusNotFound, "Not found")
			return
		}

		secret, keyExists := secrets[key]

		if !keyExists {
			c.String(http.StatusNotFound, "Not found")
			return
		}

		c.HTML(http.StatusOK, "show.tmpl", gin.H{
			"title":       SITE_TITLE,
			"itemTitle":   secret.Title,
			"itemContent": secret.Content,
		})
	})

	router.Run(":3000")
}
