-- name: CreateRefreshToken :one
INSERT INTO refresh_tokens (user_id, token, created_at, expires_at, updated_at)
VALUES ($1, $2, NOW(), NOW() + INTERVAL '60 days', NOW())
RETURNING *;

-- name: GetRefreshToken :one
SELECT * FROM refresh_tokens
WHERE token = $1;

-- name: RevokeRefreshToken :exec
UPDATE refresh_tokens
SET revoked_at = NOW(), updated_at = NOW()
WHERE token = $1;