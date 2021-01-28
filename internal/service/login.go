package service

//LoginService interface
type LoginService interface {
	Login(username, password string) bool
}

type loginService struct {
	username string
	password string
}

//NewLoginService to return static username and password
func NewLoginService() LoginService {
	return &loginService{
		username: "Something",
		password: "Password",
	}
}

//Login function to verify username and password
func (s *loginService) Login(username, password string) bool {
	return s.username == username && s.password == password
}

//Another way
//func Login(username, password string) bool {
//	return username == "Something" && password == "Password"
//}
