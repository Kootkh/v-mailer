package types

/* type Auth struct {
	Username string `yaml:"AUTH_USERNAME" env:"AUTH_USERNAME" validate:"omitempty,required_with=password"`
	Password string `yaml:"AUTH_PASSWORD" env:"AUTH_PASSWORD" validate:"omitempty,required_with=username"`
}

func (*Auth) String() string {
	return ""
}

func (a *Auth) IsSet() bool {
	if a.Username != "" && a.Password != "" {
		return true
	}
	return false
}

func (a *Auth) Set(ai *AuthItem) error {
	a.setUsername(ai.Username)
	a.setPassword(ai.Password)
	return nil
} */

type AuthItem struct {
	Username string `validate:"required_with=Password"`
	Password string `validate:"required_with=Username"`
}

func NewAuthItem() *AuthItem {
	return &AuthItem{}
}

/* func (a *Auth) setUsername(value string) error {
	a.Username = strings.TrimSpace(value)
	return nil
}

func (a *Auth) setPassword(value string) error {
	a.Password = strings.TrimSpace(value)
	return nil
} */
