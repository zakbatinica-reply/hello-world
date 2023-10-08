package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("views/index.html")
	r.Static("/css", "./css")

	r.GET("/", func(c *gin.Context) {
		headers := make(map[string]string, len(c.Request.Header))
		for k, v := range c.Request.Header {
			headers[k] = v[0]
		}

		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":   "Hello, World!",
			"headers": headers,
		})
	})

	r.Run()
}
