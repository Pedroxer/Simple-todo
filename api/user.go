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
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func newUserResp(user sqlc.User) UserResponse {
	return UserResponse{
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
	}
}
func (server *Server) createUser(ctx *gin.Context) {
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
	resp := newUserResp(user)
	ctx.JSON(http.StatusOK, resp)
}

type loginUserReq struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

type loginUserResp struct {
	AccessToken string       `json:"access_token"`
	User        UserResponse `json:"user"`
}

func (server *Server) loginUser(ctx *gin.Context) {
	var req loginUserReq
	if err := ctx.ShouldBindJSON(&req); err != nil {

		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	} // 1. получаем пользователя
	user, err := server.db.GetUser(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	// 2. Проверяем пароль
	err = util.CheckPassword(req.Password, user.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	// 3. Создаём токен
	accessToken, err := server.tokenMaker.CreateToken(user.Username, server.config.TokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	resp := loginUserResp{
		AccessToken: accessToken,
		User:        newUserResp(user),
	}
	ctx.JSON(http.StatusOK, resp)

}
