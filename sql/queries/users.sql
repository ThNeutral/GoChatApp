-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, username, password, email, access_token)
VALUES ($1, $2, $3, $4, $5, $6, encode(sha256(random()::text::bytea), 'hex'))
RETURNING *;

-- name: GetUserByAPIKey :one
SELECT * FROM users WHERE access_token = $1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: GetUserByEmailAndPassword :one
SELECT * FROM users WHERE (email = $1 and password = $2);