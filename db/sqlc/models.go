// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package sqlc

import (
	"database/sql"
	"time"
)

type List struct {
	ID        int32        `json:"id"`
	Title     string       `json:"title"`
	CreatedAt sql.NullTime `json:"created_at"`
}

type Task struct {
	ID          int64          `json:"id"`
	Name        sql.NullString `json:"name"`
	Description sql.NullString `json:"description"`
	// 1 for important, 0 for regular task
	Important sql.NullInt32 `json:"important"`
	// 1 for done, 0 for in progress
	Done      sql.NullInt32 `json:"done"`
	Deadline  time.Time     `json:"deadline"`
	CreatedAt sql.NullTime  `json:"created_at"`
}

type TasksToList struct {
	TaskID int32 `json:"task_id"`
	ListID int32 `json:"list_id"`
}

type User struct {
	ID        int64        `json:"id"`
	Username  string       `json:"username"`
	Password  string       `json:"password"`
	Email     string       `json:"email"`
	CreatedAt sql.NullTime `json:"created_at"`
}

type UserToList struct {
	ListID int32 `json:"list_id"`
	UserID int32 `json:"user_id"`
}
