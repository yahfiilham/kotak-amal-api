package server

import (
	"database/sql"
	"net/http"

	"kotak-amal/handler"
	"kotak-amal/repository"
	"kotak-amal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ApiServer struct {
	DB        *sql.DB
	Router    *gin.Engine
	Validator *validator.Validate
}

func NewServer(db *sql.DB, validator *validator.Validate) *ApiServer {
	r := gin.New()
	return &ApiServer{
		DB:        db,
		Router:    r,
		Validator: validator,
	}
}

func (server *ApiServer) ListenAndServe(port string) {
	server.Router.Use(gin.Logger())
	server.registerRouter()

	http.ListenAndServe(":"+port, server.Router)
}

func (server *ApiServer) registerRouter() {
	userRepository := repository.NewUserRepository(server.DB)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := handler.NewUserHandler(userUsecase, server.Validator)

	r := server.Router.Group("/api/v1")

	// User
	r.POST("/users", userHandler.CreateUser)
}
