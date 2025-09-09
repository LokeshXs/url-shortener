package routes

import (
	"fmt"
	"net/http"
	"github.com/Lokeshxs/url-shortener/model"
	"github.com/Lokeshxs/url-shortener/utils"
	"github.com/gin-gonic/gin"
)

func RoutingHandler(server *gin.Engine) {

	// URL Shorten Endpoint
	server.POST("/shorten", func(c *gin.Context) {

		var url model.URL

		// Get the long url from body

		err := c.ShouldBindJSON(&url)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Long URL is missing",
			})
		}

		// Call a function to shorten the URL

		domain := c.Request.Host

		fmt.Println(domain)

		url.ShortenURL(domain)

		c.JSON(http.StatusOK, gin.H{
			"message":   "Short url created successfully",
			"short_url": url.ShortURL,
		})

	})


	// URL Redirect Endpoint
	server.GET("/u/:code",func(c *gin.Context){

		// Extracting the short code
		shortCode := c.Param("code");

		// checking if the short code is empty or not of required length
		if(shortCode == "" || len(shortCode) != utils.SHORT_CODE_LENGTH){
			c.JSON(http.StatusBadRequest,gin.H{
				"message":"Invalid short url",
			})

			return;
		}

		// Get the long url from DB for the short code and incase the url is not found send 404 
		long_url := "https://www.youtube.com/watch?v=2HZmWnMWU-A";


		// Send the long url redirect response

		c.Redirect(http.StatusMovedPermanently,long_url)

	})


}
