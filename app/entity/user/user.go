package user_entity

// for table user
type User struct {
	IdUser       string `gorm:"size:64; unique; not null" json:"user_id"`
	NameUser     string `gorm:"size:32; not null" json:"user_name"`
	EmailUser    string `gorm:"size:32; primaryKey; not null; unique" json:"user_email"`
	PasswordUser string `gorm:"size:255; not null" json:"user_password"`
	StatusUser   StatusUser
	StatusUserId int8   `gorm:"not null;" json:"user_status"`
	SessionUser  string `gorm:"size:255; not null; default:-" json:"user_session"`
	CreateDate   string `gorm:"size:64" json:"create_date"`
	LastUpdate   string `gorm:"size:64" json:"last_update"`
}

type StatusUser struct {
	Id     int8   `gorm:"not null;uniqueIndex;primaryKey"`
	Status string `gorm:"not null;unique"`
}

// for query
type UserRequestLogin struct {
	UserEmail    string `json:"user_email"`
	UserPassword string `json:"user_password"`
}

type UserLogin struct {
	UserEmail    string `json:"user_email"`
	UserPassword string `json:"user_password"`
	UserStatus   string `json:"user_status"`
	UserSession  string `json:"user_session"`
}

type UserRegister struct {
	UserName      string `json:"user_name"`
	UserEmail     string `json:"user_email"`
	UserPassword1 string `json:"user_password_1"`
	UserPassword2 string `json:"user_password_2"`
	UserTerm      bool   `json:"user_term"`
}

// for auth
type UserJwt struct {
	UserEmail string `json:"user_email"`
	TokenJwt  string `json:"token_jwt"`
}

type UserSession struct {
	SessionUser string `json:"user_session"`
}
