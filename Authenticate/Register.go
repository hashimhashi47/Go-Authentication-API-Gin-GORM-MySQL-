package authenticate

import (
	config "authentication/Config"
	model "authentication/Model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var Validate model.UserInfo

	if err := c.ShouldBindJSON(&Validate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"ERROR": "INVALID JSON",
			"ON":    err.Error(),
		})
		return
	}

	user := model.User{
		Username: Validate.Username,
		Email:    Validate.Email,
		Password: Validate.Password,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"ERROR": "INVALID CREDINTIAL",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"SUCESS": "SUCESSFULY REGISTERED",
		"USER":   Validate.Username,
	})
}
