-- name: CreateUser :one
INSERT INTO users(username,display_name,bio,avatar_url)
VALUES ($1, $2, $3, $4)
RETURNING *;