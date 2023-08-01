-- name: AddTaskToList :one
INSERT INTO "tasks_to_list" (task_id, list_id)
VALUES ($1, $2) RETURNING *;

-- name: ChangeListForTask :exec
UPDATE "tasks_to_list" SET list_id = $1 
where task_id = $2;

-- name: DeleteTaskFromList :exec
DELETE FROM "tasks_to_list" where task_id = $1;