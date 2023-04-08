package auth

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(userid string) (string, error)
	ValidateToken(encodeToken string) (*jwt.Token, error)
}

type jwtService struct {
}

func NewService() *jwtService {
	return &jwtService{}
}

var SECRET_KEY = []byte("Kode_Rahasia_just_for_dummY_Exampl3")

func (j *jwtService) GenerateToken(userid string) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userid

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedMethod, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedMethod, err
	}

	return signedMethod, nil
}

func (j *jwtService) ValidateToken(encodeToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodeToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}
