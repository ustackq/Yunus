package setting

import (
	"time"

	"github.com/spf13/viper"
	"github.com/ustack/Yunus/src/app/backend/pkg/avatar"
)

var (
	// DisableRouterLog ...
	DisableRouterLog bool
	// EnableGzip ...
	EnableGzip bool
	// LogRootPath Log settings
	LogRootPath string
	// LogModes ...
	LogModes []string
	// EnableCORS ...
	EnableCORS bool

	// Cache settings

	// APP settings

	// RootPath ...
	RootPath string
	// APPSubURL url
	APPSubURL string
	// DisableGravatar disable gravatar
	DisableGravatar bool
	// CacheAdapter ...
	CacheAdapter string
	// CacheInterval ...
	CacheInterval int
	// CacheConn ...
	CacheConn string
	// CaptchaStdWidth ...
	CaptchaStdWidth int
	// CaptchaStdHeight ...
	CaptchaStdHeight int

	// Server settings

	// OfflineMode define
	OfflineMode bool
	// Cfg config
	Cfg *viper.Viper
	// SecretKey secret key
	SecretKey string

	// Picture settings

	// AvatarUploadPath Avatar Upload Path
	AvatarUploadPath string
	// LibravatarService avatar struct
	LibravatarService *avatar.Libravatar
	// GravatarSource avatar source
	GravatarSource string

	// Cron tasks
	Cron struct {
		UpdateMirror struct {
			Enabled    bool
			RunAtStart bool
			Schedule   string
		} `ini:"cron.update_mirrors"`
		RepoHealthCheck struct {
			Enabled    bool
			RunAtStart bool
			Schedule   string
			Timeout    time.Duration
			Args       []string `delim:" "`
		} `ini:"cron.repo_health_check"`
		CheckRepoStats struct {
			Enabled    bool
			RunAtStart bool
			Schedule   string
		} `ini:"cron.check_repo_stats"`
		RepoArchiveCleanup struct {
			Enabled    bool
			RunAtStart bool
			Schedule   string
			OlderThan  time.Duration
		} `ini:"cron.repo_archive_cleanup"`
	}
)

// Service struct
var Service struct {
	ActiveCodeLives                int
	ResetPwdCodeLives              int
	RegisterEmailConfirm           bool
	DisableRegistration            bool
	ShowRegistrationButton         bool
	RequireSignInView              bool
	EnableNotifyMail               bool
	EnableReverseProxyAuth         bool
	EnableReverseProxyAutoRegister bool
	EnableCaptcha                  bool
}

// init config
func init() {
	Cfg = viper.New()
}

func newCacheService() {

}
