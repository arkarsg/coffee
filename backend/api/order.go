package api

import (
	"coffeh/config"
	"coffeh/model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (s *APIServer) registerOrder() {
	orderRoutes := s.apiRoutes.Group("/order")

	orderRoutes.GET("/", s.getAllOrders)
	orderRoutes.GET("/user", s.getOrdersByTelegramUser)
	orderRoutes.POST("/", s.createOrder)
	orderRoutes.PATCH("/fulfill/:id", s.fulfillOrder)
	orderRoutes.PATCH("/cancel/:id", s.cancelOrder)
	orderRoutes.DELETE("/:id", s.deleteOrder)
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
	idParam := c.Query("id")
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
			Message: err.Error(),
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

func (s *APIServer) createOrder(c *gin.Context) {
	var createOrderRequest model.CreateOrderDTO
	if err := c.ShouldBindJSON(&createOrderRequest); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if len(createOrderRequest.Items) > config.MAX_ORDER_ITEMS {
		c.JSON(http.StatusBadRequest, &Response{
			Success: false,
			Message: "Item limit exceeded",
		})
		return
	}

	if !model.IsValidCollectionDate(createOrderRequest.CollectFrom) {
		c.JSON(http.StatusBadRequest, &Response{
			Success: false,
			Message: "Collection time for pre-order is in the past",
		})
		return
	}

	createdId, err := s.store.CreateOrder(c, createOrderRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &Response{
		Success: true,
		Message: fmt.Sprintf("Created %s", createdId),
	})
}

type orderRequest struct {
	OrderID string `uri:"id"`
}

func (s *APIServer) fulfillOrder(c *gin.Context) {
	var req orderRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	orderId, err := bson.ObjectIDFromHex(req.OrderID)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err = s.store.FulfillOrder(c, orderId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &Response{
		Success: true,
		Message: fmt.Sprintf("Fulfilled order successfully"),
	})
}

func (s *APIServer) cancelOrder(c *gin.Context) {
	var req orderRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	orderId, err := bson.ObjectIDFromHex(req.OrderID)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err = s.store.CancelOrder(c, orderId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &Response{
		Success: true,
		Message: fmt.Sprintf("Fulfilled order successfully"),
	})
}

func (s *APIServer) deleteOrder(c *gin.Context) {
	var req orderRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	orderId, err := bson.ObjectIDFromHex(req.OrderID)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err = s.store.DeleteOrder(c, orderId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &Response{
		Success: true,
		Message: fmt.Sprintf("Fulfilled order successfully"),
	})
}
