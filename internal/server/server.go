package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"real_estate_finder/real_estate_finder/internal/dto"

	"github.com/gin-gonic/gin"
)

type repositoryInterface interface{}

type geocoderInterface interface {
	AddAddress(req *dto.AddAddressRequestBody) (*dto.AddAddressResponseBody, error)
	GetAddressByID(req *dto.GetAddressByIDRequestBody) (*dto.GetAddressByIDResponseBody, error)
}

type Server struct {
	r          *gin.Engine
	repository repositoryInterface
	geocoder   geocoderInterface
}

func New(
	repository repositoryInterface,
	geocoder geocoderInterface,
) *Server {
	r := gin.New()
	r.Use(gin.Recovery()) // recovers from any panics and writes a 500
	r.Use(gin.Logger())

	return &Server{
		r:          r,
		repository: repository,
		geocoder:   geocoder,
	}
}

func (s *Server) Start(addr string) error {
	// TODO: pretty error formatting
	s.r.GET("/ping", func(c *gin.Context) {
		c.IndentedJSON(200, gin.H{
			"message": "pong",
		})
	})

	s.r.POST("address", func(c *gin.Context) {
		b, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.IndentedJSON(500, gin.H{"readall error": err.Error()})
			return
		}

		var req *dto.AddAddressRequestBody
		err = json.Unmarshal(b, &req)
		if err != nil {
			c.IndentedJSON(500, gin.H{"unmarshalling error": err.Error()})
			return
		}
		fmt.Println(req)

		resp, err := s.geocoder.AddAddress(req)
		if err != nil {
			c.IndentedJSON(500, gin.H{"geocoder error": err.Error()})
			return
		}

		c.IndentedJSON(200, &resp)
	})

	s.r.GET("address", func(c *gin.Context) {
		b, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.IndentedJSON(500, gin.H{"readall error": err.Error()})
			return
		}

		var req *dto.GetAddressByIDRequestBody
		err = json.Unmarshal(b, &req)
		if err != nil {
			c.IndentedJSON(500, gin.H{"unmarshalling error": err.Error()})
			return
		}
		fmt.Println(req)
		resp, err := s.geocoder.GetAddressByID(req)
		if err != nil {
			c.IndentedJSON(500, gin.H{"geocoder error": err.Error()})
			return
		}

		c.IndentedJSON(200, &resp)
	})


	//s.r.POST("/midpoint", func(c *gin.Context) {
	//	b, err := ioutil.ReadAll(c.Request.Body)
	//	if err != nil {
	//		c.IndentedJSON(500, gin.H{"error reading body": err})
	//	}
	//
	//	var p []models.Point
	//	if err = json.Unmarshal(b, &p); err != nil {
	//		c.IndentedJSON(500, gin.H{"unmarshalling error": err})
	//	}
	//
	//	mid, err := models.Midpoint(p)
	//	if err != nil {
	//		c.IndentedJSON(500, gin.H{
	//			"error": err.Error(),
	//		})
	//	} else {
	//		c.IndentedJSON(200, mid)
	//	}
	//})
	//
	//s.r.GET("/map", func(c *gin.Context) {
	//	c.HTML(
	//		http.StatusOK,
	//		"mapbasics.html",
	//		gin.H{
	//			"token": token,
	//		},
	//	)
	//})

	return s.r.Run(addr)
}

func (s *Server) LoadHTML(pattern string) {
	s.r.LoadHTMLGlob(pattern)
}
