package cmd

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/ustack/Yunus/src/app/backend/pkg/setting"
)

func newGin() *gin.Engine {
	g := gin.New()
	g.Use(gin.Recovery())
	if !setting.DisableRouterLog {
		g.Use(gin.Logger())
	}
	if !setting.EnableGzip {
		g.Use(gzip.Gzip(gzip.DefaultCompression))
	}
	if !setting.EnableCORS {
		g.Use(cors.Default())
	}

	return g
}
