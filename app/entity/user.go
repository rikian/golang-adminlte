package entity

type User struct {
	UserEmail    string `json:"user_email"`
	UserPassword string `json:"user_password"`
	UserStatus   string `json:"user_status"`
	UserSession  string `json:"user_session"`
}
