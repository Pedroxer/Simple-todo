// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: tasksFromList.sql

package sqlc

import (
	"context"
)

const addTaskToList = `-- name: AddTaskToList :one
INSERT INTO "tasks_to_list" (task_id, list_id)
VALUES ($1, $2) RETURNING task_id, list_id
`

type AddTaskToListParams struct {
	TaskID int32 `json:"task_id"`
	ListID int32 `json:"list_id"`
}

func (q *Queries) AddTaskToList(ctx context.Context, arg AddTaskToListParams) (TasksToList, error) {
	row := q.db.QueryRowContext(ctx, addTaskToList, arg.TaskID, arg.ListID)
	var i TasksToList
	err := row.Scan(&i.TaskID, &i.ListID)
	return i, err
}

const changeListForTask = `-- name: ChangeListForTask :exec
UPDATE "tasks_to_list" SET list_id = $1 
where task_id = $2
`

type ChangeListForTaskParams struct {
	ListID int32 `json:"list_id"`
	TaskID int32 `json:"task_id"`
}

func (q *Queries) ChangeListForTask(ctx context.Context, arg ChangeListForTaskParams) error {
	_, err := q.db.ExecContext(ctx, changeListForTask, arg.ListID, arg.TaskID)
	return err
}

const deleteTaskFromList = `-- name: DeleteTaskFromList :exec
DELETE FROM "tasks_to_list" where task_id = $1 and list_id = $2
`

type DeleteTaskFromListParams struct {
	TaskID int32 `json:"task_id"`
	ListID int32 `json:"list_id"`
}

func (q *Queries) DeleteTaskFromList(ctx context.Context, arg DeleteTaskFromListParams) error {
	_, err := q.db.ExecContext(ctx, deleteTaskFromList, arg.TaskID, arg.ListID)
	return err
}

const listAllTasks = `-- name: ListAllTasks :many
SELECT tasks.id, tasks.name, tasks.description, tasks.important, tasks.done, tasks.deadline, tasks.created_at from "tasks" tasks, "tasks_to_list"
                 where list_id = $1 and tasks.id = task_id LIMIT 10
`

func (q *Queries) ListAllTasks(ctx context.Context, listID int32) ([]Task, error) {
	rows, err := q.db.QueryContext(ctx, listAllTasks, listID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Task
	for rows.Next() {
		var i Task
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Important,
			&i.Done,
			&i.Deadline,
			&i.CreatedAt,
		); err != nil {
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
