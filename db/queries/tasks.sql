-- name: CreateTask :one
INSERT INTO "tasks" 
(name, description, important, done, deadline)
VALUES ($1,$2,$3,$4,$5) RETURNING *;

-- name: GetTask :one
SELECT * FROM "tasks" where name = $1;

-- name: ChangeTaskName :exec
UPDATE "tasks" SET name = $1 where name = $2;

-- name: ChangeDescription :exec
UPDATE "tasks" SET 
description = $1 
where name = $2;

-- name: ChangeTaskOrder :exec
UPDATE "tasks" SET 
important = $1 
where name = $2;

-- name: ChangeTaskDone :exec
UPDATE "tasks" SET 
done = $1 
where name = $2;

-- name: ChangeTaskDeadline :exec
UPDATE "tasks" SET 
deadline = $1 
where name = $2;

-- name: DeleteTask :exec
DELETE FROM "tasks" where name = $1;