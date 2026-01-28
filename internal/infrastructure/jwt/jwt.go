package jwt

import (
	"errors"
	"go-mini-admin/config"
	"go-mini-admin/internal/infrastructure/logger"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Clainms struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type JWTManager struct {
	secret string
	expire int // hours
}

func New(cfg *config.JwtConfig) *JWTManager {
	return &JWTManager{
		secret: cfg.Secret,
		expire: cfg.Expire,
	}
}

func (j *JWTManager) GenerateToken(userID uint, username string) (string, error) {
	claims := Clainms{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(j.expire) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "go-mini-admin",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString([]byte(j.secret))
}

func (j *JWTManager) ParseToken(tokenString string) (*Clainms, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Clainms{}, func(token *jwt.Token) (any, error) {
		return []byte(j.secret), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Clainms); ok && token.Valid {
		return claims, nil
	}

	logger.Warn("invalid token")
	return nil, errors.New("invalid token")
}
