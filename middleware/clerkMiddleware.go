package middleware
	
import (
	"net/http"

	"github.com/clerk/clerk-sdk-go/v2"
	clerkhttp "github.com/clerk/clerk-sdk-go/v2/http"
	
	"github.com/gin-gonic/gin"
)




func ClerkMiddleware() gin.HandlerFunc {


	return func(c *gin.Context) {
		// Wrap Clerk’s middleware
		handler := clerkhttp.WithHeaderAuthorization()(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Inject Clerk’s context into Gin’s request
			c.Request = c.Request.WithContext(r.Context())
		}))

		// Run Clerk auth check
		handler.ServeHTTP(c.Writer, c.Request)

	

		// Extract user claims
		claims, ok := clerk.SessionClaimsFromContext(c.Request.Context())
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Clerk authorization failed!",
			})
				c.Abort();
				return ;
		}

		// Save userId for later use in handlers
		c.Set("userId", claims.Subject)

		// Continue
		c.Next()
	}
}
