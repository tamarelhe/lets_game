-- name: GetGame :one
SELECT * FROM lg_games
WHERE id = $1 LIMIT 1;

-- name: ListGames :many
SELECT * FROM lg_games
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: CreateGame :one
INSERT INTO lg_games (
  id, group_id, type_id, datetime, members, location, constraints, message
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING *;

-- name: UpdateGame :one
UPDATE lg_games
  set group_id = $2,
  type_id = $3,
  datetime = $4,
  members = $5,
  location = $6,
  constraints = $7,
  message = $8
WHERE id = $1
RETURNING *;

-- name: DeleteGame :exec
DELETE FROM lg_games
WHERE id = $1;