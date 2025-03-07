package model

import tgData "github.com/telegram-mini-apps/init-data-golang"

type AuthRequest struct {
	InitData string `json:"initData" binding:"required"`
	IsMocked *bool  `json:"isMocked" binding:"required"`
}

type AuthOutput struct {
	User    tgData.User `json:"user"`
	ChatID  string      `json:"chat_id"`
	Message string      `json:"message"`
}
