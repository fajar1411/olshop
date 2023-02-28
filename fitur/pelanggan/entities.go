package pelanggan

import "mime/multipart"

type PelangganEntites struct {
	ID        uint
	Nama      string `validate:"required,min=5,required"`
	Password  string `validate:"required,min=5,required"`
	Email     string `validate:"required,email"`
	Image_url string
	Role      string
}

type PelangganService interface {
	Login(email, password string) (string, PelangganEntites, error)
	Register(newUser PelangganEntites) (PelangganEntites, error)
	Profile(id int) (PelangganEntites, error)
	UpdateUser(id int, fileData *multipart.FileHeader, Updata PelangganEntites) (PelangganEntites, error)
	DeleteUser(id int) (PelangganEntites, error)
}

type PelangganData interface {
	Login(email string) (PelangganEntites, error)
	Register(newUser PelangganEntites) (PelangganEntites, error)
	Profile(id int) (PelangganEntites, error)
	DeleteUser(id int) (PelangganEntites, error)
	UpdateUser(id int, Updata PelangganEntites) (PelangganEntites, error)
}
