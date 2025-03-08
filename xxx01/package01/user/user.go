package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

type Admin struct {
	Email    string
	Password string
	User
}

func (user User) OutputUserDetails() {
	fmt.Println(user.firstName, user.lastName, user.birthdate)
}

func (admin Admin) OutputAdminDetails() {
	fmt.Println(admin.Email, admin.Password, admin.firstName, admin.lastName, admin.birthdate)
}

func GetUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}

func (user *User) ClearUserName() {
	user.firstName = ""
	user.lastName = ""
	user.birthdate = ""
}

func NewUser(userFirstName, userLastName, userBirthdate string) (*User, error) {
	if userFirstName == "" || userLastName == "" || userBirthdate == "" {
		return nil, errors.New("User data is missing")
	}
	return &User{
		firstName: userFirstName,
		lastName:  userLastName,
		birthdate: userBirthdate,
		createdAt: time.Now(),
	}, nil
}

func NewAdmin(userFirstName, userLastName, userBirthdate, email, password string) (*Admin, error) {
	if userFirstName == "" || userLastName == "" || userBirthdate == "" || email == "" || password == "" {
		return nil, errors.New("User data is missing")
	}
	return &Admin{
		Email:    email,
		Password: password,
		User: User{
			firstName: userFirstName,
			lastName:  userLastName,
			birthdate: userBirthdate,
			createdAt: time.Now(),
		},
	}, nil
}
