package setting

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
)

func newCacheService() {

}
