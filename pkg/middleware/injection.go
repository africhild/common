package middleware

import (
	"github.com/africhild/common/pkg/injection"
	"github.com/africhild/common/pkg/util"
	"github.com/gin-gonic/gin"
)

func Injection() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.GetHeader("x-user")
		if user != "" {
			injection.SetUser(c, util.FromBase64(user))
		}

		workspace := c.GetHeader("x-workspace")
		if workspace != "" {
			injection.SetWorkspace(c, util.FromBase64(workspace))
		}

		c.Next()
	}
}
