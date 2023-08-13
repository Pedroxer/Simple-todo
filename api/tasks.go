package api

import (
	"database/sql"
	"github.com/Pedroxer/Simple-todo/db/sqlc"
	"github.com/Pedroxer/Simple-todo/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type createTaskRequest struct {
	Name        string        `json:"name" binding:"required"`
	Description string        `json:"description"` // todo: instead string use some text object
	Important   int           `json:"important"`
	Done        int           `json:"done"`
	ListId      int           `json:"list_id"`
	Deadline    util.Duration `json:"deadline"`
}

type taskResponse struct {
	Name        string    `json:"name"`
	Description string    `json:"description"` // todo: instead string use some text object
	Important   int       `json:"important"`
	Done        int       `json:"done"`
	Deadline    time.Time `json:"deadline"`
}

func (server *Server) createTask(ctx *gin.Context) {
	var req createTaskRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	taskParam := sqlc.CreateTaskParams{
		Name: req.Name,
		Description: sql.NullString{
			String: req.Description,
			Valid:  true,
		},
		Important: sql.NullInt32{
			Int32: int32(req.Important),
			Valid: true,
		},
		Done: sql.NullInt32{
			Int32: int32(req.Done),
			Valid: true,
		},
		Deadline: time.Now().Add(req.Deadline.Duration),
	}
	task, err := server.db.CreateTask(ctx, taskParam)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	taskTolistParams := sqlc.AddTaskToListParams{
		TaskID: int32(task.ID),
		ListID: int32(req.ListId),
	}
	_, err = server.db.AddTaskToList(ctx, taskTolistParams)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, task)
}
