package main

import (
	"fmt"

	"site.com/abc/user"
)

func main() {
	userFirstName := user.GetUserData("Please enter your first name: ")
	userLastName := user.GetUserData("Please enter your last name: ")
	userBirthdate := user.GetUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user.User

	appUser, err := user.NewUser(userFirstName, userLastName, userBirthdate)
	if err != nil {
		fmt.Println(err)
		return
	}

	admin := user.Admin{
		Email:    "booom@xyz.org",
		Password: "password",
		User:     *appUser,
	}

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()
	admin.OutputAdminDetails()
}
