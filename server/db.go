package main

import (
	"database/sql"
	"fmt"
	"log"
)

func ConnectToDB() *sql.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/firstdb"

	// Open a connection to the database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	// Test the database connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Connected to the database")
	return db
}

func CreateTableUsers(db *sql.DB) {
	tableQuery := `
		CREATE TABLE IF NOT EXISTS users (
			id INT AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL,
			password VARCHAR(255) NOT NULL
		)
	`
	_, err := db.Exec(tableQuery)
	if err != nil {
		log.Fatal("error in creating table users ", err)
	}
	log.Print("User table created")
}

func CreateTableFollowers(db *sql.DB) {
	tableQuery := `CREATE TABLE IF NOT EXISTS followers(
    					user_id INT NOT NULL PRIMARY KEY,
    					follower_id INT NOT NULL,
    					FOREIGN KEY (user_id) REFERENCES users(id),
    					FOREIGN KEY(follower_id) REFERENCES users(id)
					)`
	_, err := db.Exec(tableQuery)
	if err != nil {
		log.Fatal("error in creating table followers ", err)
	}
	log.Print("Followers table created")
}

func CreateTableTweets(db *sql.DB) {
	tableQuery := `
		CREATE TABLE IF NOT EXISTS tweets (
			id INT AUTO_INCREMENT PRIMARY KEY,
			user_id INT,
			text TEXT,
			timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (user_id) REFERENCES users(id)
		)
	`
	_, err := db.Exec(tableQuery)
	if err != nil {
		log.Fatal("error in creating table tweets ", err)
	}
	log.Print("Tweets table created")
}

func InsertUser(db *sql.DB, name, email, password string) (User, error) {
	insertQuery := "INSERT INTO users (name, email, password) VALUES (?, ?, ?)"

	result, err := db.Exec(insertQuery, name, email, password)
	if err != nil {
		log.Fatal("error inserting in user ", err)
		return User{}, err
	}
	log.Print("User inserted successfully")
	insertedID, _ := result.LastInsertId()

	return GetUser(db, int(insertedID))
}

func InsertFollower(db *sql.DB, user_id, follower_id int) {
	insertQuery := "INSERT INTO followers(user_id, follower_id) values (?,?)"
	_, err := db.Exec(insertQuery, user_id, follower_id)
	if err != nil {
		log.Fatal("error inserting in follower ", err)
	}
	log.Print("Follower inserted successfully")
}

func InsertTweet(db *sql.DB, user_id int, text string) (Tweet, error) {
	insertQuery := `
		INSERT INTO tweets (user_id, text)
		VALUES (?, ?)
	`
	result, err := db.Exec(insertQuery, user_id, text)
	if err != nil {
		log.Fatal("error in inserting tweet ", err)
	}
	log.Print("Tweet inserted successfully")
	insertedID, _ := result.LastInsertId()
	return GetTweetByTweetID(db, int(insertedID))
}

func GetFollowers(db *sql.DB, user_id int) {
	// we need to get followers!
	selectQuery := "SELECT follower_id FROM followers WHERE user_id = ?"
	rows, err := db.Query(selectQuery, user_id)
	if err != nil {
		log.Fatal("error querying followers: ", err)
	}
	defer rows.Close()

	followersIDArr := make([]int, 0) // Use make to create an empty slice

	for rows.Next() {
		var followerID int
		if err := rows.Scan(&followerID); err != nil {
			log.Fatal(err)
		}
		followersIDArr = append(followersIDArr, followerID)
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	// Now you can use followersIDArr as needed
	fmt.Printf("Followers ID Array for user_id = %v is -> %v\n", user_id, followersIDArr)
}

func GetFollowing(db *sql.DB, user_id int) {
	// we need to get followers!
	selectQuery := "SELECT user_id FROM followers WHERE follower_id = ?"
	rows, err := db.Query(selectQuery, user_id)
	if err != nil {
		log.Fatal("error querying following: ", err)
	}
	defer rows.Close()

	followingIDArr := make([]int, 0) // Use make to create an empty slice

	for rows.Next() {
		var followingID int
		if err := rows.Scan(&followingID); err != nil {
			log.Fatal(err)
		}
		followingIDArr = append(followingIDArr, followingID)
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Following ID Array for user_id = %v is -> %v\n", user_id, followingIDArr)
}

func GetUser(db *sql.DB, user_id int) (User, error) {
	selectQuery := "SELECT id, name, email, password FROM users WHERE id = ?"
	row := db.QueryRow(selectQuery, user_id)

	var user User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {

		return User{}, err
	}

	return user, nil
}
func GetUserByName(db *sql.DB, name string) (User, error) {
	selectQuery := "SELECT id, name, email, password FROM users WHERE name = ?"
	row := db.QueryRow(selectQuery, name)

	var user User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func GetUserByEmail(db *sql.DB, email string) (User, error) {
	selectQuery := "SELECT id, name, email, password FROM users WHERE email = ?"
	row := db.QueryRow(selectQuery, email)

	var user User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func GetTweetByTweetID(db *sql.DB, tweetID int) (Tweet, error) {
	selectQuery := "SELECT id, user_id, text, timestamp FROM tweets WHERE id = ?"
	row := db.QueryRow(selectQuery, tweetID)

	var tweet Tweet
	err := row.Scan(&tweet.ID, &tweet.UserID, &tweet.Text, &tweet.Timestamp)
	if err != nil {
		return Tweet{}, err
	}

	return tweet, nil
}

func GetAllTweets(db *sql.DB) ([]Tweet, error) {
	selectQuery := "SELECT * FROM TWEETS"
	rows, err := db.Query(selectQuery)
	if err != nil {
		return []Tweet{}, err
	}

	var tweets []Tweet

	for rows.Next() {
		var tweet Tweet
		err := rows.Scan(&tweet.ID, &tweet.UserID, &tweet.Text, &tweet.Timestamp)
		if err != nil {
			return []Tweet{}, err
		}
		user, err := GetUser(db, tweet.UserID)
		if err != nil {
			return []Tweet{}, err
		}
		tweet.Username = user.Name
		tweets = append(tweets, tweet)
	}

	return tweets, nil
}

func GetTweetsByUserID(db *sql.DB, user_id int) ([]Tweet, error) {
	selectQuery := "SELECT * FROM TWEETS WHERE user_id = ?"
	rows, err := db.Query(selectQuery, user_id)
	if err != nil {
		return []Tweet{}, err
	}

	var tweets []Tweet

	for rows.Next() {
		var tweet Tweet
		err := rows.Scan(&tweet.ID, &tweet.UserID, &tweet.Text, &tweet.Timestamp)
		if err != nil {
			return []Tweet{}, err
		}
		user, err := GetUser(db, tweet.UserID)
		if err != nil {
			return []Tweet{}, err
		}
		tweet.Username = user.Name
		tweets = append(tweets, tweet)
	}
	return tweets, nil
}

func UpdateTweet(db *sql.DB, tweetID int, text string) (Tweet, error) {
	updateQuery := "UPDATE TWEETS SET text = ? where id = ?"
	_, err := db.Exec(updateQuery, text, tweetID)
	if err != nil {
		log.Print("error updating tweet")
		return Tweet{}, err
	}
	log.Print("Tweet updated successfully")

	return GetTweetByTweetID(db, tweetID)
}

func DeleteTweet(db *sql.DB, tweetID int) error {
	deleteQuery := "DELETE FROM TWEETS WHERE id = ?"
	_, err := db.Exec(deleteQuery, tweetID)
	if err != nil {
		log.Print("error updating tweet")
		return err
	}
	log.Print("Tweet deleted successfully")
	return nil
}
