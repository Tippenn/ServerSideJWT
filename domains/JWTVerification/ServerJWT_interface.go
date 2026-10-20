package JWTVerification

import (
	"context"

	"github.com/Tippenn/ServerSideJWT/domains/JWTVerification/entities"
	"github.com/Tippenn/ServerSideJWT/domains/JWTVerification/models/request"
)

type ServerJWTUseCase interface {
	Login(ctx context.Context, requests *request.LoginRequest) error
	CheckJWTServer(ctx context.Context, requests *request.CheckJWTRequest) (*entities.ServerJWT, error)
	Logout(ctx context.Context, requests *request.LogoutRequest) error
}

type ServerJWTRepository interface {
	FindServerJWTbyID(Id string) (*entities.ServerJWT, error)
	CreateServerJWT(accessToken string, userId string) (*entities.ServerJWT, error)
	UpdateServerJWT(ServerJWT *entities.ServerJWT) error
	DeleteServerJWT(accessToken string) error
}
