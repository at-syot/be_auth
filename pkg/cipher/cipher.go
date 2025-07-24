package cipher

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func CheckHashWithPassword(hash, pw string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
}

type JWTClaims struct {
	jwt.RegisteredClaims
}

const jwtSecret = "secret"

func JWTSign() (signedStr string, err error) {
	// SigningMethod: use HMAC for single entity like -one server to one client-
	// use ECDSA when we have other microservices to communicate with.
	claims := JWTClaims{jwt.RegisteredClaims{}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	signedStr, err = token.SignedString([]byte(jwtSecret))
	return
}

func JWTParse(tokenStr string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &JWTClaims{}, func(t *jwt.Token) (any, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		return nil, errors.New("jwt claims is malform.")
	}
	return claims, nil
}
