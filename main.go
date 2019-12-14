package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-media-site/routes"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("views/*.html")
	router.Static("/assets", "./assets")

	user := router.Group("/user")
	{
		user.POST("/signup", routes.UserSignUp)
		user.POST("/login", routes.UserLogIn)
	}

	router.GET("/", routes.Home)
	router.GET("/login", routes.LogIn)
	router.GET("/signup", routes.SignUp)
	router.NoRoute(routes.NoRoute)

	// router.GET("/", func(ctx *gin.Context) {
	// 	ctx.HTML(200, "index.html", gin.H{})
	// })

	router.Run()
}
