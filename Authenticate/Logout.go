package authenticate

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)


func Logout(c *gin.Context){
	session:=sessions.Default(c)
	session.Clear()
	session.Save()
	
	c.JSON(http.StatusOK, gin.H{
		"SUCESS":"SUCESSFULLY LOGOUT",
	})
}