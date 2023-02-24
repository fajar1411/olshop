package faktory

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	userData "toko/fitur/user/data"
	userService "toko/fitur/user/service"
	userHandler "toko/routes"

	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	v := validator.New()
	userRepofaktory := userData.NewUser(db) //menginiasialisasi func new yang ada di repository
	userServiceFaktory := userService.NewService(userRepofaktory, v)
	userHandler.NewHandlerUser(userServiceFaktory, e)

}
