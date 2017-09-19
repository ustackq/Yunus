package handler

import (
	"errors"
	"net/http"
	"time"

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
	c.AbortWithError(http.StatusBadRequest, errors.New("Yunus has beed installed"))

}

// AdminConfig handler create admin account after install completed
func AdminConfig(c *gin.Context) {

	// Create admin account
	regTime := time.Now()
	admin := &models.User{}
	admin.UserName = c.PostForm("username")
	admin.Passwd = c.PostForm("passwd")
	admin.Email = c.PostForm("email")
	admin.GroupID = 1
	admin.Salt = utils.GenerateSalt()
	admin.Passwd = utils.GeneratePasswd(admin.Passwd, admin.Salt)
	admin.ReputationGroup = 5
	admin.ValidEmail = 1
	admin.IsFirstLogin = 1
	admin.RegTime = regTime
	admin.RegIP = c.Request.Host
	admin.LastLogin = regTime
	admin.LastIP = c.Request.Host
	admin.LastActive = regTime
	admin.InvitationAvailable = 10
	admin.Integral = 2000
	// ... and other element
	models.CreateUser(admin)
}
