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
	Profile(id int) (UserEntites, error)
}

type UserData interface {
	Login(email string) (UserEntites, error)
	Register(newUser UserEntites) (UserEntites, error)
	Profile(id int) (UserEntites, error)
}
