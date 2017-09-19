package models

import (
	"os"
	"time"
	"fmt"
	"bytes"
	"strings"
	"image"
	"image/png"
	"path/filepath"
	"encoding/hex"
	"crypto/sha256"
	"crypto/subtle"
	"unicode/utf8"

	"github.com/golang/glog"
	"golang.org/x/crypto/pbkdf2"
	"github.com/nfnt/resize"
	"github.com/go-xorm/xorm"
	"github.com/ustack/Yunus/src/app/backend/pkg/setting"
	"github.com/ustack/Yunus/src/app/backend/pkg/account"
	"github.com/ustack/Yunus/src/app/backend/pkg/utils"
	"github.com/ustack/Yunus/src/app/backend/pkg/utils/convert"
	"github.com/ustack/Yunus/src/app/backend/pkg/avatar"
	"github.com/ustack/Yunus/src/app/backend/errors"
	)

var (
	reservedUsernames    = []string{"assets", "css", "img", "js", "less", "plugins", "debug", "raw", "install", "api", "avatar", "user", "org", "help", "stars", "issues", "pulls", "commits", "repo", "template", "admin", "new", ".", ".."}
	reservedUserPatterns = []string{"*.keys"}
)

type ActiveData struct {
	ID int64
	UID int64
	ExpireTimeUnix int64
	ActiveCode int64
	ActiveTypeCode int
	AddTimeUnix int64
	ADDIP string
	ActiveTimeUnix int64
	ActiveIP string
}
// User universal definition
type User struct {
	ID                   int64
	UserName             string `xorm:"UNIQUE NOT NULL"`
	Email                string `xorm:"UNIQUE NOT NULL"`
	Mobile               string `xorm:"UNIQUE NOT NULL"`
	Passwd               string
	Salt                 string `xorm:"VARCHAR(10)"`
	AvatarFile           string `xorm:"VARCHAR(2048) NOT NULL"`
	Sex                  bool
	Birthday             string
	Province             string
	City                 string
	JobID                int64
	RegTime              time.Time
	RegIP                string
	LastLogin            time.Time
	LastIP               string
	OnlineTime           time.Time
	LastActive           time.Time
	Integral             int32
	NotificationUnread   int32
	InboxUnread          int32
	InboxRecv            int32
	FansCount            int32
	FriendCount          int32
	BeInvitedCount       int32
	ArticleCount         int32
	QuestionCount        int32
	AnswerCount          int32
	TopicFocusCount      int32
	InvitedCount         int32
	GroupID              int32
	ReputationGroup      int32
	Forbidden            bool
	ValidEmail           bool
	IsFirstLogin         bool
	AgreeCount           int32
	ThanksCount          int32
	ViewsCount           int32
	Reputation           int32
	ReputationUpdateTime time.Time
	WeiboVisit           bool
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

type UsersAttribute struct {
	ID int64
	UID int64
	Introduction string
	Signature string
	QQ int
	HomePage string
}

type UsersGroup struct {
	ID int64
	Type int
	Custom bool
	Name string
	ReputationLower int
	ReputationHiger int
	ReputationFactor int
}

type UsersNotificationSetting struct {
	ID int64
	UID int64
	Data string
}

type UserOnline struct {
	UID int64
	LastActive int64
	ActiveIP string
	ActiveURL string
	UserAgent string
}

type UsersQQ struct {
	ID int64
	UID int64
	NickName string
	OpenID int64
	Gender string
	CreatedTimeUnix int64
	AccessToken int64
	RefreshToken int64
	ExpiresTime int64
	FigureURL string
}

type UsersSina struct {
	ID int64
	UID int64
	Name string
	Location string
	Description string
	URL string
	ProfileImageURL string
	Gender string
	CreatedTimeUnix int64
	AccessToken int64
	LastMessageID int64
}

type UsersWeixin struct {
	ID int64
	UID int64
	OpenID int64
	ExpiresID int
	AccessToken string
	RefreshToken string
	Scope string
	HeadImgURL string
	NickName string
	Sex bool
	Province string
	City string
	Country string
	CreatedTimeUnix int64
	Latitude float64
	LongLatitude float64
	LocationUpdate int
}

type UsersYunus struct {
	ID int64
	UID int64
	YUID int64
	UserName string
	Email string
}

type UsersGoogle struct {
	ID int64
	UID int64
	Name string
	Locale string
	Picture string
	Gender string
	Email string
	Link string
	CreatedTimeUnix int64
	AccessToken int64
	RefreshToken int64
	ExpiresTime int64
}

type UsersFacebook struct {
	ID int64
	UID int64
	Name string
	Locale string
	Picture string
	Gender string
	Email string
	Link string
	CreatedTimeUnix int64
	AccessToken int64
	RefreshToken int64
	ExpiresTime int64
}

type UsersTwitter struct {
	ID int64
	UID int64
	Name string
	ScreenName string
	Location string
	TimeZone string
	Email string
	Lang string
	CreatedTimeUnix int64
	AccessToken int64
}

type UserActionHistoryData struct {
	ID int64
	Content string
	Attached string
	AddonData string
}

type UserActionHistoryFresh struct {
	ID int64
	HistoryID int64
	AssociateID int64
	AssociateType string
	CreatedTimeUnix int64
	UID int64
	Anonymous bool
}

type UserFollow struct {
	ID int64
	FansUID int64
	FriendUID int64
	CreatedTimeUnix int64
}

type WeixinReplyRule struct {
	ID int64
	AccountID int64
	KeyWord string
	Title string
	ImageFile string
	Description string
	Link string
	Enabled bool
	SortStatus int
}

type WeixinAccounts struct {
	ID int64
	WeixinmpToekn string
	WeixinAccountRole string
	WeixinAPPID string
	WeixinAPPSecret string
	WeixinmpMenu string
	WeixinSubscribeMessageKey string
	WeixinNoResultMessageKey string
	WeixinEncodingAESKey string
}

type AccountMsg struct {
	ID int64
	Type string
	MsgAuthID int64
	MsgID int64
	AccessKey string
	GroupName string
	Status string
	ErrorNum int
	MainMsg string
	ArticlesInfo string
	ItemID int64
	QuestionsInfo string
	QuestionID int64
	CreatedTimeUnix int64
	FilterCount int
	HasAttach bool
}

type WeixinQRCode struct {
	ID int64
	Ticket string
	Description string
	SubsribeNum int
}



// BeforeInsert before insert what will do
func (u *User) BeforeInsert() {
	u.RegTime = time.Now()
	u.OnlineTime = u.RegTime
	u.LastLogin = u.RegTime
	u.LastActive = u.RegTime
}

// BeforeUpdate before update what will do
func (u *User) BeforeUpdate() {
	u.OnlineTime = time.Now()
	u.LastLogin = u.OnlineTime
	u.LastActive = u.OnlineTime
}

// UpdateRegTime update registry time
func (u *User) UpdateRegTime() {
	u.RegTime = time.Now()
}

// GetUserByName returns a user by given name.
func GetUserByName(name string) (*User, error) {
	if len(name) == 0 {
		return nil, errors.UserNotExist{UserID: 0, UserName: name}
	}
	u := &User{UserName: strings.ToLower(name)}
	has, err := engine.Get(u)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, errors.UserNotExist{UserID: 0, UserName: name}
	}
	return u, nil
}

// IsUserExist check given user name and id
func IsUserExist(uid int64, name string) (bool, error) {
	if len(name) == 0 {
		return false, nil
	}
	return engine.Where("id != ?", uid).Get(&User{UserName: name})
}

// GetUserSalt return user's salt token
func GetUserSalt() (string, error) {
	return utils.RandomString(10)
}

// ResponseFormat reponse format
func (u *User) ResponseFormat() *account.Account {
	return &account.Account{
		ID: u.ID,
		Email: u.Email,
	}
}

// GenerateEmailActivateCode generates an activate code based on user information and given e-mail.
func (u *User) GenerateEmailActivateCode(email string) string {
	code := utils.CreateTimeLimitCode(
	utils.ToStr(u.ID)+email+u.Passwd,
	setting.Service.ActiveCodeLives, nil)

	// Add tail hex username
  code += hex.EncodeToString([]byte(u.UserName))
  return code
}

// GenerateActivateCode generates an activate code based on user information.
func (u *User) GenerateActivateCode() string {
		return u.GenerateEmailActivateCode(u.Email)
}

// CustomAvatarPath custom avatar path
func (u *User) CustomAvatarPath() string {
	return filepath.Join(setting.AvatarUploadPath, utils.ToStr(u.ID))
}

// GenerateRandomAvatar generates a random avatar for user.
func (u *User) GenerateRandomAvatar() (err error) {
	email := u.Email
	if len(u.Email) == 0 {
		email = u.UserName
	}

	img, err := avatar.RandomImage([]byte(email))
	if err != nil {
		return fmt.Errorf("RandomImage error: %v", err)
	}

	if err = os.MkdirAll(filepath.Dir(u.CustomAvatarPath()), os.ModePerm); err !=nil {
		return fmt.Errorf("RandomImage mkdir: %v", err)
	}

	f, err := os.Create(u.CustomAvatarPath())
	if err != nil {
		return fmt.Errorf("RandomImage create: %v", err)
	}
	defer f.Close()
	if err = png.Encode(f, img); err != nil {
		return fmt.Errorf("RandomImage Encode: %v", err)
	}
	glog.Info("user %v avatar created", u.UserName)
	return nil
}

// RelAvatarLink return relative avatar link
func (u *User) RelAvatarLink() string {
	defaultImg := setting.APPSubURL + "/img/avatar_default.png"

	if u.ID == -1 {
		return defaultImg
	}
	switch {
	case setting.DisableGravatar, setting.OfflineMode:
			if !utils.IsExist(u.CustomAvatarPath()) {
				if err := u.GenerateRandomAvatar(); err != nil {
					glog.V(3).Infof("GenerateRandomAvatar: %v", err)
				}
			}
			return setting.APPSubURL + "/avatars/" + utils.ToStr(u.ID)
	}
	return utils.AvatarLink(u.Email)
}

// AvatarLink returns user avatar abs link
func (u *User) AvatarLink() string {
	link := u.RelAvatarLink()
	if link[0] == '/' && link[1] != '/' {
		return setting.APPSubURL + strings.TrimPrefix(link, setting.APPSubURL)[1:]
	}
	return link
}

// GetFollowers return given user's followers
func (u *User) GetFollowers(page int) ([]*User, error) {
	users := make([]*User, 0, ItemsPerPage)
	session := engine.Limit(ItemsPerPage, (page - 1)*ItemsPerPage).Where("follow.follow_id=?", u.ID)
	switch  DbCfg.Type{
	case "postgres":
		session.Join("LEFT", "follow", `"user".id=follow.user_id`)
	default:
		session.Join("LEFT", "follow", "user.id=follow.user_id")
	}
	return users, session.Find(&users)
}

// IsFollowing return wether follow or not
func (u *User) IsFollowing(followID int64) bool {
	return IsFollowing(u.ID, followID)
}

// GetFollowing return user's focus user
func (u *User) GetFollowing(page int) ([]*User, error) {
	users := make([]*User, 0, ItemsPerPage)
	session := engine.Limit(ItemsPerPage, (page - 1)*ItemsPerPage).Where("follow.follow_id=?", u.ID)
	switch  DbCfg.Type{
	case "postgres":
		session.Join("LEFT", "follow", `"user".id=follow.follow_id`)
	default:
		session.Join("LEFT", "follow", "user.id=follow.follow_id")
	}
	return users, session.Find(&users)
}

// EncodePasswd encodes password with salt
func (u *User) EncodePasswd() {
	u.Passwd = fmt.Sprintf("%x", pbkdf2.Key([]byte(u.Passwd), []byte(u.Salt), 10000, 40, sha256.New))
}

// ValidatePasswd checks  the valid of given passwd
func (u *User)ValidatePasswd(passwd string) bool {
	newUser := &User{Passwd: passwd, Salt: u.Salt}
	newUser.EncodePasswd()
	return subtle.ConstantTimeCompare([]byte(u.Passwd), []byte(newUser.Passwd)) == 1
}

// UploadAvatar upload avatar
func (u *User) UploadAvatar(data []byte) error {
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("image Decode: %v", err)
	}

	size := resize.Resize(avatar.AVATARSIZE, avatar.AVATARSIZE, img, resize.NearestNeighbor)
	session := engine.NewSession()
	defer session.Close()
	if err = session.Begin(); err != nil {
		return fmt.Errorf("image session Begine: %v", err)
	}

	os.MkdirAll(setting.AvatarUploadPath, os.ModePerm)
	f, err := os.Create(u.CustomAvatarPath())
	if err != nil {
		return fmt.Errorf("image dir create: %v", err)
	}
	defer f.Close()

	if err = png.Encode(f, size); err != nil {
		return fmt.Errorf("image Encode: %v", err)
	}

	return session.Commit()
}

// DeleteAvatar delete user's avatar
func (u *User) DeleteAvatar() error {
	os.Remove(u.CustomAvatarPath())
	u.AvatarFile = ""
	if err := UpdateUser(u); err != nil {
		return fmt.Errorf("UpdateUser: %v", err)
	}
	return nil
}


func countUsers(e Engine) int64 {
	num, _ := e.Where("name IS NOT NULL ").Count(new(User))
	return num
}
// CountUsers return the number of users
func CountUsers() int64 {
	return countUsers(engine)
}

// isUsableName check wether name is reserved or not.
func isUsableName(names, patterns []string, name string) error {
		name = strings.TrimSpace(strings.ToLower(name))
		if utf8.RuneCountInString(name) == 0 {
			return errors.EmptyName{}
		}

		for n := range names {
			if name == names[n] {
				return errors.ErrNameReserved{Name: name}
			}
		}

		for _, pat := range patterns {
			if pat[0] == '*' && strings.HasSuffix(name, pat[1:]) || (pat[len(pat)-1] == '*' && strings.HasPrefix(name, pat[:len(pat)-1])) {
				return errors.ErrNamePatternNotAllowed{Pattern: pat}
		}
	}
		return nil
}

// IsUsableUserName ...
func IsUsableUserName(name string) error {
	return isUsableName(reservedUsernames, reservedUserPatterns, name)
}

// CreateUser create a new user record
func CreateUser(u *User) error {
	if err := IsUsableUserName(u.UserName); err != nil {
		return err
	}

	exist, err := IsUserExist(0, u.UserName)
	if err != nil {
		return err
	}else if exist {
		return errors.ErrUserAlreadyExist{Name: u.UserName}
	}

	u.AvatarFile = u.CustomAvatarPath()
	if u.Salt, err = GetUserSalt(); err != nil {
		return err
	}

	u.EncodePasswd()

	session := engine.NewSession()
	defer session.Close()

	if err = session.Begin(); err != nil {
		return err
	}

	if _, err = session.Insert(u); err != nil {
		return err
	} else if err = os.MkdirAll(UserPath(u.UserName), os.ModePerm); err != nil {
		return err
	}
	return session.Commit()
}

// UserPath returns the path absolute path of user .
func UserPath(userName string) string {
	return filepath.Join(setting.RootPath, strings.ToLower(userName))
}

// Users return users number in given pages
func Users(page, size int) ([]*User, error) {
	users := make([]*User, 0, size)
	return users, engine.Limit(size, (page-1)*size).Where("id IS NOT NULL").Asc("id").Find(&users)
}

// parseUserFromCode returns user by username encoded in code.
// It returns nil if code or username is invalid.
func parseUserFromCode(code string) (user *User) {
	if len(code) <= utils.TIME_LIMIT_CODE_LENGTH {
		return nil
	}

	// Use tail hex username to query user
	hexStr := code[utils.TIME_LIMIT_CODE_LENGTH:]
	if b, err := hex.DecodeString(hexStr); err == nil {
		if user, err = GetUserByName(string(b)); user != nil {
			return user
		} else if !errors.IsUserNotExist(err) {
			glog.V(2).Infof("GetUserByName: %v", err)
		}
	}

	return nil
}

// VerifyUserActiveCode verify active code
func VerifyUserActiveCode(code string) *User {
	ms := setting.Service.ActiveCodeLives
	if use := parseUserFromCode(code); use != nil {
		prefix := code[:utils.TIME_LIMIT_CODE_LENGTH]
		data := convert.ToStr(use) + use.Email + use.UserName + use.Passwd
		if utils.VerifyTimeLimitCode(data, ms, prefix) {
			return use
		}
	}
	return nil
}

// VerifyActiveEmailCode return verify active code when active account
func VerifyActiveEmailCode(code, email string) *Email {
	ms := setting.Service.ActiveCodeLives

	if user := parseUserFromCode(code); user != nil {
		// time limit code
		prefix := code[:utils.TIME_LIMIT_CODE_LENGTH]
		data := convert.ToStr(user.ID) + email + user.UserName + user.Passwd

		if utils.VerifyTimeLimitCode(data, ms, prefix) {
			emailAddress := &Email{Email: email}
			if has, _ := engine.Get(emailAddress); has {
				return emailAddress
			}
		}
	}
	return nil
}

// UpdateUserName changes user's name
func UpdateUserName(u *User, newName string) (err error) {
	if err = IsUsableUserName(newName); err != nil {
		return err
	}
	err = updateUser(engine, u)
	if err != nil {
		return err
	}
	return nil
}

// updateUser update user
func updateUser(e Engine, u *User) error {
	_, err := e.Id(u.ID).AllCols().Update(u)
	return err
}

// UpdateUser update user's info
func UpdateUser(u *User) error {
	return updateUser(engine, u)
}

// deleteBeans deletes all given beans, beans should contain delete conditions.
func deleteBeans(e Engine, beans ...interface{}) (err error) {
	for bean := range beans {
		if _, err = e.Delete(bean); err != nil {
			return err
		}
	}
	return nil
}

//deleteUser delete specific user
func deleteUser(e *xorm.Session, u *User) (err error) {
	// need update watch title

	// update Follow
	followers := make([]*Follow, 0, 10)
	if err = e.Find(&followers, &Follow{UserID: u.ID}); err != nil {
		return fmt.Errorf("get followers: %v", err)
	}
	for _, follower := range followers {
		if _, err = e.Exec("UPDATE `user` SET friend_count=friend_count-1 WHERE id=?", follower.UserID); err != nil {
			return fmt.Errorf("decrease user follower number[%d]: %v", follower.UserID, err)
		}
	}

	if err = deleteBeans(engine, &Follow{FollowID: u.ID}); err != nil {
		return fmt.Errorf("deleteBeans: %v", err)
	}
	os.RemoveAll(UserPath(u.UserName))
	os.Remove(u.CustomAvatarPath())
	return nil
}

// DeleteUser completely and permanently delete everything about the given user.
func DeleteUser(u *User) (err error) {
	session := engine.NewSession()
	defer session.Close()

	if err = session.Begin(); err != nil {
		return err
	}

	if err = deleteUser(session, u); err != nil {
		return err
	}

	return session.Commit()

}

// getUserByID get user according given userID
func getUserID(e Engine, id int64) (*User, error) {
	u := new(User)
	exist, err := e.Id(id).Get(u)
	if err != nil {
		return nil, err
	} else if !exist {
		return nil, errors.UserNotExist{UserName: "", UserID: id}
	}
	return u, nil
}

// GetUserByID ...
func GetUserByID(id int64) (*User, error) {
	return getUserID(engine, id)
}

// GetUserIDsByNames returns a slice of ids corresponds to names.
func GetUserIDsByNames(names []string) []int64 {
	ids := make([]int64, 0, len(names))
	for _, name := range names {
		u, err := GetUserByName(name)
		if err != nil {
			continue
		}
		ids = append(ids, u.ID)
	}
	return ids
}

// GetUserByEmail returns the user object by given e-mail if exists.
func GetUserByEmail(email string) (*User, error) {
	if len(email) == 0 {
		return nil, errors.UserNotExist{UserID: 0, UserName: "email"}
	}

	email = strings.ToLower(email)
	// First try to find the user by primary email
	user := &User{Email: email}
	exist, err := engine.Get(user)
	if err != nil {
		return nil, err
	}
	if exist {
		return user, nil
	}


	emailAddress := &Email{Email: email, IsActivation: true}
	has, err := engine.Get(emailAddress)
	if err != nil {
		return nil, err
	}
	if has {
		return GetUserByID(emailAddress.UID)
	}

	return nil, errors.UserNotExist{UserID: 0, UserName: email}
}
// Follow define relationship of user and its Follower.
type Follow struct {
	ID         int64
	UserID     int64 `xorm:"UNIQUE(follow)"`
	FollowID      int64 `xorm:"UNIQUE(follow)"`
	FollowTime time.Time
}

// IsFollowing checkout wether user  has beed focused by followID
func IsFollowing(userID, followID int64) bool {
	follow, _ := engine.Get(&Follow{UserID: userID, FollowID: followID})
	return follow
}

// Followed userID action follow
func Followed(userID, followID int64) (err error){
	// check followID not equal userID and has beed Followed
	if userID == followID || IsFollowing(userID, followID) {
		return nil
	}
	session := engine.NewSession()
	defer session.Close()
	if err = session.Begin(); err != nil {
		return err
	}

	if _, err = session.Insert(&Follow{UserID: userID, FollowID: followID}); err != nil {
		return err
	}

	if _, err = session.Exec("UPDATE `user` SET fans_count = fans_count + 1 where id = ?", followID); err != nil {
		return err
	}

	if _, err = session.Exec("UPDATE `user` SET friend_count = friend_count + 1 where id = ?", userID); err != nil {
		return err
	}

	return session.Commit()
}

// Unfollowed unfollow people by user
func Unfollowed(userID, followID int64) (err error) {
	if userID == followID || !IsFollowing(userID, followID) {
		return nil
	}

	session := engine.NewSession()
	defer session.Close()
	if err = session.Begin(); err != nil {
		return err
	}

	if _, err = session.Delete(&Follow{UserID: userID, FollowID: followID}); err != nil {
		return err
	}

	if _, err = session.Exec("UPDATE `user` SET fans_count = fans_count - 1 where id = ?", followID); err != nil {
		return err
	}

	if _, err = session.Exec("UPDATE `user` SET friend_count = friend_count - 1 where id = ?", userID); err != nil {
		return err
	}
	return session.Commit()
}
