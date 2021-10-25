package token

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	authError "main/internal/error"
	"time"
)

func CreateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Username": username,
		"exp":      time.Now().Add(time.Second).Unix(),
	})

	accessSecret := viper.GetString("SECRET_ACCESS")

	tokenString, err := token.SignedString([]byte(accessSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func parseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("SECRET_ACCESS")), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := parseToken(tokenString)
	if err != nil {
		return nil, err
	}

	if token.Valid {
		return token, nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return nil, errors.New(authError.ErrNotEvenAToken)
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return nil, errors.New(authError.ErrTokenIsExpired)
		} else {
			return nil, errors.New(authError.ErrorInvalidToken)
		}
	} else {
		return nil, errors.New(authError.ErrorInvalidToken)
	}
}
