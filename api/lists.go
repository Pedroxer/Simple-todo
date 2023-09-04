package api

import (
	"github.com/Pedroxer/Simple-todo/db/sqlc"
	"github.com/gin-gonic/gin"
	"net/http"
)

type createListRequest struct {
	Title  string `json:"title" binding:"required"`
	UserId int    `json:"user_id" binding:"required"`
}

type ListResponse struct {
	Title string `json:"title"`
}

func newListResponse(list sqlc.List) ListResponse {
	return ListResponse{Title: list.Title}
}
func (server *Server) createList(ctx *gin.Context) {
	var req createListRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	list, err := server.db.CreateList(ctx, req.Title)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	listToUserParams := sqlc.AddListToUserParams{
		UserID: int32(req.UserId),
		ListID: int32(list.ID),
	}
	_, err = server.db.AddListToUser(ctx, listToUserParams)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	resp := newListResponse(list)
	ctx.JSON(http.StatusOK, resp)
}

type getListRequest struct {
	ListId int `json:"list_id" binding:"required"`
}

func (server *Server) getList(ctx *gin.Context) {
	var req getListRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	list, err := server.db.GetList(ctx, int64(req.ListId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, newListResponse(list))
}

type changeListTitleRequest struct {
	List_id int    `json:"list_id"  binding:"required"`
	Title   string `json:"title" binding:"required"`
}

func (server *Server) changeListTitle(ctx *gin.Context) {
	var req changeListTitleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := sqlc.ChangeListNameParams{
		Title: req.Title,
		ID:    int64(req.List_id),
	}
	err := server.db.ChangeListName(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	list, err := server.db.GetList(ctx, int64(req.List_id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, newListResponse(list))
}

type deleteListRequest struct {
	List_id int `json:"list_id" binding:"required"`
	User_id int `json:"user_id" binding:"required"`
}

func (server *Server) deleteList(ctx *gin.Context) {
	var req deleteListRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	// 1. удалить все таски из листа
	tasks, err := server.db.ListAllTasks(ctx, int32(req.List_id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	for i := 0; i < len(tasks); i++ {
		arg := sqlc.DeleteTaskFromListParams{
			TaskID: int32(tasks[i].ID),
			ListID: int32(req.List_id),
		}
		err = server.db.DeleteTaskFromList(ctx, arg)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	}

	// 2. удалить лист из юзера
	arg := sqlc.DeleteListFromUserParams{
		ListID: int32(req.List_id),
		UserID: int32(req.User_id),
	}
	err = server.db.DeleteListFromUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	// 3. удалить лист
	err = server.db.DeleteList(ctx, int64(arg.ListID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, "Deleted")
}
