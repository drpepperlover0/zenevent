package api

import (
	"time"

	"github.com/drpepperlover0/internal/structs"
	"github.com/golang-jwt/jwt/v5"
)

var (
	secretKey = []byte("snitzel")
)

type Claims struct {
	jwt.RegisteredClaims
	TokenType string `json:"type"`
	Username  string `json:"username,omitempty"`
	OrgName   string `json:"org_name,omitempty"`
}

func GenerateUserJWT(user structs.User) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		},
		TokenType: "user",
		Username:  user.Username,
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GenerateOrgJWT(org structs.Organizer) (string, error) {

	orgToken := jwt.NewWithClaims(jwt.SigningMethodHS512, Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
		},
		TokenType: "org",
		OrgName:   org.Name,
	})

	tokenString, err := orgToken.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseNameJWT(tokenString string) (string, error) {

	var claims Claims

	token, err := jwt.ParseWithClaims(tokenString, &claims, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if !token.Valid || err != nil {
		return "", err
	}

	if claims.TokenType == "user" {
		return claims.Username, nil
	} else if claims.TokenType == "org" {
		return claims.OrgName, nil
	}

	return "", nil
}
