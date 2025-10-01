package routes

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Lokeshxs/url-shortener/middleware"
	"github.com/Lokeshxs/url-shortener/model"
	"github.com/Lokeshxs/url-shortener/utils"
	"github.com/Lokeshxs/url-shortener/webhook"
	"github.com/gin-gonic/gin"
)

func RoutingHandler(server *gin.Engine) {

	// URL Shorten Endpoint //:PRIVATE
	server.POST("/shorten", middleware.ClerkMiddleware(), middleware.RateLimiter, func(c *gin.Context) {

		var url model.URL

		// Get the long url from body

		err := c.ShouldBindJSON(&url)

		if err != nil || len(url.LongURL) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Long URL is missing",
			})
		}

		// Calling a function to shorten the URL

		var userId = c.GetString("userId")

		err = url.ShortenURL(userId, url.ExpiredAT)

		if err != nil {

			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to generate short url. Try again!",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message":   "Short url created successfully",
			"short_url": url.ShortURL,
		})

	})

	// URL Redirect Endpoint  //:PUBLIC
	server.GET("/u/:code", func(c *gin.Context) {

		// Extracting the short code
		shortCode := c.Param("code")

		// checking if the short code is empty or not of required length
		if shortCode == "" || len(shortCode) != utils.SHORT_CODE_LENGTH {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid  url",
			})

			return
		}

		// Get the long url from DB for the short code and incase the url is not found send 404
		long_url, err := model.GetURL(shortCode)

		if err != nil {
			errMsg := err
			statusCode := http.StatusInternalServerError
			if err == sql.ErrNoRows {
				errMsg = errors.New("Invalid URL")
				statusCode = http.StatusNotFound
			}

			c.JSON(statusCode, gin.H{
				"message": errMsg.Error(),
			})

			return

		}

		// Send the long url redirect response

		c.Redirect(http.StatusTemporaryRedirect, long_url)

	})

	// URL STATS Endpoint //:PRIVATE

	server.GET("/stats", middleware.ClerkMiddleware(), func(c *gin.Context) {

		var userId = c.GetString("userId")
		page := c.DefaultQuery("page", "1")
		limit := c.DefaultQuery("limit", "10")

		pageNumber, err := strconv.ParseInt(page, 10, 64)
		pageLimit, err := strconv.ParseInt(limit, 10, 64)

		
		if err != nil || pageLimit < 5 {
			pageLimit = 10
		}
		if err != nil || pageNumber < 1 {
			pageNumber = 1
		}

		resultedUrls, totalRecords, err := model.GetAllURLS(userId, pageNumber, pageLimit)

		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to fetch the stats",
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"stats":      resultedUrls,
			"total":      totalRecords,
			"page":       pageNumber,
			"limit":      pageLimit,
			"totalPages": (totalRecords + pageLimit - 1) / pageLimit,
		})

	})

	// WEBHOOKS Handlers //:PUBLIC
	server.POST("/webhook/signup", webhook.ClerkSignUp)
}
