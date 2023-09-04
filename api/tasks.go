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

func createTaskResponse(task sqlc.Task) taskResponse {
	return taskResponse{
		Name:        task.Name,
		Description: task.Description.String,
		Important:   int(task.Important.Int32),
		Done:        int(task.Done.Int32),
		Deadline:    task.Deadline,
	}
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
		return
	}

	ctx.JSON(http.StatusOK, createTaskResponse(task))
}

type getTaskRequest struct {
	Name string `json:"name" binding:"required"`
}

func (serv *Server) getTask(ctx *gin.Context) {
	var req getTaskRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	task, err := serv.db.GetTask(ctx, req.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, createTaskResponse(task))
}

type updateTaskNameRequest struct {
	Id   int    `json:"task_id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

func (serv *Server) updateName(ctx *gin.Context) {
	var req updateTaskNameRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := sqlc.ChangeTaskNameParams{
		ID:   int64(req.Id),
		Name: req.Name,
	}
	err := serv.db.ChangeTaskName(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	task, err := serv.db.GetTask(ctx, req.Name)
	ctx.JSON(http.StatusOK, createTaskResponse(task))
}

type updateDescriptionRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func (serv *Server) updateDescription(ctx *gin.Context) {
	var req updateDescriptionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := sqlc.ChangeDescriptionParams{
		Description: sql.NullString{
			String: req.Description,
			Valid:  true,
		},
		Name: req.Name,
	}
	err := serv.db.ChangeDescription(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	task, err := serv.db.GetTask(ctx, req.Name)
	ctx.JSON(http.StatusOK, createTaskResponse(task))
}

type updateOrderRequest struct {
	Name  string `json:"name" binding:"required"`
	Order int    `json:"order" binding:"required"`
}

func (serv *Server) updateOrder(ctx *gin.Context) {
	var req updateOrderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := sqlc.ChangeTaskOrderParams{
		Important: sql.NullInt32{
			Int32: int32(req.Order),
			Valid: true,
		},
		Name: "",
	}
	err := serv.db.ChangeTaskOrder(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	task, err := serv.db.GetTask(ctx, req.Name)
	ctx.JSON(http.StatusOK, createTaskResponse(task))
}

type updateDoneRequest struct {
	Name string `json:"name" binding:"required"`
	Done int    `json:"done" binding:"required"`
}

func (serv *Server) updateDone(ctx *gin.Context) {
	var req updateDoneRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := sqlc.ChangeTaskDoneParams{
		Done: sql.NullInt32{Int32: int32(req.Done)},
		Name: req.Name,
	}
	err := serv.db.ChangeTaskDone(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	task, err := serv.db.GetTask(ctx, req.Name)
	ctx.JSON(http.StatusOK, createTaskResponse(task))
}

type updateDeadLineRequest struct {
	Name     string    `json:"name" binding:"required"`
	DeadLine time.Time `json:"description" binding:"required"`
}

func (serv *Server) updateDeadline(ctx *gin.Context) {
	var req updateDeadLineRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := sqlc.ChangeTaskDeadlineParams{
		Deadline: req.DeadLine,
		Name:     req.Name,
	}

	err := serv.db.ChangeTaskDeadline(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	task, err := serv.db.GetTask(ctx, req.Name)
	ctx.JSON(http.StatusOK, createTaskResponse(task))
}

type deleteTaskRequest struct {
	Name   string `json:"name" binding:"required"`
	ListID int    `json:"list_id" binding:"required"`
}

func (serv *Server) deleteTask(ctx *gin.Context) {
	var req deleteTaskRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	task, err := serv.db.GetTask(ctx, req.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := sqlc.DeleteTaskFromListParams{
		TaskID: int32(task.ID),
		ListID: int32(req.ListID),
	}

	err = serv.db.DeleteTask(ctx, req.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = serv.db.DeleteTaskFromList(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, "Deleted")
}
