-- name: CreateUser :one
INSERT INTO "users"
(username, password, email) 
values ($1,$2,$3) RETURNING *;

-- name: GetUser :one
SELECT *
FROM "users" where username = $1;

-- name: ChangePassword :exec
UPDATE "users" SET
    password = $2
    where username = $1;

-- name: ChangeEmail :exec
UPDATE "users" SET
    email = $2
    where username = $1;

-- name: DeleteUser :exec
DELETE FROM "users" where username = $1; 
