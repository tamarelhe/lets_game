// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: game.sql

package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/tabbed/pqtype"
)

const createGame = `-- name: CreateGame :one
INSERT INTO lg_games (
  id, group_id, type_id, datetime, members, location, constraints, message
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING id, group_id, type_id, datetime, members, location, constraints, message, created_at
`

type CreateGameParams struct {
	ID          uuid.UUID             `json:"id"`
	GroupID     uuid.UUID             `json:"group_id"`
	TypeID      uuid.NullUUID         `json:"type_id"`
	Datetime    time.Time             `json:"datetime"`
	Members     pqtype.NullRawMessage `json:"members"`
	Location    pqtype.NullRawMessage `json:"location"`
	Constraints pqtype.NullRawMessage `json:"constraints"`
	Message     sql.NullString        `json:"message"`
}

func (q *Queries) CreateGame(ctx context.Context, arg CreateGameParams) (LgGame, error) {
	row := q.db.QueryRowContext(ctx, createGame,
		arg.ID,
		arg.GroupID,
		arg.TypeID,
		arg.Datetime,
		arg.Members,
		arg.Location,
		arg.Constraints,
		arg.Message,
	)
	var i LgGame
	err := row.Scan(
		&i.ID,
		&i.GroupID,
		&i.TypeID,
		&i.Datetime,
		&i.Members,
		&i.Location,
		&i.Constraints,
		&i.Message,
		&i.CreatedAt,
	)
	return i, err
}

const deleteGame = `-- name: DeleteGame :exec
DELETE FROM lg_games
WHERE id = $1
`

func (q *Queries) DeleteGame(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteGame, id)
	return err
}

const getGame = `-- name: GetGame :one
SELECT id, group_id, type_id, datetime, members, location, constraints, message, created_at FROM lg_games
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetGame(ctx context.Context, id uuid.UUID) (LgGame, error) {
	row := q.db.QueryRowContext(ctx, getGame, id)
	var i LgGame
	err := row.Scan(
		&i.ID,
		&i.GroupID,
		&i.TypeID,
		&i.Datetime,
		&i.Members,
		&i.Location,
		&i.Constraints,
		&i.Message,
		&i.CreatedAt,
	)
	return i, err
}

const listGames = `-- name: ListGames :many
SELECT id, group_id, type_id, datetime, members, location, constraints, message, created_at FROM lg_games
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListGamesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListGames(ctx context.Context, arg ListGamesParams) ([]LgGame, error) {
	rows, err := q.db.QueryContext(ctx, listGames, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []LgGame{}
	for rows.Next() {
		var i LgGame
		if err := rows.Scan(
			&i.ID,
			&i.GroupID,
			&i.TypeID,
			&i.Datetime,
			&i.Members,
			&i.Location,
			&i.Constraints,
			&i.Message,
			&i.CreatedAt,
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

const updateGame = `-- name: UpdateGame :one
UPDATE lg_games
  set group_id = $2,
  type_id = $3,
  datetime = $4,
  members = $5,
  location = $6,
  constraints = $7,
  message = $8
WHERE id = $1
RETURNING id, group_id, type_id, datetime, members, location, constraints, message, created_at
`

type UpdateGameParams struct {
	ID          uuid.UUID             `json:"id"`
	GroupID     uuid.UUID             `json:"group_id"`
	TypeID      uuid.NullUUID         `json:"type_id"`
	Datetime    time.Time             `json:"datetime"`
	Members     pqtype.NullRawMessage `json:"members"`
	Location    pqtype.NullRawMessage `json:"location"`
	Constraints pqtype.NullRawMessage `json:"constraints"`
	Message     sql.NullString        `json:"message"`
}

func (q *Queries) UpdateGame(ctx context.Context, arg UpdateGameParams) (LgGame, error) {
	row := q.db.QueryRowContext(ctx, updateGame,
		arg.ID,
		arg.GroupID,
		arg.TypeID,
		arg.Datetime,
		arg.Members,
		arg.Location,
		arg.Constraints,
		arg.Message,
	)
	var i LgGame
	err := row.Scan(
		&i.ID,
		&i.GroupID,
		&i.TypeID,
		&i.Datetime,
		&i.Members,
		&i.Location,
		&i.Constraints,
		&i.Message,
		&i.CreatedAt,
	)
	return i, err
}
