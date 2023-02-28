package handler

import "toko/fitur/pelanggan"

type PelangganReponse struct {
	Nama  string `json:"name"`
	Email string `json:"email"`
}
type UpdateReponse struct {
	Nama      string `json:"name"`
	Email     string `json:"email"`
	Image_url string `json:"url"`
	Password  string `json:"Password"`
}

// type RegisterResponse struct {
// 	Nama  string `json:"name"`
// 	Email string `json:"email"`
// }
type LoginResponse struct {
	Nama  string `json:"nama"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func ToRegisterResponse(data pelanggan.PelangganEntites) PelangganReponse {
	return PelangganReponse{
		Nama:  data.Nama,
		Email: data.Email,
	}
}

func ToResponses(data pelanggan.PelangganEntites) PelangganReponse {
	return PelangganReponse{

		Nama:  data.Nama,
		Email: data.Email,
	}
}
func ToLoginRespon(data pelanggan.PelangganEntites, token string) LoginResponse {
	return LoginResponse{

		Nama:  data.Nama,
		Email: data.Email,
		Token: token,
	}
}
func UpdateRespons(data pelanggan.PelangganEntites) UpdateReponse {
	return UpdateReponse{

		Nama:      data.Nama,
		Email:     data.Email,
		Password:  data.Password,
		Image_url: data.Image_url,
	}
}
