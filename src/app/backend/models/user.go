package models

import (
	"strings"
	"time"

	"github.com/ustack/Yunus/src/app/backend/errors"
)

var (
	reservedUsernames    = []string{"assets", "css", "img", "js", "less", "plugins", "debug", "raw", "install", "api", "avatar", "user", "org", "help", "stars", "issues", "pulls", "commits", "repo", "template", "admin", "new", ".", ".."}
	reservedUserPatterns = []string{"*.keys"}
)

// User universal definition
type User struct {
	UID                  int64
	UserName             string
	Email                string
	Mobile               string
	Salt                 string
	AvatarFile           string
	Sex                  bool
	Birthday             string
	Province             string
	City                 string
	JobID                int16
	RegTime              int64
	RegIP                string
	LastLogin            time.Time
	LastIP               string
	OnlineTime           time.Time
	LastActive           time.Time
	NotificationUnread   int32
	InboxUnread          int32
	InboxRecv            int32
	FansCount            int64
	FriendCount          int64
	BeInvitedCount       int32
	ArticleCount         int32
	QuestionCount        int32
	AnswerCount          int16
	TopicFocusCount      int16
	InvitedCount         int16
	GroupID              int16
	ReputationGroup      int16
	Forbidden            bool
	ValidEmail           bool
	IsFirstLogin         bool
	AgreeCount           int32
	ThanksCount          int32
	ViewsCount           int32
	Reputation           int32
	ReputationUpdateTime time.Time
	WeiboVisit           bool
	Integral             int32
	DraftCount           int32
	CommonEmail          string
	URLToken             string
	URLTokenUpdate       string
	Verified             string
	DefaultTimeZone      string
	EmailSetting         string
	WeixinSetting        string
	RecentTopics         string
}

// UpdateRegTime update registry time
func (u *User) UpdateRegTime() {
	u.RegTime = time.Now().Unix()
}

// ValidateUserName whether the username is exist
func ValidateUserName(name string) error {
	name = strings.TrimSpace(strings.ToLower(name))
	if len(name) == 0 {
		return errors.EmptyName{}
	}
	if strings.Contains(strings.Join(reservedUsernames, ","), name) {
		return errors.NameReserved{}
	}
	return nil
}

// ExistUserName validate username
func ExistUserName(name string) (bool, error) {
	var nameResult string
	if len(name) == 0 {
		return false, nil
	}
	 rows, err := x.Query("select * from user", name)
	 if err != nil {
		 return false, err
	 }
	 defer rows.Close()
	 for rows.Next() {
		 err = rows.Scan(&nameResult)
		 if err != nil {
			 return false, err
		 }
	 }
	 if strings.Compare(name, nameResult) {
		 return true, nil
	 }
	 return false, err
}

// CreateUser create a new user
func CreateUser(u *User) error{
	if err = ValidateUserName(u.UserName); err != nil {
		return err
	}
	if isExist, err := ExistUserName(u.UserName); err != nil {
		return err
	} else if isExist {
		return errors.ErrUserAlreadyExist{}
	}
	u.Email = strings.ToLower(u.Email)
	isExist, err = IsEmailUsed(u.Email)
	if err != nil {
		return err
	}else if isExist {
		return errors.ErrEmailAlreadyUsed{Email:u.Email}
	}
	// need refacotr
	_, err := x.Exec("insert ", u)
	if err != nil {
		return err
	}
	return nil

}
