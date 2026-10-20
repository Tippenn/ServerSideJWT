package infrastructure

import "gorm.io/gorm"

type Database interface {
	GetInstance() *gorm.DB
}
