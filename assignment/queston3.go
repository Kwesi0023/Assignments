package main


import (
	"gorm.io/gorm"
)
	
type User struct {
	ID       int	`json:"id"`
	Profile  Profile  `json:"profile"`
	Settings Settings `json:"settings"`
}

type Profile struct {
	Username string	`json:"username"`
	Email    string  `json:"email"`
}

type Settings struct {
	Theme         string	`json:"theme"`
	Notifications bool		`json:"notifications"`
}

func PrintDetails(name string, mail string) (string, string) {
	userDetails := Profile{Username: "Kwesi", Email: "kkwesi@gmail.com"}
	var userDetails string
	names := userDetails.Username
	mails := userDetails.Email
	return names, mails
}

func theThemes(colors string) string {
	choose := Settings{Theme: "Dark", Notifications: true}
	var choose string
	colors := choose.Theme
	return colors
}
