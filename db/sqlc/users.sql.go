// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: users.sql

package sqlc

import (
	"context"
)

const changeEmail = `-- name: ChangeEmail :exec
UPDATE "users" SET
    email = $2
    where username = $1
`

type ChangeEmailParams struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (q *Queries) ChangeEmail(ctx context.Context, arg ChangeEmailParams) error {
	_, err := q.db.ExecContext(ctx, changeEmail, arg.Username, arg.Email)
	return err
}

const changePassword = `-- name: ChangePassword :exec
UPDATE "users" SET
    password = $2
    where username = $1
`

type ChangePasswordParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (q *Queries) ChangePassword(ctx context.Context, arg ChangePasswordParams) error {
	_, err := q.db.ExecContext(ctx, changePassword, arg.Username, arg.Password)
	return err
}

const createUser = `-- name: CreateUser :one
INSERT INTO "users"
(username, password, email) 
values ($1,$2,$3) RETURNING id, username, password, email, created_at
`

type CreateUserParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Username, arg.Password, arg.Email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Email,
		&i.CreatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM "users" where username = $1
`

func (q *Queries) DeleteUser(ctx context.Context, username string) error {
	_, err := q.db.ExecContext(ctx, deleteUser, username)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, username, password, email, created_at
FROM "users" where username = $1
`

func (q *Queries) GetUser(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Email,
		&i.CreatedAt,
	)
	return i, err
}
