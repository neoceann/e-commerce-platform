-- name: CreateUser :one
INSERT INTO users (email, first_name, last_name, phone, user_role, password_hash)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, email, first_name, last_name, phone, user_role, password_hash, created_at, updated_at;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1;

-- name: UpdatePassword :exec
UPDATE users
SET password_hash = $2, updated_at = NOW()
WHERE id = $1;

-- name: UserExists :one
SELECT EXISTS(SELECT 1 FROM users WHERE email = $1);