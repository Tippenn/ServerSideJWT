package wizards

import (
	"github.com/Tippenn/ServerSideJWT/config"
	serverJWTHttp "github.com/Tippenn/ServerSideJWT/domains/JWTVerification/handlers/http"
	serverJWTRepository "github.com/Tippenn/ServerSideJWT/domains/JWTVerification/repository"
	serverJWTUsecase "github.com/Tippenn/ServerSideJWT/domains/JWTVerification/usecase"
	"github.com/Tippenn/ServerSideJWT/domains/users/handlers/http"
	"github.com/Tippenn/ServerSideJWT/domains/users/repository"
	"github.com/Tippenn/ServerSideJWT/domains/users/usecase"
	"github.com/Tippenn/ServerSideJWT/infrastructure"
)

var (
	Config = config.GetConfig()

	//Storage
	PostgresDatabase = infrastructure.NewPostgresDatabase(Config)

	//repo
	UserRepo      = repository.NewUserRepository(PostgresDatabase)
	ServerJWTRepo = serverJWTRepository.NewServerJWTRepository(PostgresDatabase)

	//usecase
	ServerJWTUc = serverJWTUsecase.NewServerJWTUseCase(ServerJWTRepo)
	UserUc      = usecase.NewUserUseCase(UserRepo, ServerJWTUc)

	//http
	ServerJWTHttp = serverJWTHttp.NewServerJWTHttp(ServerJWTUc)
	UserHttp      = http.NewUserHttp(UserUc)
)
