package domain

type AuthConfig struct {
	Username     string
	PasswordHash string
}

type Credentials struct {
	Username string
	Password string
}
