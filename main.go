package main

import (
	config "authentication/Config"
	middileware "authentication/Middileware"
	routes "authentication/Routes"

	"github.com/gin-gonic/gin"
)

func main(){
	g:=gin.Default()
	config.DBconfig()
	g.Use(middileware.Cookie())
	routes.Routes(g)
	g.Run(":8000")
}

