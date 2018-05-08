package utils

import (
	"crypto/md5"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"

	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"

	"github.com/golang/glog"
	"github.com/nu7hatch/gouuid"
	"github.com/ustack/Yunus/src/app/backend/pkg/setting"
	"github.com/ustack/Yunus/src/app/backend/pkg/types"
	"github.com/ustack/Yunus/src/app/backend/pkg/utils/convert"
	"github.com/ustack/Yunus/src/app/backend/pkg/version"
)

// VerifyTimeLimitCode verify time limit code
func VerifyTimeLimitCode(data string, minutes int, code string) bool {
	if len(code) <= 18 {
		return false
	}

	// split code
	start := code[:12]
	lives := code[12:18]
	if d, err := convert.StrTo(lives).Int(); err == nil {
		minutes = d
	}

	// right active code
	retCode := CreateTimeLimitCode(data, minutes, start)
	if retCode == code && minutes > 0 {
		// check time is expired or not
		before, _ := time.ParseInLocation("200601021504", start, time.Local)
		now := time.Now()
		if before.Add(time.Minute*time.Duration(minutes)).Unix() > now.Unix() {
			return true
		}
	}

	return false
}

const TIME_LIMIT_CODE_LENGTH = 12 + 6 + 40

// Info return build info
func Info() *types.YunusInfo {
	return &types.YunusInfo{
		Name:       "Yunus",
		Release:    version.RELEASE,
		Build:      version.COMMIT,
		Repository: version.REPO,
	}
}

// HashEmail hashes email address to MD5 string.
// https://en.gravatar.com/site/implement/hash/
func HashEmail(email string) string {
	email = strings.ToLower(strings.TrimSpace(email))
	h := md5.New()
	h.Write([]byte(email))
	return hex.EncodeToString(h.Sum(nil))
}

// AvatarLink returns relative avatar link to the site domain by given email,
// which includes app sub-url as prefix. However, it is possible
// to return full URL if user enables Gravatar-like service.
func AvatarLink(email string) (url string) {
	if setting.LibravatarService != nil &&
		strings.Contains(email, "@") {
		var err error
		url, err = setting.LibravatarService.FromEmail(email)
		if err != nil {
			glog.V(1).Infof("AvatarLink.LibravatarService.FromEmail [%s]: %v", email, err)
		}
	}
	if len(url) == 0 && !setting.DisableGravatar {
		url = setting.GravatarSource + HashEmail(email)
	}
	if len(url) == 0 {
		url = setting.APPSubURL + "/img/avatar_default.png"
	}
	return url
}

// InstallCheck return install check
func InstallCheck() (bool, error) {
	cfgPath := os.Getenv("YunusCFG_PATH")
	if cfgPath != "" {
		if _, err := os.Stat(cfgPath); err != nil {
			return false, err
		}
	}
	return true, nil
}

// GenerateSalt return salt six lenth string
func GenerateSalt() string {
	salt, err := uuid.NewV4()
	if err != nil {
		glog.Error(err)
	}
	return salt.String()[:6]
}

// GeneratePasswd ...
func GeneratePasswd(passwd, salt string) string {
	tmp := md5.Sum([]byte(passwd))
	a := tmp[:]
	passwd = string(a) + salt
	return fmt.Sprintf("%x", md5.Sum([]byte(passwd)))
}

// CreateTimeLimitCode generates a time limit code based on given input data.
// Format: 12 length date time string + 6 minutes string + 40 sha1 encoded string
func CreateTimeLimitCode(data string, minutes int, startInf interface{}) string {
	format := "200601021504"

	var start, end time.Time
	var startStr, endStr string

	if startInf == nil {
		// Use now time create code
		start = time.Now()
		startStr = start.Format(format)
	} else {
		// use start string create code
		startStr = startInf.(string)
		start, _ = time.ParseInLocation(format, startStr, time.Local)
		startStr = start.Format(format)
	}

	end = start.Add(time.Minute * time.Duration(minutes))
	endStr = end.Format(format)

	// create sha1 encode string
	sh := sha1.New()
	sh.Write([]byte(data + setting.SecretKey + startStr + endStr + ToStr(minutes)))
	encoded := hex.EncodeToString(sh.Sum(nil))

	code := fmt.Sprintf("%s%06d%s", startStr, minutes, encoded)
	return code
}

type argInt []int

// ToStr convert any type to string.
func ToStr(value interface{}, args ...int) (s string) {
	switch v := value.(type) {
	case bool:
		s = strconv.FormatBool(v)
	case float32:
		s = strconv.FormatFloat(float64(v), 'f', argInt(args).Get(0, -1), argInt(args).Get(1, 32))
	case float64:
		s = strconv.FormatFloat(v, 'f', argInt(args).Get(0, -1), argInt(args).Get(1, 64))
	case int:
		s = strconv.FormatInt(int64(v), argInt(args).Get(0, 10))
	case int8:
		s = strconv.FormatInt(int64(v), argInt(args).Get(0, 10))
	case int16:
		s = strconv.FormatInt(int64(v), argInt(args).Get(0, 10))
	case int32:
		s = strconv.FormatInt(int64(v), argInt(args).Get(0, 10))
	case int64:
		s = strconv.FormatInt(v, argInt(args).Get(0, 10))
	case uint:
		s = strconv.FormatUint(uint64(v), argInt(args).Get(0, 10))
	case uint8:
		s = strconv.FormatUint(uint64(v), argInt(args).Get(0, 10))
	case uint16:
		s = strconv.FormatUint(uint64(v), argInt(args).Get(0, 10))
	case uint32:
		s = strconv.FormatUint(uint64(v), argInt(args).Get(0, 10))
	case uint64:
		s = strconv.FormatUint(v, argInt(args).Get(0, 10))
	case string:
		s = v
	case []byte:
		s = string(v)
	default:
		s = fmt.Sprintf("%v", v)
	}
	return s
}

func (a argInt) Get(i int, args ...int) (r int) {
	if i >= 0 && i < len(a) {
		r = a[i]
	} else if len(args) > 0 {
		r = args[0]
	}
	return
}

// IsExist checks whether a file or directory exists.
// It returns false when the file or directory does not exist.
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

// RandomString returns generated random string in given length of characters.
// It also returns possible error during generation.
func RandomString(n int) (string, error) {
	buffer := make([]byte, n)
	max := big.NewInt(int64(len(alphanum)))

	for i := 0; i < n; i++ {
		index, err := randomInt(max)
		if err != nil {
			return "", err
		}

		buffer[i] = alphanum[index]
	}

	return string(buffer), nil
}

func randomInt(max *big.Int) (int, error) {
	rand, err := rand.Int(rand.Reader, max)
	if err != nil {
		return 0, err
	}

	return int(rand.Int64()), nil
}
