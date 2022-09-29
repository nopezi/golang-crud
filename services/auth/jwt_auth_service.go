package services

import (
	"errors"
	env "riskmanagement/lib/env"
	models "riskmanagement/models/user"
	"time"

	"github.com/dgrijalva/jwt-go"
	"gitlab.com/golang-package-library/logger"
)

// JWTAuthService service relating to authorization
type JWTAuthService struct {
	env    env.Env
	logger logger.Logger
}

// NewJWTAuthService creates a new auth service
func NewJWTAuthService(env env.Env, logger logger.Logger) JWTAuthService {
	return JWTAuthService{
		env:    env,
		logger: logger,
	}
}

// Authorize authorizes the generated token
func (s JWTAuthService) Authorize(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(s.env.JWTSecret), nil
	})
	if token.Valid {
		return true, nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return false, errors.New("token malformed")
		}
		if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return false, errors.New("token expired")
		}
	}
	return false, errors.New("couldn't handle token")
}

// CreateToken creates jwt auth token
func (s JWTAuthService) CreateToken(user models.User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID,
		"name":  user.Name,
		"email": *user.Email,
	})

	tokenString, err := token.SignedString([]byte(s.env.JWTSecret))

	if err != nil {
		s.logger.Zap.Error("JWT validation failed: ", err)
	}

	return tokenString
}

// CreateTokenGlobal creates jwt auth token
func (s JWTAuthService) CreateTokenGlobal() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"appID":   "v2",
		"appName": "riskmanagement",
		"exp":     time.Now().Add(time.Minute * 60).Unix(),
	})

	tokenString, err := token.SignedString([]byte(s.env.JWTSecret))

	if err != nil {
		s.logger.Zap.Error("JWT validation failed: ", err)
	}

	return tokenString
}
