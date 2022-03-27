package server

import (
	"real_estate_finder/real_estate_finder/internal/dto"

	"github.com/gin-gonic/gin"
)

type repositoryInterface interface{}

type webscraperInterface interface {
	ParseRentalOffers() error
	GetRentalOfferse() ([]*dto.RentalOffer, error)
}

type Server struct {
	r          *gin.Engine
	repository repositoryInterface
	webscraper webscraperInterface
}

func New(
	repository repositoryInterface,
	webscraper webscraperInterface,
) *Server {
	r := gin.Default() // Default With the Logger and Recovery middleware already attached

	return &Server{
		r:          r,
		repository: repository,
		webscraper: webscraper,
	}
}

func (s *Server) Start(addr string) error {
	s.r.GET("rental-offers/parse", func(c *gin.Context) {
		err := s.webscraper.ParseRentalOffers()
		if err != nil {
			c.IndentedJSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.IndentedJSON(200, gin.H{
			"message": "success",
		})
	})

	s.r.GET("rental-offers/", func(c *gin.Context) {
		resp, err := s.webscraper.GetRentalOfferse()
		if err != nil {
			c.IndentedJSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.IndentedJSON(200, resp)
	})

	//s.r.GET("addresses", func(c *gin.Context) {
	//	resp, err := s.geocoder.FindAll()
	//	if err != nil {
	//		c.IndentedJSON(500, gin.H{
	//			"error": err.Error(),
	//		})
	//		return
	//	}
	//
	//	c.IndentedJSON(200, &resp)
	//})

	//s.r.POST("addresses", func(c *gin.Context) {
	//	b, err := ioutil.ReadAll(c.Request.Body)
	//	if err != nil {
	//		c.IndentedJSON(500, gin.H{
	//			"error": err.Error(),
	//		})
	//		return
	//	}
	//	var req *dto.AddAddressRequestBody
	//	err = json.Unmarshal(b, &req)
	//	if err != nil {
	//		c.IndentedJSON(500, gin.H{
	//			"error": err.Error(),
	//		})
	//		return
	//	}
	//
	//	if req.FullAddress == "" {
	//		c.IndentedJSON(400, gin.H{
	//			"error": "address len is 0",
	//		})
	//		return
	//	}
	//	resp, err := s.geocoder.AddAddress(req.FullAddress)
	//	if err != nil {
	//		c.IndentedJSON(500, gin.H{
	//			"error": err.Error(),
	//		})
	//		return
	//	}
	//
	//	c.IndentedJSON(200, &resp)
	//})

	return s.r.Run(addr)
}

func (s *Server) LoadHTML(pattern string) {
	s.r.LoadHTMLGlob(pattern)
}
