package http

import (
	"github.com/Tippenn/ServerSideJWT/domains/JWTVerification"
	"github.com/gin-gonic/gin"
)

type ServerJWTHttp struct {
	uc JWTVerification.ServerJWTUseCase
}

func NewServerJWTHttp(uc JWTVerification.ServerJWTUseCase) *ServerJWTHttp {
	return &ServerJWTHttp{uc: uc}
}

func (handler *ServerJWTHttp) Logout(c *gin.Context) {

}
