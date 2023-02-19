package handler

import "toko/fitur/user"

type UserReponse struct {
	// ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func ToResponse(data user.UserEntites) UserReponse {
	return UserReponse{

		Name:     data.Nama,
		Email:    data.Email,
		Password: data.Password,
	}
}
