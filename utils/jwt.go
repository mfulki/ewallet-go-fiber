package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mfulki/ewallet-go-fiber/entity"
)

func CreateAccessToken(user *entity.User, secretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Name,
		"iat":      time.Now().Unix(),
		"iss":      os.Getenv("ISSUER"),
		"exp":      time.Now().Add(10 * time.Minute).Unix(),
	})

	signed, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return signed, nil
}

func ParseAndVerify(signed string, secretKey string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(signed, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	}, jwt.WithIssuer(os.Getenv("ISSUER")),
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}),
		jwt.WithExpirationRequired(),
	)
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, nil
	} else {
		return nil, fmt.Errorf("unknown claims")
	}

}
