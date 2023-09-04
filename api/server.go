package api

import (
	"fmt"
	"github.com/Pedroxer/Simple-todo/db/sqlc"
	"github.com/Pedroxer/Simple-todo/token"
	"github.com/Pedroxer/Simple-todo/util"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router     *gin.Engine
	config     util.Config
	db         *sqlc.Queries
	tokenMaker token.Maker // just interface
}

func NewServer(config util.Config, db *sqlc.Queries) (*Server, error) {
	Jwtm, err := token.NewJwtToken(config.TokenKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		db:         db,
		tokenMaker: Jwtm,
	}
	server.SetupRoutes()
	return server, nil
}

func (server *Server) SetupRoutes() {
	router := gin.Default()
	authRoutes := router.Group("/").Use(authMiddlware(server.tokenMaker))

	router.POST("/user", server.createUser)
	router.GET("/user", server.getUser)
	router.POST("/user/login", server.loginUser)

	authRoutes.POST("/task", server.createTask)
	authRoutes.GET("/task", server.getTask)
	authRoutes.POST("/task/name", server.updateName)
	authRoutes.POST("/task/description", server.updateDescription)
	authRoutes.POST("/task/order", server.updateOrder)
	authRoutes.POST("task/done", server.updateDone)
	authRoutes.POST("/task/deadline", server.updateDeadline)
	authRoutes.DELETE("/task", server.deleteTask)

	authRoutes.POST("/list", server.createList)
	authRoutes.GET("/list", server.getList)
	authRoutes.POST("/list/title", server.changeListTitle)
	authRoutes.DELETE("/list", server.deleteList)

	server.router = router
}

func (server *Server) Start(config util.Config) error {
	return server.router.Run(config.ServerAddress)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
