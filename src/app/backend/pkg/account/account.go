package account

// ProviderType define type
type ProviderType int

// Google init
const (
	Google ProviderType = iota
	Weixin
	Weibo

	Black int = iota
	White
)

// RuleAction define
type RuleAction map[string]interface{}

// Account define account struct
type Account struct {
	ID              int64
	UID             int64
	AccountType     ProviderType
	Name            string
	Locale          string
	Picture         string
	Gender          string
	Email           string
	Link            string
	CreatedTimeUnix int64
	AccessToken     int64
	RefreshToken    int64
	ExpiresTime     int64
}

// Provider define interface func
type Provider interface {
	GetAccessRule() RuleAction
	IsEnabled() bool
	RegisterAction()
}
