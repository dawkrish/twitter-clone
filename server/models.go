package main

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Tweet struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	Text      string `json:"text"`
	Timestamp string `json:"timestamp"`
	Username  string `json:"username"`
}
