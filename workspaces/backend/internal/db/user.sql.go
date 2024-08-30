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
SELECT id, email FROM users WHERE email = ? LIMIT 1
`

type GetUserByEmailRow struct {
	ID    string
	Email string
}

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (GetUserByEmailRow, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i GetUserByEmailRow
	err := row.Scan(&i.ID, &i.Email)
	return i, err
}
