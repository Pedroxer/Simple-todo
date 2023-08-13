package api

import (
	"github.com/Pedroxer/Simple-todo/db/sqlc"
	"github.com/gin-gonic/gin"
	"net/http"
)

type createListRequest struct {
	Title  string `json:"title"`
	UserId int    `json:"user_id"`
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
