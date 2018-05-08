package mail

import (
	"fmt"
  "gopkg.in/gomail.v2"
	"github.com/gin-gonic/gin"

	"github.com/golang/glog"
	"github.com/ustack/Yunus/src/app/backend/pkg/setting"
)

const (
	MAIL_AUTH_ACTIVATE        = "auth/activate"
	MAIL_AUTH_ACTIVATE_EMAIL  = "auth/activate_email"
	MAIL_AUTH_RESET_PASSWORD  = "auth/reset_passwd"
	MAIL_AUTH_REGISTER_NOTIFY = "auth/register_notify"

	MAIL_ISSUE_COMMENT = "issue/comment"
	MAIL_ISSUE_MENTION = "issue/mention"

	MAIL_NOTIFY_COLLABORATOR = "notify/collaborator"
)

type User interface {
	ID() int64
	DisplayName() string
	Email() string
	GenerateActivateCode() string
	GenerateEmailActivateCode(string) string
}

type Repository interface {
	FullName() string
	HTMLURL() string
	ComposeMetas() map[string]string
}

type Issue interface {
	MailSubject() string
	Content() string
	HTMLURL() string
}

type HTMLOptions struct {
	// Layout template name. Overrides Options.Layout.
	Layout string
}

// MailRender return text tpl
type MailRender interface {
	HTMLString(string, interface{}, ...HTMLOptions) (string, error)
}

var mailRender MailRender

func SendUserMail(c *gin.Context, u User, from, tpl, code, subject, info string) {
	data := map[string]interface{}{
		"Username":          u.DisplayName(),
		"ActiveCodeLives":   setting.Service.ActiveCodeLives,
		"ResetPwdCodeLives": setting.Service.ResetPwdCodeLives,
		"Code":              code,
	}
	content, err := mailRender.HTMLString(string(tpl), data)
	if err != nil {
		glog.V(3).Infof("HTML render", err)
		return
	}

	message := NewMessage([]string{u.Email()}, from, subject, content)
	message.Info = fmt.Sprintf("User-ID: %d, %s", u.ID(), info)
	Send(message)
}
// SendActivateMail ...
func SendActivateMail(c *gin.Context, u User) {
  SendUserMail(c, u, MailService.From, MAIL_AUTH_ACTIVATE, u.GenerateActivateCode(), "mail.activate_account", "activate account")
}

// SendRestPasswordMail ...
func SendRestPasswordMail(c *gin.Context, u User) {
  SendUserMail(c, u, MailService.From, MAIL_AUTH_RESET_PASSWORD, u.GenerateActivateCode(), "mail.reset_password", "reset password")
}


// SendActivateEmailMail sends confirmation email.
func SendActivateEmailMail(c *gin.Context, u User, email string) {
	data := map[string]interface{}{
		"Username":        u.DisplayName(),
		"ActiveCodeLives": setting.Service.ActiveCodeLives / 60,
		"Code":            u.GenerateEmailActivateCode(email),
		"Email":           email,
	}
	body, err := mailRender.HTMLString(string(MAIL_AUTH_ACTIVATE_EMAIL), data)
	if err != nil {
		glog.V(2).Infof("HTMLString: %v", err)
		return
	}

	msg := NewMessage([]string{email}, MAIL_AUTH_ACTIVATE_EMAIL, "mail.activate_email", body)
	msg.Info = fmt.Sprintf("UID: %d, activate email", u.ID())

	Send(msg)
}
// SendRegisterNotifyMail triggers a notify e-mail by admin created a account.
func SendRegisterNotifyMail(c *gin.Context, u User) {
	data := map[string]interface{}{
		"Username": u.DisplayName(),
	}
	body, err := mailRender.HTMLString(string(MAIL_AUTH_REGISTER_NOTIFY), data)
	if err != nil {
		glog.V(2).Infof("HTMLString: %v", err)
		return
	}

	msg := NewMessage([]string{u.Email()}, MAIL_AUTH_REGISTER_NOTIFY ,"mail.register_notify", body)
	msg.Info = fmt.Sprintf("UID: %d, registration notify", u.ID())

	Send(msg)
}

func composeTplData(subject, body, link string) map[string]interface{} {
	data := make(map[string]interface{}, 10)
	data["Subject"] = subject
	data["Body"] = body
	data["Link"] = link
	return data
}

func composeIssueMessage(issue Issue, repo Repository, doer User, tplName string, tos []string, info string) *Message {
	subject := issue.MailSubject()
	body := tplName
	data := composeTplData(subject, body, issue.HTMLURL())
	data["Doer"] = doer
	content, err := mailRender.HTMLString(tplName, data)
	if err != nil {
		glog.V(2).Infof( "HTMLString (%s): %v", tplName, err)
	}
	from := gomail.NewMessage().FormatAddress(MailService.FromEmail, doer.DisplayName())
	msg := NewMessage(tos, from, subject, content)
	msg.Info = fmt.Sprintf("Subject: %s, %s", subject, info)
	return msg
}

// SendIssueCommentMail composes and sends issue comment emails to target receivers.
func SendIssueCommentMail(issue Issue, repo Repository, doer User, tos []string) {
	if len(tos) == 0 {
		return
	}

	Send(composeIssueMessage(issue, repo, doer, MAIL_ISSUE_COMMENT, tos, "issue comment"))
}

// SendIssueMentionMail composes and sends issue mention emails to target receivers.
func SendIssueMentionMail(issue Issue, repo Repository, doer User, tos []string) {
	if len(tos) == 0 {
		return
	}
	Send(composeIssueMessage(issue, repo, doer, MAIL_ISSUE_MENTION, tos, "issue mention"))
}
