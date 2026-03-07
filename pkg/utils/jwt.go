package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecretKey = []byte("scecret_key")

type CustomClaims struct {
	UserID    string `json:"user_id"`
	TokenType string `json:"token_type"`
	jwt.RegisteredClaims
}

func GenerateToken(userID, tokentype string, ttl time.Duration) (string, error) {
	claims := CustomClaims{
		UserID:    userID,
		TokenType: tokentype,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecretKey)
}

func GenerateTokenPair(userID string) (string, string, error) {
	access, err := GenerateToken(userID, "access", 2*time.Hour)
	if err != nil {
		return "", "", err
	}
	refresh, err := GenerateToken(userID, "refresh", 7*24*time.Hour)
	if err != nil {
		return "", "", err
	}
	return access, refresh, nil
}

func ParseAndValidate(tokenStr, expectedType string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretKey, nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return "", errors.New("invalid token")
	}
	if claims.TokenType != expectedType {
		return "", errors.New("invalid token type")
	}
	return claims.UserID, nil
}
