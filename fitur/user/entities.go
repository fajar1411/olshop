package user

type UserEntites struct {
	ID       uint
	Nama     string
	Password string
	Email    string
}

type UserService interface {
	Login(email, password string) (string, UserEntites, error)
	Register(newUser UserEntites) (UserEntites, error)
}

type UserData interface {
	Login(password string) (UserEntites, error)
	Register(newUser UserEntites) (UserEntites, error)
}
