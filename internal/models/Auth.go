package models

// Auth - Описание структуры для объекта аутентификации "Auth" ...
type Auth struct {
	Username string
	Password string
}

func NewAuth() *Auth {
	auth := new(Auth)
	return auth
}
