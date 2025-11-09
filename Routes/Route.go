package routes

import (
	authenticate "authentication/Authenticate"
	middileware "authentication/Middileware"

	"github.com/gin-gonic/gin"
)

// Routes
func Routes(e *gin.Engine) {
	e.POST("/Register", authenticate.Register)
	e.POST("/Login", authenticate.Login)
	e.POST("/Logout", middileware.Middilewares(), authenticate.Logout)
	e.GET("/Profile", middileware.Middilewares(), authenticate.Profile)

}
