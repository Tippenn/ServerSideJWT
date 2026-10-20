package usecase

import (
	"context"

	"github.com/Tippenn/ServerSideJWT/domains/JWTVerification"
	"github.com/Tippenn/ServerSideJWT/domains/JWTVerification/entities"
	"github.com/Tippenn/ServerSideJWT/domains/JWTVerification/models/request"
)

type serverJWTUseCase struct {
	repo JWTVerification.ServerJWTRepository
}

func NewServerJWTUseCase(repo JWTVerification.ServerJWTRepository) JWTVerification.ServerJWTUseCase {
	return serverJWTUseCase{repo: repo}
}

// Login implements [JWTVerification.ServerJWTUseCase].
func (s serverJWTUseCase) Login(ctx context.Context, requests *request.LoginRequest) error {
	_, err := s.repo.CreateServerJWT(requests.AccessToken, requests.UserId)

	if err != nil {
		return err
	}

	return nil
}

// Logout implements [JWTVerification.ServerJWTUseCase].
func (s serverJWTUseCase) Logout(ctx context.Context, requests *request.LogoutRequest) error {
	jwtData, err := s.repo.FindServerJWTbyID(requests.AccessToken)
	if err != nil {
		return err
	}

	err = s.repo.UpdateServerJWT(jwtData)

	if err != nil {
		return err
	}

	return nil
}

// CheckJWTServer implements [JWTVerification.ServerJWTUseCase].
func (s serverJWTUseCase) CheckJWTServer(ctx context.Context, requests *request.CheckJWTRequest) (*entities.ServerJWT, error) {
	jwtData, err := s.repo.FindServerJWTbyID(requests.AccessToken)
	if err != nil {
		return nil, err
	}

	return jwtData, nil
}
