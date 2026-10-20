package http

import (
	"fmt"
	"net/http"

	"github.com/Tippenn/ServerSideJWT/domains/users"
	"github.com/Tippenn/ServerSideJWT/domains/users/models/request"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Userhttp struct {
	uc users.UserUseCase
}

func NewUserHttp(uc users.UserUseCase) *Userhttp {
	return &Userhttp{uc: uc}
}

func (handler *Userhttp) SignUp(c *gin.Context) {
	ctx := c.Request.Context()
	requestBody := &request.SignUpRequest{}

	err := c.ShouldBindJSON(&requestBody)

	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"error": "Invalid request",
		})
		return
	}

	validate := validator.New()
	err = validate.StructCtx(ctx, requestBody)

	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	accessToken, err := handler.uc.SignUp(ctx, requestBody)

	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"access_token": accessToken,
	})
}

func (handler *Userhttp) Login(c *gin.Context) {
	ctx := c.Request.Context()
	requestBody := &request.LoginRequest{}
	err := c.ShouldBindJSON(&requestBody)

	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"error": "Invalid request",
		})
		return
	}

	validate := validator.New()
	err = validate.StructCtx(ctx, requestBody)

	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	accessToken, err := handler.uc.Login(ctx, requestBody)

	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"access_token": accessToken,
	})
}

func (handler *Userhttp) Dashboard(c *gin.Context) {
	ctx := c.Request.Context()
	Id := c.GetString("user_Id")

	value, exists := c.Get("user_Id")
	fmt.Printf("exists=%v value=%v type=%T\n", exists, value, value)

	user, err := handler.uc.Dashboard(ctx, Id)
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":       user.Username,
		"name":     user.Name,
		"username": user.Username,
	})
}


