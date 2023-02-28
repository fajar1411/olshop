package routes

import (
	pelanggan "toko/fitur/pelanggan"
	handlerpelanggan "toko/fitur/pelanggan/handler"
	"toko/middlewares"

	"github.com/labstack/echo/v4"
)

func NewHandlerPelanggan(Service pelanggan.PelangganService, e *echo.Echo) {
	handlers := &handlerpelanggan.PelangganHandler{
		PelangganServices: Service,
	}

	e.POST("/register", handlers.Register, middlewares.JWTMiddleware())
	e.POST("/login", handlers.Login)
	e.GET("/profile", handlers.Profile, middlewares.JWTMiddleware())
	e.PUT("/user", handlers.Update, middlewares.JWTMiddleware())

}
