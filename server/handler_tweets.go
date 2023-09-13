package main

import (
	"database/sql"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
)

func GetAllTweetsHandler(writer http.ResponseWriter, request *http.Request, db *sql.DB) {
	tweets, err := GetAllTweets(db)
	//log.Print("tweets : ", tweets)
	if err != nil {
		RespondWithError(writer, http.StatusBadRequest, "cannot fetch tweets")
	}
	RespondWithJSON(writer, 200, tweets)
}

func GetTweetsByUserIDHandler(writer http.ResponseWriter, request *http.Request, db *sql.DB) {
	//log.Print(claims)
	userId, err := middlewareAuth(request)
	if err != nil {
		RespondWithError(writer, http.StatusUnauthorized, err.Error())
	}
	//log.Print("user id: ", userId)
	tweets, err := GetTweetsByUserID(db, userId)
	//log.Print("tweets : ", tweets)
	if err != nil {
		RespondWithError(writer, http.StatusBadRequest, "cannot fetch tweets")
	}
	RespondWithJSON(writer, 200, tweets)
}

func UpdateTweetHandler(writer http.ResponseWriter, request *http.Request, db *sql.DB) {
	tweetId := chi.URLParam(request, "tweetID")
	tweetIdNumerical, _ := strconv.Atoi(tweetId)

	userId, err := middlewareAuth(request)
	if err != nil {
		RespondWithError(writer, http.StatusUnauthorized, err.Error())
		return
	}
	type requestParams struct {
		Text string `json:"text"`
	}
	decoder := json.NewDecoder(request.Body)
	requestBody := requestParams{}
	err = decoder.Decode(&requestBody)
	if err != nil {
		RespondWithError(writer, http.StatusBadRequest, err.Error())
		return
	}
	if requestBody.Text == "" {
		RespondWithError(writer, http.StatusBadRequest, "the new content is empty, enter something")
		return
	}
	oldTweet, err := GetTweetByTweetID(db, tweetIdNumerical)
	if err != nil {
		RespondWithError(writer, http.StatusBadRequest, "tweet does not exist")
		return
	}
	log.Println(oldTweet)
	if oldTweet.UserID != userId {
		RespondWithError(writer, http.StatusUnauthorized, "you are not the author of this tweet, you cannot change this tweet")
		return
	}
	newTweet, err := UpdateTweet(db, tweetIdNumerical, requestBody.Text)
	if err != nil {
		RespondWithError(writer, http.StatusBadRequest, err.Error())
		return
	}
	RespondWithJSON(writer, 200, newTweet)
}

func PostTweetHandler(writer http.ResponseWriter, request *http.Request, db *sql.DB) {
	userId, err := middlewareAuth(request)
	if err != nil {
		RespondWithError(writer, http.StatusUnauthorized, err.Error())
		return
	}
	type requestParams struct {
		Text string `json:"text"`
	}
	decoder := json.NewDecoder(request.Body)
	requestBody := requestParams{}
	err = decoder.Decode(&requestBody)
	if err != nil {
		RespondWithError(writer, http.StatusBadRequest, err.Error())
		return
	}
	if requestBody.Text == "" {
		RespondWithError(writer, http.StatusBadRequest, "the new content is empty, enter something")
		return
	}
	newTweet, err := InsertTweet(db, userId, requestBody.Text)
	if err != nil {
		RespondWithError(writer, 200, err.Error())
		return
	}
	RespondWithJSON(writer, 200, newTweet)
}

func DeleteTweetHandler(writer http.ResponseWriter, request *http.Request, db *sql.DB) {
	log.Print("in delete tweet handler")
	tweetId := chi.URLParam(request, "tweetID")
	tweetIdNumerical, _ := strconv.Atoi(tweetId)

	_, err := middlewareAuth(request)
	if err != nil {
		RespondWithError(writer, http.StatusUnauthorized, err.Error())
		return
	}
	err = DeleteTweet(db, tweetIdNumerical)
	if err != nil {
		RespondWithError(writer, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJSON(writer, 200, "Deleted successfully")
}
