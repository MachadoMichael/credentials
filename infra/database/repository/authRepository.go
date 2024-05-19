package repository

import (
	"database/sql"
	"errors"

	"github.com/MachadoMichael/GoAPI/schema"
)

type AuthRepository interface {
	Login(credentials schema.Credentials) (*schema.Credentials, error)
	Create(credentials schema.Credentials) error
	Delete(email string) error
	UpdatePassword(email string, oldPassword string, newPassword string) error
}

type BasicAuthRepo struct {
	db *sql.DB
}

func NewBasicAuthRepo(db *sql.DB) *BasicAuthRepo {
	return &BasicAuthRepo{db: db}
}

func (b *BasicAuthRepo) Login(credentials schema.Credentials) (*schema.Credentials, error) {
	var foundEmail string
	err := b.db.QueryRow("SELECT Email FROM users WHERE Email =? AND Password =?", credentials.Email, credentials.Password).Scan(&foundEmail)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("no such user")
		}
		return nil, err
	}
	return &schema.Credentials{Email: foundEmail}, nil
}

func (b *BasicAuthRepo) Create(credentials schema.Credentials) error {
	_, err := b.db.Exec("INSERT INTO users (Email, Password) VALUES (?,?)", credentials.Email, credentials.Password)
	return err
}

func (b *BasicAuthRepo) Delete(email string) error {

	result, err := b.db.Exec("DELETE FROM users WHERE Email =?", email)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("no user found with this email")
	}

	return nil
}

func (b *BasicAuthRepo) UpdatePassword(email string, oldPassword string, newPassword string) error {
	var foundEmail string
	err := b.db.QueryRow("SELECT Email FROM users WHERE Email =? AND Password =?", email, oldPassword).Scan(&foundEmail)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("old password does not match")
		}
		return err
	}

	_, err = b.db.Exec("UPDATE users SET Password =? WHERE Email =?", newPassword, email)
	return err
}
