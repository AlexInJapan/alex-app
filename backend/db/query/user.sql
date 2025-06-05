-- name: CreateUser :one
INSERT INTO users (
    name, email
) VALUES (
    $1, $2
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;