package handler

import (
	"github.com/labstack/echo"
	"github.com/mitsu9/clean-architecture-go/internal/api/usecase"
)

type GetUserHandler struct {
	GetUserUsecase *usecase.GetUserUsecase
}

func NewGetUserHandler(getUserUsecase *usecase.GetUserUsecase) GetUserHandler {
	return GetUserHandler{
		GetUserUsecase: getUserUsecase,
	}
}

func (h *GetUserHandler) GetUser(ctx echo.Context) error {
	// TODO: impl
	return nil
}
