package server

import (
	"encoding/json"
	"io/ioutil"
	"real_estate_finder/real_estate_finder/internal/models"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery()) // recovers from any panics and writes a 500
	router.Use(gin.Logger())

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.IndentedJSON(200, gin.H{
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

	return router
}
