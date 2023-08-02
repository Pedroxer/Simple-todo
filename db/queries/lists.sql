-- name: CreateList :one
INSERT INTO "lists" (title)
values ($1) RETURNING *;

-- name: GetList :one
SELECT * FROM "lists" where title = $1;

-- name: ChangeListName :exec
UPDATE "lists" SET title = $1 where id = $2;



-- name: DeleteList :exec
DELETE FROM "lists" where title = $1;
