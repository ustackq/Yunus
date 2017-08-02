package handler

import (
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"

	"github.com/ustack/Yunus/src/app/backend/pkg/setting"
)

// Captchaer return server capthaer
func Captchaer(c *gin.Context) {
	captcha.Server(setting.CaptchaStdWidth, setting.CaptchaStdHeight)
}
