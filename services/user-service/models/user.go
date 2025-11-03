package models

import (
	"errors"
	"rearatrox/event-booking-api/services/user-service/db"
	"rearatrox/event-booking-api/services/user-service/utils"
)

type User struct {
	ID       int64
	Email    string `binding: "required"`
	Password string `binding: "required"`
}

func (u *User) ValidateCredentials() error {
	query := `SELECT id, Password FROM users where email=?`

	statement, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer statement.Close()

	result := db.DB.QueryRow(query, u.Email)
	if err != nil {
		return err
	}
	var hash []byte
	err = result.Scan(&u.ID, &hash)
	if err != nil {
		return err
	}

	isPasswordValid := utils.CheckPasswordHash(hash, u.Password)
	if !isPasswordValid {
		return errors.New("credentials invalid")
	}

	return nil
}

func (u *User) SaveUser() error {
	query := `INSERT INTO users(email, password) VALUES (?, ?)`

	statement, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer statement.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := statement.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	u.ID = id
	return err
}
