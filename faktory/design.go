package faktory

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	"toko/config"
	pelangganData "toko/fitur/pelanggan/data"
	pelanganService "toko/fitur/pelanggan/service"
	"toko/helper"
	pelangganHandler "toko/routes"

	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	v := validator.New()
	cfg := config.GetConfig()
	cld := helper.NewCloud(cfg)
	userRepofaktory := pelangganData.NewPelanggan(db) //menginiasialisasi func new yang ada di repository
	userServiceFaktory := pelanganService.NewService(userRepofaktory, v, cld)
	pelangganHandler.NewHandlerPelanggan(userServiceFaktory, e)

}
