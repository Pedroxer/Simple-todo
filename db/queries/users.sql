-- name: CreateUser :one
INSERT INTO "users"
(username, password, email) 
values ($1,$2,$3) RETURNING *;

-- name: SelectUser :one
SELECT username, password, email 
FROM "users" where id = $1;

-- name: UpdateUser :exec
UPDATE "users" SET
    password = $2,
    email = $3
    where id = $1;

-- name: DeleteUser :exec
DELETE FROM "users" where id = $1; 
