package api

import (
	"github.com/Pedroxer/Simple-todo/db/sqlc"
	"github.com/Pedroxer/Simple-todo/util"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	config util.Config
	db     *sqlc.Queries
}

func NewServer(config util.Config, db *sqlc.Queries) *Server {
	server := &Server{
		config: config,
		db:     db,
	}
	server.SetupRoutes()
	return server
}

func (server *Server) SetupRoutes() {
	router := gin.Default()

	router.POST("/user", server.CreateUser)
	server.router = router
}

func (server *Server) Start(config util.Config) error {
	return server.router.Run(config.ServerAddress)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
