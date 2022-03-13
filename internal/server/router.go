package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"real_estate_finder/real_estate_finder/internal/models"

	"github.com/gin-gonic/gin"
)

func NewRouter(token string) *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery()) // recovers from any panics and writes a 500
	router.Use(gin.Logger())

	router.GET("/ping", func(c *gin.Context) {
		c.IndentedJSON(200, gin.H{
			"message": "pong",
		})
	})

	router.POST("/midpoint", func(c *gin.Context) {
		b, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.IndentedJSON(500, gin.H{"error reading body": err})
		}

		var p []models.Point
		if err = json.Unmarshal(b, &p); err != nil {
			c.IndentedJSON(500, gin.H{"unmarshalling error": err})
		}

		mid, err := models.Midpoint(p)
		if err != nil {
			c.IndentedJSON(500, gin.H{
				"error": err.Error(),
			})
		} else {
			c.IndentedJSON(200, mid)
		}
	})

	router.GET("/map", func(c *gin.Context) {
		c.HTML(
			http.StatusOK,
			"mapbasics.html",
			gin.H{
				"token": token,
			},
		)
	})

	return router
}
