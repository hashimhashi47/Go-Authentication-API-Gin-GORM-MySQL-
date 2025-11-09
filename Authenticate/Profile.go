package authenticate

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Profile(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"SUCESS":   "YOU ARE ON PROFILE PAGE",
		"HEY USER": login.Username,
	})
}
