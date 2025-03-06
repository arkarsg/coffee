package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *APIServer) registerPing() {
	s.router.GET("/", home)
	s.router.GET("/ping", pong)
}

func home(c *gin.Context) {
	c.JSON(http.StatusOK, &Response{
		Success: true,
		Message: "Hello from Coffe(EH) API Server",
	})
}

func pong(c *gin.Context) {
	c.JSON(http.StatusOK, &Response{
		Success: true,
		Message: "Pong",
	})
}
