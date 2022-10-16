-- name: GetUser :one
SELECT * FROM lg_users
WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM lg_users
WHERE email = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM lg_users
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: CreateUser :one
INSERT INTO lg_users (
  id, name, email, password, avatar, is_active, groups
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
)
RETURNING *;

-- name: UpdateUser :one
UPDATE lg_users
  set name = $2,
  email = $3,
  avatar = $4,
  groups = $5
WHERE id = $1
RETURNING *;

-- name: UpdateUserPassword :exec
UPDATE lg_users
  set password = $2
WHERE id = $1;

-- name: InactivateUser :exec
UPDATE lg_users
  set is_active = false
WHERE id = $1;

-- name: ActivateUser :exec
UPDATE lg_users
  set is_active = true
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM lg_users
WHERE id = $1;