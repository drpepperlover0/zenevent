package api

import (
	"time"

	"github.com/drpepperlover0/internal/structs"
	"github.com/golang-jwt/jwt/v5"
)

var (
	userSecretKey = []byte("snitzel")
	orgSecretKey  = []byte("sunnyBunny")
)

func GenerateUserJWT(user structs.User) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 2).Unix(),
		"sub": user.Username,
		"aud": structs.Role1,
	})

	tokenString, err := token.SignedString(userSecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GenerateOrgJWT(org structs.Organizer) (string, error) {

	orgToken := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 12).Unix(),
		"sub": org.IndividEmail,
		"aud": structs.Role2,
	})

	tokenString, err := orgToken.SignedString(orgSecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
