package middileware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)


//cookie and session
func Cookie() gin.HandlerFunc{
	store:=cookie.NewStore([]byte("THE-SECRET-KEY"))
	return sessions.Sessions("mysession", store)
}