package middleware

import (

	"net/http"
	"sync"
	"time"

	"github.com/Lokeshxs/url-shortener/utils"

	"github.com/gin-gonic/gin"
)

// User Requests Info Struct
type userRequestsInfo struct {
	requests int
	resetAt  time.Time
}

// Initializing the mutex for making sure the no 2 or more goroutines can access the code at the same time

var mu sync.Mutex

// Global Map to store requests data
var requestsMap = make(map[string]*userRequestsInfo)

func RateLimiter(c *gin.Context) {

	// Locking the mutex
	mu.Lock()

	// defering the unlock whenever the func execution finishes
	defer mu.Unlock()



	
// User Id from clerk
	var userId = c.GetString("userId");




	var currentTime = time.Now()

	// If user does not exist or window expired -> reset

	data, exists := requestsMap[userId]

	if !exists || currentTime.After(data.resetAt) {
		requestsMap[userId] = &userRequestsInfo{
			requests: 1,
			resetAt:  currentTime.Add(utils.RATE_LIMIT_TIME_WINDOW),
		}
	} else {
		// Within the Window
		if data.requests >= utils.RATE_LIMIT_REQUESTS_WINDOW {

			// Rate limit exceeded

			c.JSON(http.StatusTooManyRequests, gin.H{
				"message": "Rate limit exceeded, Will reset in 1 min",
			})

			c.Abort()
			return
		}

		// Updating the requests
		data.requests++
	}

	

	c.Next()

}
