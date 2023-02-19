package routes

import (
	"toko/fitur/user"
	handleruser "toko/fitur/user/handler"
	"toko/middlewares"

	"github.com/labstack/echo/v4"
)

func NewHandlerUser(Service user.UserService, e *echo.Echo) {
	handlers := &handleruser.UserHandler{
		UserServices: Service,
	}

	e.POST("/user", handlers.Create, middlewares.JWTMiddleware())
	e.POST("/user", handlers.Login, middlewares.JWTMiddleware())

}
