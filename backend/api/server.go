package api

import (
	"coffeh/config"
	"coffeh/db"
	"fmt"

	"github.com/gin-gonic/gin"
)

type APIServer struct {
	config    *config.Env
	store     *db.Store
	router    *gin.Engine
	apiRoutes *gin.RouterGroup
}

type Response struct {
	Success bool `json:"success"`
	Message any  `json:"message"`
}

func (s *APIServer) initRoutes() {
	r := gin.Default()
	apiRoutes := r.Group("/api")

	s.router = r
	s.apiRoutes = apiRoutes

	// Register API routes here
	s.registerPing()
	s.registerOrder()
}

func NewApiServer(store *db.Store, config *config.Env) (*APIServer, error) {
	server := &APIServer{
		config: config,
		store:  store,
	}

	server.initRoutes()

	return server, nil
}

func (s *APIServer) Start() error {
	port := fmt.Sprintf(":%s", s.config.Port)
	return s.router.Run(port)
}

func errorResponse(err error) *Response {
	return &Response{
		Success: false,
		Message: err.Error(),
	}
}
