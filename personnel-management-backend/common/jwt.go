package common

import (
	"crypto/ed25519"
	"fmt"
	"personnel-management-backend/dao"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	privateKey ed25519.PrivateKey
	publicKey  ed25519.PublicKey
)

func init() {
	var err error
	publicKey, privateKey, err = ed25519.GenerateKey(nil)
	if err != nil {
		panic(err)
	}
}

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
		// panic(err)
		return nil, fmt.Errorf("failed to parse token: %v", err)
	}
	if !token.Valid {
		// panic(err)
		return nil, fmt.Errorf("token is invalid")
	}
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, fmt.Errorf("failed to get token claims")
	}
	return claims, nil
}

// func RefreshToken(tokenString string) (string, error) {
// 	// 解析 token，提取过期时间
// 	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (interface{}, error) {
// 		return publicKey, nil
// 	})
// 	if err != nil {
// 		return "", err
// 	}

// 	claims, ok := token.Claims.(*Claims)
// 	if !ok || !token.Valid {
// 		return "", errors.New("invalid token")
// 	}

// 	// 如果过期时间小于 5 分钟，则刷新 token
// 	if claims.ExpiresAt.Time.Sub(time.Now()) < 5*time.Minute {
// 		// 生成新的过期时间和签发时间
// 		newExpirationTime := time.Now().Add(7 * 24 * time.Hour)
// 		newIssuedAtTime := time.Now()

// 		// 更新 claims
// 		claims.ExpiresAt = &jwt.NumericDate{Time: newExpirationTime}
// 		claims.IssuedAt = &jwt.NumericDate{Time: newIssuedAtTime}

// 		// 创建新的 token
// 		newToken := jwt.NewWithClaims(jwt.SigningMethodEdDSA, claims)
// 		newTokenString, err := newToken.SignedString(privateKey)
// 		if err != nil {
// 			return "", err
// 		}

// 		return newTokenString, nil
// 	}

// 	return "", errors.New("token is still valid")
// }
