package faktory

import (
	"github.com/labstack/echo/v4"

	userData "toko/fitur/user/data"
	userService "toko/fitur/user/service"
	userHandler "toko/routes"

	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	userRepofaktory := userData.NewUser(db) //menginiasialisasi func new yang ada di repository
	userServiceFaktory := userService.NewService(userRepofaktory)
	userHandler.NewHandlerUser(userServiceFaktory, e)

}
