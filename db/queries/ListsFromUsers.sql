-- name: AddListToUser :one
INSERT INTO "user_to_list" (user_id, list_id)
VALUES ($1, $2) RETURNING *;

-- name: DeleteListFromUser :exec
DELETE FROM "user_to_list" where list_id = $1;