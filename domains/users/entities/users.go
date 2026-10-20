package entities

import "github.com/Tippenn/ServerSideJWT/domains/JWTVerification/entities"

type User struct {
	Id       string `gorm:"primaryKey;type:varchar(255)"`
	Username string
	Name     string
	Password string
	Enabled  int
	ServerJWTs []entities.ServerJWT `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE"`
}
