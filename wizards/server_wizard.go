package wizards

import (
	"github.com/Tippenn/ServerSideJWT/shared/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterServer(router *gin.Engine) {
	auth := router.Group("/auth")
	{
		auth.POST("/login", UserHttp.Login)
		auth.POST("/signup", UserHttp.SignUp)
	}

	protected := router.Group("/auth")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/dashboard", UserHttp.Dashboard)
		protected.GET("/logout", ServerJWTHttp.Logout)
	}
}
