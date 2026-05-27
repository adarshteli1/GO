package middleware

import (
	"net/http"
	"notes/casbin"

	"github.com/gin-gonic/gin"
)

func CasbinMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, _ := c.Get("role")
		path := c.FullPath()
		method := c.Request.Method

		ok, _ := casbin.Enforcer.Enforce(role, path, method)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}
