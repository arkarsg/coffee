package api

import (
	"coffeh/config"
	"coffeh/db"
	"coffeh/model"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
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
	s.registerDrink()
}

func NewApiServer(store *db.Store, config *config.Env) (*APIServer, error) {
	server := &APIServer{
		config: config,
		store:  store,
	}

	server.initRoutes()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("valid_drinkvariant", model.ValidateDrinkVariant)
		v.RegisterValidation("valid_drinkcategory", model.ValidateDrinkCategory)
	}

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
