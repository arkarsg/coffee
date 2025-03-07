package middleware

import (
	"coffeh/config"
	"coffeh/db"
	"coffeh/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	tgData "github.com/telegram-mini-apps/init-data-golang"
)

func IsTelegramUser(cfg *config.Env, db *db.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		var authRequest model.AuthRequest
		if err := c.ShouldBindJSON(&authRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid auth format"})
			c.Abort()
			return
		}

		err := tgData.Validate(authRequest.InitData, cfg.TelegramToken, config.TELEGRAM_INIT_DATA_EXPIRATION)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "Invalid auth data"})
			c.Abort()
			return
		}

		parsedData, err := tgData.Parse(authRequest.InitData)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "Failed to parse auth data"})
			c.Abort()
			return
		}

		authOutput := &model.AuthOutput{
			User:    parsedData.User,
			ChatID:  fmt.Sprint(parsedData.ChatInstance),
			Message: "Using Telegram Init Data",
		}

		c.Set(config.TG_AUTH_OUTPUT_KEY, authOutput)
		c.Next()
	}
}

func IsAdmin(db *db.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		authOutput, exists := c.Get(config.TG_AUTH_OUTPUT_KEY)
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "Unauthorized"})
			c.Abort()
			return
		}
		parsedAuthOutput, ok := authOutput.(*model.AuthOutput)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "Invalid user data"})
			c.Abort()
			return
		}

		_, err := db.FindUserByTelegramID(c, parsedAuthOutput.User.ID)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"success": false, "message": "Access denied, user is not an admin"})
			c.Abort()
			return
		}

		c.Next()
	}
}
