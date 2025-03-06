package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *APIServer) registerOrder() {
	orderRoutes := s.apiRoutes.Group("/order")
	orderRoutes.GET("/", s.getAllOrders)
	orderRoutes.GET("/:id", s.getOrdersByTelegramUser)
}

func (s *APIServer) getAllOrders(c *gin.Context) {
	allOrders, err := s.store.GetAllOrders(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &Response{
			Success: false,
			Message: "Something went wrong",
		})
		return
	}
	c.JSON(http.StatusOK, &Response{
		Success: true,
		Message: gin.H{
			"orders": allOrders,
		},
	})
}

func (s *APIServer) getOrdersByTelegramUser(c *gin.Context) {
	idParam := c.Param("id")
	telegramId, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, &Response{
			Success: false,
			Message: "Invalid user ID",
		})
		return
	}

	orders, err := s.store.FindAllOrdersByTelegramId(c, telegramId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &Response{
			Success: false,
			Message: "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, &Response{
		Success: true,
		Message: gin.H{
			"orders": orders,
		},
	})
}
