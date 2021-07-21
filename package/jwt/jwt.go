package jwt

import (
	"time"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"github.com/golang-jwt/jwt"
)

type Claim struct {
	Id int `json:"id"`
	jwt.StandardClaims
}

func GenerateToken(id int) (string, error) {
	claim := Claim{id, jwt.StandardClaims{
		ExpiresAt: time.Now().Unix() + g.Cfg().GetInt64("api.ExpiredAt"),
		Issuer:    "go-admin",
	}}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	ss, err := token.SignedString(gconv.Bytes(g.Cfg().Get("api.Key")))

	return ss, err
}

func ParseToken(tokenString string) (interface{}, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claim{}, func(token *jwt.Token) (interface{}, error) {
		return gconv.Bytes(g.Cfg().Get("api.Key")), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claim); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
