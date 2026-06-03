package src

import (
	"time"
	"errors"

	"github.com/golang-jwt/jwt/v5"
)



func NewJWTService (secret string) *JWTService {
	return &JWTService{
		secret: []byte(secret),
	}
}

func (s *JWTService) GenerateJWT (
	UserID string,
	Role UserRole,
) (string, error) {
	claims := Claims{
		UserID: UserID,
		Role: Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(
				time.Now().Add(24 * time.Hour),
			),
		},
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString(s.secret)
}

func (s *JWTService) VerifyJWT(
    tokenString string,
) (*Claims, error) {

    claims := &Claims{}

    token, err := jwt.ParseWithClaims(
        tokenString,
        claims,
        func(token *jwt.Token) (interface{}, error) {
			if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
				return nil, errors.New("unexpected signing method")
			}

			return s.secret, nil
		},
    )

    if err != nil {
        return nil, err
    }

    if !token.Valid {
        return nil, errors.New("invalid token")
    }

    return claims, nil
}