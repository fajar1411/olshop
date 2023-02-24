package user

type UserEntites struct {
	ID       uint
	Nama     string `validate:"required,min=5,required"`
	Password string `validate:"required,min=5,required"`
	Email    string `validate:"required,email"`
}

type UserService interface {
	Login(email, password string) (string, UserEntites, error)
	Register(newUser UserEntites) (UserEntites, error)
	Profile(id int) (UserEntites, error)
	UpdateUser(id int, Updata UserEntites) (UserEntites, error)
}

type UserData interface {
	Login(email string) (UserEntites, error)
	Register(newUser UserEntites) (UserEntites, error)
	Profile(id int) (UserEntites, error)
	UpdateUser(id int, Updata UserEntites) (UserEntites, error)
}
