package repository

import (
	"errors"
	"time"

	"github.com/Tippenn/ServerSideJWT/domains/JWTVerification"
	"github.com/Tippenn/ServerSideJWT/domains/JWTVerification/entities"
	"github.com/Tippenn/ServerSideJWT/infrastructure"
)

type serverJWTRepository struct {
	db infrastructure.Database
}

func NewServerJWTRepository(db infrastructure.Database) JWTVerification.ServerJWTRepository {
	return serverJWTRepository{db: db}
}

// CreateServerJWT implements [JWTVerification.ServerJWTRepository].
func (s serverJWTRepository) CreateServerJWT(accessToken string, userId string) (*entities.ServerJWT, error) {
	user := &entities.ServerJWT{
		Id:        accessToken,
		UserId:    userId,
		Revoked:   false,
		ExpiresAt: time.Now().Add(1 * time.Hour),
	}

	result := s.db.GetInstance().Create(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

// DeleteServerJWT implements [JWTVerification.ServerJWTRepository].
func (s serverJWTRepository) DeleteServerJWT(accessToken string) error {
	panic("unimplemented")
}

// FindServerJWTbyID implements [JWTVerification.ServerJWTRepository].
func (s serverJWTRepository) FindServerJWTbyID(Id string) (*entities.ServerJWT, error) {
	var user *entities.ServerJWT
	result := s.db.GetInstance().First(&user, "id=?", Id)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

// UpdateServerJWT implements [JWTVerification.ServerJWTRepository].
func (s serverJWTRepository) UpdateServerJWT(ServerJWT *entities.ServerJWT) error {
	var user *entities.ServerJWT
	result := s.db.GetInstance().Model(&user).
		Where("id = ?", ServerJWT.Id).
		Update("revoked", ServerJWT.Revoked)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("token not found")
	}
	return nil
}
