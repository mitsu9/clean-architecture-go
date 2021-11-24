package api

import (
	"github.com/labstack/echo"
	"github.com/mitsu9/clean-architecture-go/internal/api/interface/handler"
)

func StartServer() {
	e := echo.New()
	api := e.Group("/api")
	if err := handler.RegisterHandlers(api); err != nil {
		// TODO: output log
	}
}
