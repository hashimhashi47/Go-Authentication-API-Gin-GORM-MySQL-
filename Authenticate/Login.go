package authenticate

import (
	config "authentication/Config"
	model "authentication/Model"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)



var login model.User

func Login(c *gin.Context) {
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"ERROR": "INVALID JSON",
			"ON":    err.Error(),
		})
		return
	}

	if err := config.DB.Where("email = ?", user.Email).First(&login).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"ERROR": "USER NOT FOUND",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(login.Password), []byte(user.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"ERROR": "INVALID PASSWORD",
		})
		return
	}

	session := sessions.Default(c)
	session.Set("mysession", login.Username)
	session.Save()

	c.JSON(http.StatusOK, gin.H{
		"SUCESS": "WELCOME BACK",
		"USER":   login.Username,
	})
}
