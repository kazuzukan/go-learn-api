package auth

import (
	"errors"
	"os"

	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(userId int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
}

var SECRET_KEY = []byte(os.Getenv("JWT_SECRET_KEY"))

func NewServices() *jwtService {
	return &jwtService{}
}

func (s jwtService) GenerateToken(userId int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (s jwtService) ValidateToken(token string) (*jwt.Token, error) {
	tkn, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid token")
		}

		// akan ngereturn secretkey ketika methodnya tepat
		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return tkn, err
	}

	return tkn, nil
}
