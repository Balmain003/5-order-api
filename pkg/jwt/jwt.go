package jwt

import "github.com/golang-jwt/jwt"

type JwtData struct {
	Phone string
}
type JWT struct {
	Secret string
}

func NewJwt(secret string) *JWT {
	return &JWT{
		Secret: secret,
	}
}

func (j *JWT) CreateJwt(phone string) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"phone": phone,
	})
	s, err := jwtToken.SignedString([]byte(j.Secret))
	if err != nil {
		return "", err
	}
	return s, nil
}

func (j *JWT) Parse(tokenString string) (bool, *JwtData) {
	claims := jwt.MapClaims{}

	t, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.Secret), nil
	})
	if err != nil || !t.Valid {
		return false, nil
	}

	phoneRaw, ok := claims["phone"]
	if !ok {
		return false, nil
	}
	phone, ok := phoneRaw.(string)
	if !ok {
		return false, nil
	}
	return true, &JwtData{
		Phone: phone,
	}
}
