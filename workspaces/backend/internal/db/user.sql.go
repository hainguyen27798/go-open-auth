// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: user.sql

package db

import (
	"context"
	"database/sql"
)

const createNewUser = `-- name: CreateNewUser :exec
INSERT INTO users (id, name, email, password, status, social_provider, image, verification_code)
VALUES (?, ?, ?, ?, ?, ?, ?, ?)
`

type CreateNewUserParams struct {
	ID               string
	Name             string
	Email            string
	Password         sql.NullString
	Status           UsersStatus
	SocialProvider   NullUsersSocialProvider
	Image            sql.NullString
	VerificationCode sql.NullString
}

func (q *Queries) CreateNewUser(ctx context.Context, arg CreateNewUserParams) error {
	_, err := q.db.ExecContext(ctx, createNewUser,
		arg.ID,
		arg.Name,
		arg.Email,
		arg.Password,
		arg.Status,
		arg.SocialProvider,
		arg.Image,
		arg.VerificationCode,
	)
	return err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, created_at, updated_at, name, email, password, status, social_provider, image, verify, verification_code
FROM users
WHERE email = ?
LIMIT 1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.Status,
		&i.SocialProvider,
		&i.Image,
		&i.Verify,
		&i.VerificationCode,
	)
	return i, err
}
