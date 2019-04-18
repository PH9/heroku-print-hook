package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.POST("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "{}")

		b, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			log.Fatal("Could not read JSON body", err)
			return
		}

		body := string(b)
		log.Print(body)
	})

	router.Run(":" + port)
}
