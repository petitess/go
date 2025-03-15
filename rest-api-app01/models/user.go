package models

import (
	"errors"

	"site.org/abc/db"
	"site.org/abc/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u User) Save() (User, error) {
	query := "INSERT INTO users(email,password) VALUES (?, ?)"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return u, err
	}
	defer stmt.Close()

	encryptedPass, err := utils.HashPassword(u.Password)
	if err != nil {
		return u, err
	}

	result, err := stmt.Exec(u.Email, encryptedPass)

	if err != nil {
		return u, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return u, err
	}
	u.ID = id
	u.Password = encryptedPass
	return u, nil
}

func GetAllUsers() ([]User, error) {
	query := "SELECT * FROM users"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func GetUserById(id int64) (*User, error) {
	query := "SELECT * FROM users WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var user User
	err := row.Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByEmail(email string) (*User, error) {
	query := "SELECT * FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, email)

	var user User
	err := row.Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (user User) Update() error {
	query := `
	UPDATE users
	SET email = ?, password = ?
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Email, user.Password, user.ID)
	return err
}

func (user User) Delete(id int64) error {
	query := `
	DELETE FROM users
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.ID)
	return err
}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)
	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)
	if err != nil {
		return errors.New("credentials invalid. " + err.Error())
	}

	passwordIsValid, err := utils.CheckPasswordHash(u.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("credentials invalid. " + err.Error())
	}
	return nil
}
