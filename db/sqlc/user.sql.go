// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: user.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

const activateUser = `-- name: ActivateUser :exec
UPDATE lg_users
  set is_active = true
WHERE id = $1
`

func (q *Queries) ActivateUser(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, activateUser, id)
	return err
}

const createUser = `-- name: CreateUser :one
INSERT INTO lg_users (
  id, name, email, password, avatar, is_active, groups
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
)
RETURNING id, name, email, password, avatar, is_active, created_at, groups
`

type CreateUserParams struct {
	ID       uuid.UUID      `json:"id"`
	Name     string         `json:"name"`
	Email    string         `json:"email"`
	Password string         `json:"password"`
	Avatar   sql.NullString `json:"avatar"`
	IsActive bool           `json:"is_active"`
	Groups   []string       `json:"groups"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (LgUser, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.ID,
		arg.Name,
		arg.Email,
		arg.Password,
		arg.Avatar,
		arg.IsActive,
		pq.Array(arg.Groups),
	)
	var i LgUser
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.Avatar,
		&i.IsActive,
		&i.CreatedAt,
		pq.Array(&i.Groups),
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM lg_users
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, name, email, password, avatar, is_active, created_at, groups FROM lg_users
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id uuid.UUID) (LgUser, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i LgUser
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.Avatar,
		&i.IsActive,
		&i.CreatedAt,
		pq.Array(&i.Groups),
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, name, email, password, avatar, is_active, created_at, groups FROM lg_users
WHERE email = $1 LIMIT 1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (LgUser, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i LgUser
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.Avatar,
		&i.IsActive,
		&i.CreatedAt,
		pq.Array(&i.Groups),
	)
	return i, err
}

const inactivateUser = `-- name: InactivateUser :exec
UPDATE lg_users
  set is_active = false
WHERE id = $1
`

func (q *Queries) InactivateUser(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, inactivateUser, id)
	return err
}

const listUsers = `-- name: ListUsers :many
SELECT id, name, email, password, avatar, is_active, created_at, groups FROM lg_users
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListUsersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListUsers(ctx context.Context, arg ListUsersParams) ([]LgUser, error) {
	rows, err := q.db.QueryContext(ctx, listUsers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []LgUser{}
	for rows.Next() {
		var i LgUser
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Email,
			&i.Password,
			&i.Avatar,
			&i.IsActive,
			&i.CreatedAt,
			pq.Array(&i.Groups),
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUser = `-- name: UpdateUser :one
UPDATE lg_users
  set name = $2,
  email = $3,
  avatar = $4,
  groups = $5
WHERE id = $1
RETURNING id, name, email, password, avatar, is_active, created_at, groups
`

type UpdateUserParams struct {
	ID     uuid.UUID      `json:"id"`
	Name   string         `json:"name"`
	Email  string         `json:"email"`
	Avatar sql.NullString `json:"avatar"`
	Groups []string       `json:"groups"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (LgUser, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.ID,
		arg.Name,
		arg.Email,
		arg.Avatar,
		pq.Array(arg.Groups),
	)
	var i LgUser
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.Avatar,
		&i.IsActive,
		&i.CreatedAt,
		pq.Array(&i.Groups),
	)
	return i, err
}

const updateUserPassword = `-- name: UpdateUserPassword :exec
UPDATE lg_users
  set password = $2
WHERE id = $1
`

type UpdateUserPasswordParams struct {
	ID       uuid.UUID `json:"id"`
	Password string    `json:"password"`
}

func (q *Queries) UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) error {
	_, err := q.db.ExecContext(ctx, updateUserPassword, arg.ID, arg.Password)
	return err
}
