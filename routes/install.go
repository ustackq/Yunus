package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ustack/Yunus/src/app/backend/handler"
)

// RegisterInstall ...
func RegisterInstall(r *gin.Engine) {
	r.POST("/install", handler.YunusInstall)
}