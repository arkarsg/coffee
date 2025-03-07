package api

import (
	"coffeh/config"
	"coffeh/db"
	"coffeh/model"
	"fmt"

	"github.com/gin-contrib/cors"
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
	// apiRoutes.Use(middleware.IsTelegramUser(s.config, s.store))

	s.router = r
	s.apiRoutes = apiRoutes

	// s.setupCORS()

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

func (s *APIServer) setupCORS() {
	corsConfig := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization", "X-Requested-With", "Accept", "Origin"},
		AllowCredentials: true,
	}

	s.router.Use(cors.New(corsConfig))
}

func errorResponse(err error) *Response {
	return &Response{
		Success: false,
		Message: err.Error(),
	}
}
