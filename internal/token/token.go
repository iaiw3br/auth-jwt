package token

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	authError "main/internal/error"
	"time"
)

func CreateTokens(username string) (map[string]string, error) {
	tokens := make(map[string]string)

	accessToken, err := createToken(username, "SECRET_ACCESS", time.Minute)
	if err != nil {
		return nil, err
	}

	refreshToken, err := createToken(username, "SECRET_REFRESH", time.Hour)
	if err != nil {
		return nil, err
	}

	tokens["accessToken"] = accessToken
	tokens["refreshToken"] = refreshToken

	return tokens, nil
}

func createToken(username, secretKey string, exp time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Username": username,
		"exp":      time.Now().Add(exp).Unix(),
	})

	tokenString, err := token.SignedString([]byte(viper.GetString(secretKey)))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func parseToken(tokenString, key string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString(key)), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func ValidateToken(tokenString, key string) (*jwt.Token, error) {
	token, err := parseToken(tokenString, key)
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
			return nil, errors.New(authError.ErrInvalidToken)
		}
	} else {
		return nil, errors.New(authError.ErrInvalidToken)
	}
}

func CheckUpdateTokens(accessToken, refreshToken string) bool {
	_, accessTokenError := ValidateToken(accessToken, "SECRET_ACCESS")
	_, refreshTokenError := ValidateToken(refreshToken, "SECRET_REFRESH")

	if refreshTokenError == nil && (accessTokenError != nil || accessToken == "") {
		return true
	}

	return false
}

func GetUsernameFromToken(tokenString, key string) (string, error) {
	token, err := ValidateToken(tokenString, key)
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		username, ok := claims["Username"].(string)
		if !ok {
			return "", err
		}
		return username, nil
	}
	return "", nil
}
