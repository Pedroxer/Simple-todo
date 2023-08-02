-- name: AddListToUser :one
INSERT INTO "user_to_list" (user_id, list_id)
VALUES ($1, $2) RETURNING *;

-- name: DeleteListFromUser :exec
DELETE FROM "user_to_list" where list_id = $1 and user_id = $2;

-- name: ListAllUserLists :many
SELECT lists.* from "lists" lists, "user_to_list"
    where user_id = $1 and lists.id = list_id LIMIT 10;

-- name: G