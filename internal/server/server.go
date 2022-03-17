package server

import (
	"encoding/json"
	"io/ioutil"
	"real_estate_finder/real_estate_finder/internal/dto"

	"github.com/gin-gonic/gin"
)

type repositoryInterface interface{}

type geocoderInterface interface {
	FindAll() ([]*dto.FindAllResponseBody, error)
	AddAddress(addr string) (*dto.AddAddressResponseBody, error)
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
	r := gin.Default() // Default With the Logger and Recovery middleware already attached

	return &Server{
		r:          r,
		repository: repository,
		geocoder:   geocoder,
	}
}

func (s *Server) Start(addr string) error {
	s.r.GET("addresses", func(c *gin.Context) {
		resp, err := s.geocoder.FindAll()
		if err != nil {
			c.IndentedJSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.IndentedJSON(200, &resp)
	})

	s.r.POST("addresses", func(c *gin.Context) {
		b, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.IndentedJSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		var req *dto.AddAddressRequestBody
		err = json.Unmarshal(b, &req)
		if err != nil {
			c.IndentedJSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		if req.FullAddress == "" {
			c.IndentedJSON(400, gin.H{
				"error": "address len is 0",
			})
			return
		}
		resp, err := s.geocoder.AddAddress(req.FullAddress)
		if err != nil {
			c.IndentedJSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.IndentedJSON(200, &resp)
	})
	//s.r.POST("addresses", func(c *gin.Context) {
	//	reqLat, ok := c.GetPostForm("lat")
	//	if !ok {
	//		c.IndentedJSON(http.StatusBadRequest, gin.H{
	//			"error": "no latitude argument provided",
	//		})
	//	}
	//	lat, err := strconv.ParseFloat(reqLat, 64)
	//	if err != nil {
	//		c.IndentedJSON(http.StatusBadRequest, gin.H{
	//			"error": err.Error(),
	//		})
	//	}
	//
	//	reqLng, ok := c.GetPostForm("lng")
	//	if !ok {
	//		c.IndentedJSON(http.StatusBadRequest, gin.H{
	//			"error": "no longitude argument provided",
	//		})
	//	}
	//	lng, err := strconv.ParseFloat(reqLng, 64)
	//	if err != nil {
	//		c.IndentedJSON(http.StatusBadRequest, gin.H{
	//			"error": err.Error(),
	//		})
	//	}
	//
	//	resp, err := s.geocoder.AddAddress(lat, lng)
	//	if err != nil {
	//		c.IndentedJSON(500, gin.H{"geocoder error": err.Error()})
	//		return
	//	}
	//
	//	c.IndentedJSON(200, &resp)
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
