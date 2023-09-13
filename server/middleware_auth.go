package main

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func middlewareAuth(request *http.Request) (int, error) {
	authorizationString := request.Header.Get("Authorization")
	if authorizationString == "" {
		return -1, errors.New("token not present")
	}
	//log.Print("Authorization String : ", authorizationString)
	tokenString := strings.TrimPrefix(authorizationString, "Bearer ")
	//log.Print("Token String: ", tokenString)
	claims := jwt.StandardClaims{}

	// Checking whether this a valid jwt or not
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretKey, nil
	})
	if err != nil {
		//log.Print("inside valid jwt or not")
		return -1, err
	}

	// checking whether the token is valid or not
	if !token.Valid {
		//log.Printf("whether token is valid or not")
		return -1, err
	}
	if time.Now().UTC().Unix() > claims.ExpiresAt {
		return -1, errors.New("token has expired")
	}
	userId, _ := strconv.Atoi(claims.Subject)
	return userId, nil
}
