package jwt

import "github.com/golang-jwt/jwt"

type JWT struct {
	Secret string
}

func NewJwt(secret string) *JWT {
	return &JWT{
		Secret: secret,
	}
}

func (j *JWT) CreateJwt(number string) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"number": number,
	})
	s, err := jwtToken.SignedString([]byte(j.Secret))
	if err != nil {
		return "", err
	}
	return s, nil
}
