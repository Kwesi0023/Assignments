package main

type User struct {
	ID       int
	Profile  Profile
	Settings Settings
}

type Profile struct {
	Username string
	Email    string
}

type Settings struct {
	Theme         string
	Notifications bool
}

func PrintDetails(name string, mail string) (string, string) {
	//userDetails := Profile{Username: "Kwesi", Email: "kkwesi@gmail.com"}
	var userDetails string
	names := userDetails.Username
	mails := userDetails.Email
	return names, mails
}

func theThemes(colors string) string {
	choose := Settings{Theme: "Dark", Notifications: true}
	//var choose string
	colors := choose.Theme
	return colors
}
