-- name: GetGroup :one
SELECT * FROM lg_groups
WHERE id = $1 LIMIT 1;

-- name: ListGroups :many
SELECT * FROM lg_groups
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: CreateGroup :one
INSERT INTO lg_groups (
  id, name, avatar, members
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: UpdateGroup :one
UPDATE lg_groups
  set name = $2,
  avatar = $3,
  members = $4
WHERE id = $1
RETURNING *;

-- name: DeleteGroup :exec
DELETE FROM lg_groups
WHERE id = $1;