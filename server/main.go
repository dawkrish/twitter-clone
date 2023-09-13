package main

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

var db *sql.DB
var jwtSecretKey = []byte("secret")

func main() {
	r := chi.NewRouter()
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})
	r.Use(corsMiddleware.Handler)
	db = ConnectToDB()
	defer db.Close()

	r.Get("/api", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(`{"name": "hello"}`))
	})
	r.Get("/loginstate", func(writer http.ResponseWriter, request *http.Request) {
		_, err := middlewareAuth(request)
		if err != nil {
			RespondWithError(writer, http.StatusUnauthorized, "not logged in")
		} else {
			type response struct {
				Message string `json:"message"`
			}
			responseBody := response{
				Message: "loggedin",
			}
			RespondWithJSON(writer, 200, responseBody)
		}
	})
	r.Post("/signup", func(writer http.ResponseWriter, request *http.Request) {
		SignupHandler(writer, request, db)
	})
	r.Post("/login", func(writer http.ResponseWriter, request *http.Request) {
		LoginHandler(writer, request, db)
	})
	r.Get("/alltweets", func(writer http.ResponseWriter, request *http.Request) {
		GetAllTweetsHandler(writer, request, db)
	})

	r.Get("/tweets", func(writer http.ResponseWriter, request *http.Request) {
		GetTweetsByUserIDHandler(writer, request, db)
	})
	r.Post("/tweets", func(writer http.ResponseWriter, request *http.Request) {
		PostTweetHandler(writer, request, db)
	})
	r.Put("/tweets/{tweetID}", func(writer http.ResponseWriter, request *http.Request) {
		UpdateTweetHandler(writer, request, db)
	})

	r.Delete("/tweets/{tweetID}", func(writer http.ResponseWriter, request *http.Request) {
		DeleteTweetHandler(writer, request, db)
	})

	//CreateTableUsers(db)
	//CreateTableFollowers(db)
	//CreateTableFollowing(db)
	//CreateTableTweets(db)
	//InsertUser(db, "unnat", "b@b.com", "root2")
	//GetFollowers(db, 1)
	//GetFollowing(db, 2)
	//log.Print(InsertTweet(db, 7, "Reece's first tweet"))
	//log.Print(InsertTweet(db, 7, "Reece's second tweet"))
	//log.Print(InsertUser(db, "vansh", "c@c.com", "root3"))
	//log.Println(GetUser(db, 2))
	//log.Println(GetTweet(db, 1))
	//log.Print(GetTweetsByUserID(db, 7))
	log.Println("Listening at localhost:8080")
	http.ListenAndServe(":8080", r)
}
