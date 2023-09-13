package main

import (
	"database/sql"
	"encoding/json"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"strconv"
	"time"
)

func SignupHandler(writer http.ResponseWriter, request *http.Request, db *sql.DB) {
	decoder := json.NewDecoder(request.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		return
	}
	if user.Name == "" {
		RespondWithError(writer, http.StatusBadRequest, "name field is empty")
		return
	}
	if user.Email == "" {
		RespondWithError(writer, http.StatusBadRequest, "email field is empty")
		return
	}
	if user.Password == "" {
		RespondWithError(writer, http.StatusBadRequest, "password field is empty")
		return
	}
	// No username that like should exists
	_, err = GetUserByName(db, user.Name)
	if err == nil {
		RespondWithError(writer, http.StatusBadRequest, "this name is taken, try another")
		return
	}
	// No email like that should exists
	_, err = GetUserByEmail(db, user.Email)
	if err == nil {
		RespondWithError(writer, http.StatusBadRequest, "this email is taken, try another")
		return
	}
	hashed_password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		RespondWithError(writer, http.StatusInternalServerError, "cannot encrypt password")
		return
	}
	newUser, err := InsertUser(db, user.Name, user.Email, string(hashed_password))
	if err != nil {
		log.Print(err)
	}
	encoder := json.NewEncoder(writer)
	encoder.Encode(newUser)
	return
}

func LoginHandler(writer http.ResponseWriter, request *http.Request, db *sql.DB) {

	type requestParam struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}
	decoder := json.NewDecoder(request.Body)
	var requestParams requestParam
	err := decoder.Decode(&requestParams)
	if err != nil {
		RespondWithError(writer, http.StatusInternalServerError, "cannot parse json")
	}
	if requestParams.Name == "" {
		RespondWithError(writer, http.StatusBadRequest, "name field is empty")
		return
	}
	if requestParams.Password == "" {
		RespondWithError(writer, http.StatusBadRequest, "password field is empty")
		return
	}
	// First check whether the name is in server or not
	user, err := GetUserByName(db, requestParams.Name)
	if err != nil {
		RespondWithError(writer, http.StatusBadRequest, "no user found")
		return
	}
	// Then check whether the password matches or not
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestParams.Password))
	if err != nil {
		RespondWithError(writer, http.StatusBadRequest, "password does not match")
		return
	}
	expirationTime := time.Now().Add(time.Hour * 4)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		Issuer:    "krishnansh-server",
		Subject:   strconv.Itoa(user.ID),
	})
	tokenString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		RespondWithError(writer, http.StatusInternalServerError, "cannot make the token string")
	}
	type response struct {
		Id    int    `json:"id"`
		Token string `json:"token"`
	}
	responseBody := response{Id: user.ID, Token: tokenString}
	RespondWithJSON(writer, 200, responseBody)
}
