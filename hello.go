package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func getEnvironment() string {
	if _, present := os.LookupEnv("KUBERNETES_SERVICE_HOST"); present {
		return "Running in Kubernetes ☸️"
	}
	if _, err := os.Stat("/.dockerenv"); err == nil {
		return "Running in Docker 🐳"
	}
	return "Running Natively 💻"
}

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
			"title":       "Hello, World!",
			"environment": getEnvironment(),
			"headers":     headers,
		})
	})

	r.Run()
}
