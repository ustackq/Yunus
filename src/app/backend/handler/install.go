package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"

	"github.com/ustack/Yunus/src/app/backend/models"
	"github.com/ustack/Yunus/src/app/backend/pkg/utils"
)

// YunusInstall handler install work
func YunusInstall(c *gin.Context) {
	if _, err := utils.InstallCheck(); err != nil {
		//  Database settings
		models.DbCfg.Host = c.PostForm("dbhost")
		models.DbCfg.User = c.PostForm("user")
		models.DbCfg.Name = c.PostForm("name")
		models.DbCfg.Passwd = c.PostForm("passwd")
		// GeneratorconnStr return sql connection string
		connStr := models.GeneratorconnStr()
		// EnSureTableExist which ensure tables has beeen created
		err := models.EnSureTableExist(models.DbCfg.Type, connStr)
		if err != nil {
			glog.Fatal(err)
		}
	}
	c.AbortWithError(http.StateClosed, errors.New("Yunus has beed installed"))

}

// AdminConfig handler create admin account after install completed
func AdminConfig(c *gin.Context) {
	engine, err := models.GetSQLEngine()
	if err != nil {
		glog.Fatal(err)
	}
	// Create admin account
	admin := &models.User{}
	admin

}
