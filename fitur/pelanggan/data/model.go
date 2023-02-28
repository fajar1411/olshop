package data

import (
	"toko/fitur/pelanggan"

	"gorm.io/gorm"
)

type Pelanggan struct {
	gorm.Model
	Nama      string
	Password  string
	Email     string `gorm:"unique"`
	Image_url string
	Role      string
}

func FromEntities(dataCore pelanggan.PelangganEntites) Pelanggan { //fungsi yang mengambil data dari entities usercore dan merubah data ke user gorm(model.go)
	return Pelanggan{

		Email:     dataCore.Email,
		Password:  dataCore.Password,
		Nama:      dataCore.Nama,
		Role:      dataCore.Role,
		Image_url: dataCore.Image_url,
	}

}
func ToCore(data Pelanggan) pelanggan.PelangganEntites {
	return pelanggan.PelangganEntites{
		ID:        data.ID,
		Nama:      data.Nama,
		Email:     data.Email,
		Password:  data.Password,
		Role:      data.Role,
		Image_url: data.Image_url,
	}
}
func (dataModel *Pelanggan) ModelsToCore() pelanggan.PelangganEntites { //fungsi yang mengambil data dari  user gorm(model.go)  dan merubah data ke entities usercore
	return pelanggan.PelangganEntites{
		Nama:     dataModel.Nama,
		Email:    dataModel.Email, //mapping data core ke data gorm model
		Password: dataModel.Password,
		Role:     dataModel.Role,
	}
}
