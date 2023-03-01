package common

import (
	"crypto/ed25519"
	"errors"
	"personnel-management-backend/dao"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// var publicKey ed25519.PublicKey
// var privateKey ed25519.PrivateKey

var publicKey, privateKey, _ = ed25519.GenerateKey(nil)

type Claims struct {
	UserId uint
	jwt.RegisteredClaims
}

func CreateToken(user *dao.User) (string, error) {
	// 过期时间七天后
	expirationTime := &jwt.NumericDate{
		Time: time.Now().Add(7 * 24 * time.Hour),
	}

	claims := &Claims{
		UserId: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			// 过期时间
			ExpiresAt: expirationTime,
			// 签发时间
			IssuedAt: &jwt.NumericDate{
				Time: time.Now(),
			},
			// 签发标识
			Issuer:  "Garage",
			Subject: "user token",
		},
	}

	// 创建token
	token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, claims)
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	claims, ok := token.Claims.(*Claims)
	if ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid claims")
}
