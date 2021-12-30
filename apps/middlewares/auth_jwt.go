package middlewares

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

func CreateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(viper.GetString(`jwt.secret`)))
}
