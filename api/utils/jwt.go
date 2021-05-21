package utils

import (
	"strings"
	"strconv"
	"os"
	"time"
	"jango/models"
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
		ID uint `json:"id"`
		Email string `json:"email"`
		Name string `json:"name"`
		Avatar string `json:"avatar"`
		Birthday int64 `json:"birthday"`
		Gender int `json:"gender"`
		City string `json:"city"`
		FacebookId string `json:"facebookId"`
		About string `json:"about"`
		jwt.StandardClaims
	}

func SignToken(user models.User) (string, error) {
	expirationTime := time.Now().Add(30 * 24 * time.Hour).Unix()

	claims := Claims{
			user.ID,
			user.Email,
			user.Name,
			user.Avatar,
			user.Birthday,
			user.Gender,
			user.City,
			user.FacebookId,
			user.About,
			jwt.StandardClaims{
				ExpiresAt: expirationTime,
			},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("TOKEN_SECRET")))

	return tokenString, err
}

func ValidateToken(tokenString string) (string) {
	splitToken := strings.Split(tokenString, " ")
	if len(splitToken) > 1 {
		tokenString = splitToken[1]
	} else {
		return ""
	}
	
	claims := Claims{}
	token, _ := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})
	
	if token != nil && token.Valid {
		return strconv.Itoa(int(claims.ID))
	} else {
		return ""
	}
}