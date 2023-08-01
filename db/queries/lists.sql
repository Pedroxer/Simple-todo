-- name: CreateList :one
INSERT INTO "lists" (title)
values ($1);

-- name: GetList :one
SELECT * FROM "lists" where title = $1;

-- name: ChangeListName :exec
UPDATE "lists" SET title = $1 where title = $2;

-- name: ChangeDescription :exec
UPDATE "tasks" SET 
description = $1 
where name = $2;

-- name: DeleteList :exec
DELETE FROM "lists" where title = $1;
