package service

type LoginService interface {
	Login(username string, password string) bool
}

type loginService struct {
	username string
	password string
}

func NewLoginService() LoginService {
	return &loginService{
		username: "admin",
		password: "@#$RF@!718",
	}
}

func (service *loginService) Login(username string, password string) bool {
	return service.username == username && service.password == password
}
