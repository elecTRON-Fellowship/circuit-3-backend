-- name: CreateUser :one
INSERT INTO users (first_name, last_name, user_name,  email, password, phone_no)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1
LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1
LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
LIMIT $1
OFFSET $2;

-- name: UpdateFirstName :exec
UPDATE users
SET first_name = $2
WHERE id = $1;

-- name: UpdateLastName :exec
UPDATE users
SET last_name = $2
WHERE id = $1;

-- name: UpdateUserName :exec
UPDATE users
SET user_name = $2
WHERE id = $1;

-- name: UpdateEmail :exec
UPDATE users
SET email = $2
WHERE id = $1;

-- name: UpdatePassword :exec
UPDATE users
SET password = $2
WHERE id = $1;

-- name: UpdatePhoneNo :exec
UPDATE users
SET phone_no = $2
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
