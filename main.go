package main

import (
	"fmt"
	"log"

	JWTVerification "github.com/Tippenn/ServerSideJWT/domains/JWTVerification/entities"
	users "github.com/Tippenn/ServerSideJWT/domains/users/entities"
	"github.com/Tippenn/ServerSideJWT/wizards"

	"github.com/gin-gonic/gin"
)

func main() {
	err := wizards.PostgresDatabase.GetInstance().AutoMigrate(
		&users.User{},
		&JWTVerification.ServerJWT{},
	)

	if err != nil {
		log.Fatal("migration failed: ", err)
	}
	router := gin.Default()

	wizards.RegisterServer(router)
	router.Run(fmt.Sprintf(":%d", wizards.Config.Server.Port))
}
