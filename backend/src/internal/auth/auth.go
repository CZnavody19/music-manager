package auth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/CZnavody19/music-manager/src/db/config"
	"github.com/CZnavody19/music-manager/src/domain"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

type contextKey string

const (
	authHeader  = "Authorization"
	TokenCtxKey = contextKey("token")
	tokenLength = 64
	bcryptCost  = 12
)

type Auth struct {
	enabled     bool
	configStore *config.ConfigStore
	token       string
}

func NewAuth(cs *config.ConfigStore, enabled bool) (*Auth, error) {
	return &Auth{
		enabled:     enabled,
		configStore: cs,
	}, nil
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil

}

func checkHash(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func generateToken() string {
	b := make([]byte, tokenLength)
	rand.Read(b)

	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(b)
}

func (a *Auth) Middleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			token := r.Header.Get(authHeader)
			if token == "" {
				next.ServeHTTP(w, r)
				return
			}

			ctx = context.WithValue(ctx, TokenCtxKey, token)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func (a *Auth) Login(ctx context.Context, credentials *domain.Credentials) (string, error) {
	config, err := a.configStore.GetAuthConfig(ctx)
	if err != nil {
		return "", err
	}

	if credentials.Username != config.Username {
		return "", fmt.Errorf("invalid username")
	}

	err = checkHash(credentials.Password, config.PasswordHash)
	if err != nil {
		return "", fmt.Errorf("invalid password")
	}

	token := generateToken()

	a.token = token
	return token, nil
}

func (a *Auth) Logout(ctx context.Context) error {
	a.token = ""
	return nil
}

func (a *Auth) CheckToken(ctx context.Context, token string) error {
	if !a.enabled {
		return nil
	}

	if a.token != token {
		return fmt.Errorf("invalid token")
	}

	return nil
}
