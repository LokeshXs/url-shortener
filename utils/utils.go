package utils

import (

	"math/rand"

	"github.com/gin-gonic/gin"
)

func GenerateShortCode(length int16) string {

	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	bytesSlice := make([]byte, length)

	for i := range bytesSlice {

		bytesSlice[i] = charset[rand.Intn(len(charset))]
	}

	return string(bytesSlice)

}

func GetBaseURL(c *gin.Context) string {
	// scheme := "http"
	// host := c.Request.Header.Get("X-Forwarded-Host")
	// if host == "" {

	// 	if c.Request.TLS != nil {
	// 		scheme = "https"
	// 	}
	// 	host = c.Request.Host
	// }

	// baseURL := fmt.Sprintf("%s://%s", scheme, host)

	return SHORTEN_BASE_URL
}
