package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService interface {
	GenerateToken(userID string, roles []string) (string, error)
	ValidateToken(tokenString string) (*Claims, error)
}

type authService struct {
	secretKey []byte
}

type Claims struct {
	UserID string   `json:"user_id"`
	Roles  []string `json:"roles"`
	jwt.RegisteredClaims
}

func NewAuthService(secretKey string) AuthService {
	return &authService{secretKey: []byte(secretKey)}
}

func (s *authService) GenerateToken(userID string, roles []string) (string, error) {
	claims := &Claims{
		UserID: userID,
		Roles:  roles,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.secretKey)
}

func (s *authService) ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (any, error) {
		return s.secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
