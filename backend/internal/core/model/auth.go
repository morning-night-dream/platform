package model

type Auth struct {
	UserID   string `json:"userId"`
	LoginID  string `json:"loginId"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
