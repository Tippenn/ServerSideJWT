package usecase

import (
	"context"
	"errors"

	serverJWT "github.com/Tippenn/ServerSideJWT/domains/JWTVerification"
	serverJWTRequest "github.com/Tippenn/ServerSideJWT/domains/JWTVerification/models/request"
	"github.com/Tippenn/ServerSideJWT/domains/users"
	"github.com/Tippenn/ServerSideJWT/domains/users/entities"
	"github.com/Tippenn/ServerSideJWT/domains/users/models/request"
	"github.com/Tippenn/ServerSideJWT/domains/users/models/response"
	"github.com/Tippenn/ServerSideJWT/shared/utils"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	repo  users.UserRepository
	jwtUC serverJWT.ServerJWTUseCase
}

func NewUserUseCase(repo users.UserRepository, jwtUC serverJWT.ServerJWTUseCase) users.UserUseCase {
	return userUseCase{repo: repo, jwtUC: jwtUC}
}

// Dashboard implements [users.UserUseCase].
func (u userUseCase) Dashboard(ctx context.Context, Id string) (*entities.User, error) {
	user, err := u.repo.FindUserbyID(Id)
	if err != nil {
		return nil, err
	}

	return user, err
}

// Login implements [users.UserUseCase].
func (u userUseCase) Login(ctx context.Context, requests *request.LoginRequest) (*response.LoginResponse, error) {
	user, err := u.repo.FindUserbyUsername(requests.Username)

	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requests.Password))
	if err != nil {
		return nil, errors.New("Unauthorized")
	}

	if user.Enabled == 0 {
		return nil, errors.New("user is disabled")
	}

	token, err := utils.GenerateToken(user.Id)
	if err != nil {
		return nil, errors.New("token is not generated correctly")
	}

	loginRequest := serverJWTRequest.LoginRequest{
		UserId: user.Id,
		AccessToken: token,
	}
	err = u.jwtUC.Login(ctx, &loginRequest)

	if err != nil {
		return nil, errors.New("token is not verified on the database")
	}

	return &response.LoginResponse{AccessToken: token}, nil
}

// SignUp implements [users.UserUseCase].
func (u userUseCase) SignUp(ctx context.Context, requests *request.SignUpRequest) (*response.SignUpResponse, error) {
	if requests.Password != requests.ConfirmPassword {
		return nil, errors.New("Password don't match")
	}

	b, err := bcrypt.GenerateFromPassword([]byte(requests.Password), 10)
	if err != nil {
		return nil, err
	}

	u.repo.CreateUser(ctx,
		uuid.New().String(),
		requests.Username,
		requests.Name,
		string(b),
	)

	return &response.SignUpResponse{
		Information: "HALO BERHASIL SIGN UP HEHEHEHE",
	}, nil
}

// Logout implements [users.UserUseCase].
func (u userUseCase) Logout(ctx context.Context, request *request.LogoutRequest) (*response.LogoutResponse, error) {

	//ini checkout in tokennya dari server jwt
	u.jwtUC.Logout(ctx, &serverJWTRequest.LogoutRequest{AccessToken: request.AccessToken})
	//return info aja
	return &response.LogoutResponse{
		Information: "HALO BERHASIL SIGN UP HEHEHEHE",
	}, nil
}
