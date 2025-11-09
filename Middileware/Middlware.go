package middileware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

// authenticate middileware
func Middilewares() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessions := sessions.Default(c)
		user := sessions.Get("mysession")

		if user == nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"ERROR": "USER NOT FOUND, LOGIN IN FIRST",
			})
			c.Abort()
			return
		}
		c.Next()
	}

}
