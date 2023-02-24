package handler

import "toko/fitur/user"

type UserReponse struct {
	Nama  string `json:"name"`
	Email string `json:"email"`
}
type UpdateReponse struct {
	Nama     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"Password"`
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

func ToRegisterResponse(data user.UserEntites) UserReponse {
	return UserReponse{
		Nama:  data.Nama,
		Email: data.Email,
	}
}

func ToResponses(data user.UserEntites) UserReponse {
	return UserReponse{

		Nama:  data.Nama,
		Email: data.Email,
	}
}
func ToLoginRespon(data user.UserEntites, token string) LoginResponse {
	return LoginResponse{

		Nama:  data.Nama,
		Email: data.Email,
		Token: token,
	}
}
func UpdateRespons(data user.UserEntites) UpdateReponse {
	return UpdateReponse{

		Nama:     data.Nama,
		Email:    data.Email,
		Password: data.Password,
	}
}
