// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: lists.sql

package sqlc

import (
	"context"
)

const changeListName = `-- name: ChangeListName :exec
UPDATE "lists" SET title = $1 where id = $2
`

type ChangeListNameParams struct {
	Title string `json:"title"`
	ID    int64  `json:"id"`
}

func (q *Queries) ChangeListName(ctx context.Context, arg ChangeListNameParams) error {
	_, err := q.db.ExecContext(ctx, changeListName, arg.Title, arg.ID)
	return err
}

const createList = `-- name: CreateList :one
INSERT INTO "lists" (title)
values ($1) RETURNING id, title, created_at
`

func (q *Queries) CreateList(ctx context.Context, title string) (List, error) {
	row := q.db.QueryRowContext(ctx, createList, title)
	var i List
	err := row.Scan(&i.ID, &i.Title, &i.CreatedAt)
	return i, err
}

const deleteList = `-- name: DeleteList :exec
DELETE FROM "lists" where id = $1
`

func (q *Queries) DeleteList(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteList, id)
	return err
}

const getList = `-- name: GetList :one
SELECT id, title, created_at FROM "lists" where id = $1
`

func (q *Queries) GetList(ctx context.Context, id int64) (List, error) {
	row := q.db.QueryRowContext(ctx, getList, id)
	var i List
	err := row.Scan(&i.ID, &i.Title, &i.CreatedAt)
	return i, err
}
