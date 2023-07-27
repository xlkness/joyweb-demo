package token

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type RegisteredTokenClaims struct {
	User string `json:"user,omitempty"`
	jwt.RegisteredClaims
}

var (
	tokenSign = []byte("token_sign")
)

func ValidToken(token string) (*jwt.Token, error) {
	checkToken, err := jwt.ParseWithClaims(token, &RegisteredTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return tokenSign, nil
	})
	if err != nil {
		return nil, fmt.Errorf("ParseWithClaims token (%v) error:%v", token, err)
	} else {
		if !checkToken.Valid {
			return nil, fmt.Errorf("ParseWithClaims token (%v) invalid", token)
		}
		//else {
		//	claims := checkToken.Claims.(*jwt.RegisteredClaims)
		//	if claims.ExpiresAt.Before(time.Now()) {
		//		return nil, fmt.Errorf("token expired:%v", claims.ExpiresAt)
		//	}
		//}
	}

	return checkToken, nil
}

func GetToken(user string, expire time.Duration) (string, error) {
	claims := &RegisteredTokenClaims{
		User: user,
	}
	jwtClaims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(expire)),
		Issuer:    "test",
	}
	claims.RegisteredClaims = jwtClaims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(tokenSign)
	if err != nil {
		return "", fmt.Errorf("token SignedString error:%v", err)
	}

	return ss, nil
}
