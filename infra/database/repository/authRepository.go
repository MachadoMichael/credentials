package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/MachadoMichael/GoAPI/schema"
	"github.com/go-redis/redis/v8"
)

type AuthRepository interface {
	Login(credentials schema.Credentials) (*schema.Credentials, error)
	Create(credentials schema.Credentials) error
	Delete(email string) error
	UpdatePassword(email string, oldPassword string, newPassword string) error
}

type BasicAuthRepo struct {
	ctx context.Context
	db  *redis.Client
}

func NewBasicAuthRepo(ctx context.Context, db *redis.Client) *BasicAuthRepo {
	return &BasicAuthRepo{ctx: ctx, db: db}
}

func (b *BasicAuthRepo) Login(credentials schema.Credentials) (*schema.Credentials, error) {
	var foundEmail string
	result, err := b.db.Set(b.ctx, credentials.Email, credentials.Password).Result()

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
