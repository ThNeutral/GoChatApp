-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, access_token_updated_at, username, password, email, access_token)
VALUES ($1, $2, $3, $4, $5, $6, $7, encode(sha256(random()::text::bytea), 'hex'))
RETURNING *;

-- name: GetUserByAPIKey :one
SELECT * FROM users WHERE access_token = $1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: GetUserByEmailAndPassword :one
SELECT * FROM users WHERE (email = $1 and password = $2);

-- name: UpdateAccessTokenAndGetUser :one 
UPDATE users SET 
access_token = encode(sha256(random()::text::bytea), 'hex'), 
access_token_updated_at = $3 
WHERE (email = $1 and password = $2)
RETURNING *;

-- name: UpdateAccessTokenExpiryTimeAndGetUser :one 
UPDATE users SET 
access_token_updated_at = $2 
WHERE email = $1
RETURNING *;