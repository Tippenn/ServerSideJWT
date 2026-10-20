package repository

import (
	"context"

	"github.com/Tippenn/ServerSideJWT/domains/users"
	"github.com/Tippenn/ServerSideJWT/domains/users/entities"
	"github.com/Tippenn/ServerSideJWT/infrastructure"
)

type userRepository struct {
	db infrastructure.Database
}

func NewUserRepository(db infrastructure.Database) users.UserRepository {
	return userRepository{db: db}
}

// CreateUser implements [users.UserRepository].
func (u userRepository) CreateUser(ctx context.Context, userId string, username string, name string, password string) error {
	user := &entities.User{
		Id:       userId,
		Username: username,
		Name:     name,
		Password: password,
		Enabled:  1,
	}

	result := u.db.GetInstance().Create(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// FindUserbyID implements [users.UserRepository].
func (u userRepository) FindUserbyID(Id string) (*entities.User, error) {
	var user *entities.User
	result := u.db.GetInstance().First(&user, "id=?", Id)

	if result.Error != nil {
		println("repo")
		return nil, result.Error
	}

	return user, nil
}

// FindUserbyUsername implements [users.UserRepository].
func (u userRepository) FindUserbyUsername(username string) (*entities.User, error) {
	var user *entities.User
	result := u.db.GetInstance().First(&user, "username=?", username)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

// UpdateUser implements [users.UserRepository].
func (u userRepository) UpdateUser(username *entities.User) error {
	panic("unimplemented")
}
