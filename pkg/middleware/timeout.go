package middleware

import (
	"net/http"
	"time"

	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
)

func testResponse(c *gin.Context) {
	c.String(http.StatusRequestTimeout, "timeout")
}

func Timeout() gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(time.Second*10),
		timeout.WithHandler(func(c *gin.Context) {
			c.Next()
		}),
		timeout.WithResponse(testResponse),
	)
}
