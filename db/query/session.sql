-- name: GetSession :one
SELECT * FROM lg_sessions
WHERE id = $1 LIMIT 1;


-- name: CreateSession :one
INSERT INTO lg_sessions (
  id, email, refresh_token, user_agent, client_ip, is_blocked, expires_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
)
RETURNING *;