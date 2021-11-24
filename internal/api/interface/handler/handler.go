package handler

import (
	"github.com/labstack/echo"
	"github.com/mitsu9/clean-architecture-go/internal/api/infrastructure/dynamodb"
	"github.com/mitsu9/clean-architecture-go/internal/api/usecase"
)

type Handler struct {
	GetUserHandler
}

func CreateHandler() (Handler, error) {
	// repository
	// TODO: db instance を作って渡す
	// session, err := session.NewSession(config.NewDynamoConfig())
	//	if err != nil {
	//		return nil, err
	//	}
	// db := dynamo.New(session)
	userRepository := dynamodb.NewUserRepository(nil)

	// usecase
	getUserUsecase := usecase.NewGetUserUsecase(&userRepository)

	// handler
	return Handler{
		GetUserHandler: GetUserHandler{
			GetUserUsecase: getUserUsecase,
		},
	}, nil
}

func RegisterHandlers(router *echo.Group) error {
	handler, err := CreateHandler()
	if err != nil {
		return err
	}
	router.GET("/users/:user_id", handler.GetUserHandler.GetUser)
	return nil
}
