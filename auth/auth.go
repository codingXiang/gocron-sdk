package auth

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

const (
	Authorization = "Auth-Token"
	tokenDuration = 10 * time.Minute
)

type Request interface {
	Get() (string, error)
}

type Jwt struct {
	AuthSecret string
	UserId     int
	Username   string
}

func NewJwt(id int, name, secret string) *Jwt {
	log.Println("gocron jwt config, name = ", name , ", secret = ", secret[:len(secret) / 2], "...")
	return &Jwt{
		UserId:     id,
		Username:   name,
		AuthSecret: secret,
	}
}

func (h *Jwt) Get() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(tokenDuration).Unix()
	claims["uid"] = h.UserId
	claims["iat"] = time.Now().Unix()
	claims["issuer"] = "gocron"
	claims["username"] = h.Username
	claims["is_admin"] = 1
	token.Claims = claims
	return token.SignedString([]byte(h.AuthSecret))
}
