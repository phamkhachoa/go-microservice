package auth

import (
	"github.com/golang-jwt/jwt"
	"time"
)

type PayloadClaims struct {
	jwt.StandardClaims
}

func GenTokenJWT(payload jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString([]byte("Secret"))
}

func CreateToken(username string) (string, error) {
	now := time.Now()
	exp := now.Add(time.Hour * 24 * 7)

	return GenTokenJWT(&PayloadClaims{
		StandardClaims: jwt.StandardClaims{
			Id:        username,
			ExpiresAt: exp.Unix(),
			IssuedAt:  now.Unix(),
			Issuer:    "auth",
		},
	})
}

func ParseJwtTokenSubject(token string) (*jwt.StandardClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(jwtToken *jwt.Token) (interface{}, error) {
		return []byte("Secret"), nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*jwt.StandardClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

// validate token

func VerifyTokenSubject(token string) (claims *jwt.StandardClaims, error error) {
	claims, err := ParseJwtTokenSubject(token)
	if err != nil {
		return nil, err
	}
	if err = claims.Valid(); err != nil {
		return nil, err
	}
	return claims, nil
}
