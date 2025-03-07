package api

import (
	"coffeh/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (s *APIServer) registerDrink() {
	drinkRoutes := s.apiRoutes.Group("/drink")

	drinkRoutes.GET("/", s.getAllDrinks)
	drinkRoutes.GET("/:slug", s.getDrink)
	drinkRoutes.POST("/", s.createDrink)
	drinkRoutes.PATCH("/:slug", s.updateDrink)
	drinkRoutes.DELETE("/:slug", s.deleteDrink)
}

func (s *APIServer) getAllDrinks(c *gin.Context) {
	drinks, err := s.store.GetAllDrinks(c)
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
			"drinks": drinks,
		},
	})
}

func (s *APIServer) getDrink(c *gin.Context) {
	drinkSlug := c.Param("slug")
	if !slug.IsSlug(drinkSlug) {
		c.JSON(http.StatusBadRequest, &Response{
			Success: false,
			Message: "Invalid drink slug",
		})
		return
	}

	drink, err := s.store.GetDrink(c, drinkSlug)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, &Response{
				Success: false,
				Message: "No such drink",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, &Response{
			Success: false,
			Message: "Something went wrong",
		})
		return
	}
	c.JSON(http.StatusOK, &Response{
		Success: true,
		Message: drink,
	})
}

func (s *APIServer) createDrink(c *gin.Context) {
	var createDrinkRequest model.CreateDrinkDTO
	if err := c.ShouldBindJSON(&createDrinkRequest); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := s.store.CreateDrink(c, createDrinkRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &Response{
		Success: true,
		Message: fmt.Sprintf("Created drink: %s", createDrinkRequest.Name),
	})
}

func (s *APIServer) updateDrink(c *gin.Context) {
	drinkSlug := c.Param("slug")

	if !slug.IsSlug(drinkSlug) {
		c.JSON(http.StatusBadRequest, &Response{
			Success: false,
			Message: "Invalid drink slug",
		})
		return
	}

	var updateDrinkRequest model.UpdateDrinkDTO
	if err := c.ShouldBindJSON(&updateDrinkRequest); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := s.store.UpdateDrinkBySlug(c, drinkSlug, updateDrinkRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &Response{
		Success: true,
		Message: fmt.Sprintf("Updated drink: %s", drinkSlug),
	})
}

func (s *APIServer) deleteDrink(c *gin.Context) {
	drinkSlug := c.Param("slug")
	if !slug.IsSlug(drinkSlug) {
		c.JSON(http.StatusBadRequest, &Response{
			Success: false,
			Message: "Invalid drink slug",
		})
		return
	}

	err := s.store.DeleteDrink(c, drinkSlug)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &Response{
		Success: true,
		Message: fmt.Sprintf("Deleted drink: %s", drinkSlug),
	})
}
