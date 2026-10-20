package http

import (
	"log"
	"net/http"

	"github.com/Tippenn/ServerSideJWT/domains/JWTVerification"
	"github.com/Tippenn/ServerSideJWT/domains/JWTVerification/models/request"
	"github.com/gin-gonic/gin"
)

type ServerJWTHttp struct {
	uc JWTVerification.ServerJWTUseCase
}

func NewServerJWTHttp(uc JWTVerification.ServerJWTUseCase) *ServerJWTHttp {
	return &ServerJWTHttp{uc: uc}
}

func (handler *ServerJWTHttp) Logout(c *gin.Context) {
	ctx := c.Request.Context()
	Id := c.GetString("user_Id")
	err := handler.uc.Logout(ctx, &request.LogoutRequest{UserId: Id})

	if err != nil {
		log.Println("logout failed:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to logout"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "logout successful"})
}
