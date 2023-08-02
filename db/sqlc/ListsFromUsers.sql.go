// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: ListsFromUsers.sql

package sqlc

import (
	"context"
)

const addListToUser = `-- name: AddListToUser :one
INSERT INTO "user_to_list" (user_id, list_id)
VALUES ($1, $2) RETURNING list_id, user_id
`

type AddListToUserParams struct {
	UserID int32 `json:"user_id"`
	ListID int32 `json:"list_id"`
}

func (q *Queries) AddListToUser(ctx context.Context, arg AddListToUserParams) (UserToList, error) {
	row := q.db.QueryRowContext(ctx, addListToUser, arg.UserID, arg.ListID)
	var i UserToList
	err := row.Scan(&i.ListID, &i.UserID)
	return i, err
}

const deleteListFromUser = `-- name: DeleteListFromUser :exec
DELETE FROM "user_to_list" where list_id = $1 and user_id = $2
`

type DeleteListFromUserParams struct {
	ListID int32 `json:"list_id"`
	UserID int32 `json:"user_id"`
}

func (q *Queries) DeleteListFromUser(ctx context.Context, arg DeleteListFromUserParams) error {
	_, err := q.db.ExecContext(ctx, deleteListFromUser, arg.ListID, arg.UserID)
	return err
}

const listAllUserLists = `-- name: ListAllUserLists :many
SELECT lists.id, lists.title, lists.created_at from "lists" lists, "user_to_list"
    where user_id = $1 and lists.id = list_id LIMIT 10
`

func (q *Queries) ListAllUserLists(ctx context.Context, userID int32) ([]List, error) {
	rows, err := q.db.QueryContext(ctx, listAllUserLists, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []List
	for rows.Next() {
		var i List
		if err := rows.Scan(&i.ID, &i.Title, &i.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
