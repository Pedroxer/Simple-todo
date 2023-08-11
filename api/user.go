package api

import (
	"database/sql"
	"github.com/Pedroxer/Simple-todo/db/sqlc"
	"github.com/Pedroxer/Simple-todo/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateUserRequest struct for handling user input
type CreateUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
	Email    string `json:"email" binding:"required,email"`
}

type UserResponse struct {
	Username  string       `json:"username"`
	Password  string       `json:"password"`
	Email     string       `json:"email"`
	CreatedAt sql.NullTime `json:"CreatedAt"`
}

func newUserResp(user sqlc.User) UserResponse {
	return UserResponse{
		Username:  user.Username,
		Password:  user.Password,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
}
func (server *Server) CreateUser(ctx *gin.Context) {
	var req CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {

		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	req.Password = util.HashPassword(req.Password)

	arg := sqlc.CreateUserParams{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}
	user, err := server.db.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	resp := newUserResp(user)
	ctx.JSON(http.StatusOK, resp)
}

type getUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
}

func (server *Server) getUser(ctx *gin.Context) {
	var req getUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {

		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	user, err := server.db.GetUser(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, user)
}

//func (server *Server) LoginUser(ctx *gin.Context){
//	var
//}
