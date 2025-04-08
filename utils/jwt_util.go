package utils

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	// 密钥
	secretKey = "123456"
	// token有效期（12小时）
	tokenExpire = time.Hour * 12
)

// 自定义Claims
type Claims struct {
	UserID   int    `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenerateToken 生成token
func GenerateToken(userID int, username string) (string, error) {
	claims := Claims{
		UserID:   userID,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenExpire).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	// 生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

// ParseToken 解析token
func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("无效的Token")
}
