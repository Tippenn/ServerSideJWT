package users

import (
	"context"

	"github.com/Tippenn/ServerSideJWT/domains/users/entities"
	"github.com/Tippenn/ServerSideJWT/domains/users/models/request"
	"github.com/Tippenn/ServerSideJWT/domains/users/models/response"
)

type UserUseCase interface {
	Login(ctx context.Context, requests *request.LoginRequest) (*response.LoginResponse, error)
	SignUp(ctx context.Context, requests *request.SignUpRequest) (*response.SignUpResponse, error)
	Dashboard(ctx context.Context, Id string) (*entities.User, error)
	Logout(ctx context.Context, requests *request.LogoutRequest) (*response.LogoutResponse, error)
}

type UserRepository interface {
	FindUserbyUsername(username string) (*entities.User, error)
	FindUserbyID(Id string) (*entities.User, error)
	CreateUser(ctx context.Context, userId string, username string, name string, password string) error
	UpdateUser(username *entities.User) error
}
