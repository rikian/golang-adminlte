package entity

type User struct {
	UserEmail    string `json:"user_emai"`
	UserPassword string `json:"user_password"`
	UserStatus   string
	UserSession  string
}
