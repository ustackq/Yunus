package mail

import (
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"net/smtp"
	"os"
	"strings"
	"time"

	"gopkg.in/gomail.v2"

	"github.com/golang/glog"
	"github.com/ustack/Yunus/src/app/backend/pkg/utils/html2text"
)

// Mailer represents mail service.
type Mailer struct {
	QueueLength       int
	Subject           string
	Host              string
	Port              string
	From              string
	FromEmail         string
	User, Passwd      string
	DisableHelo       bool
	HeloHostname      string
	SkipVerify        bool
	UseCertificate    bool
	CertFile, KeyFile string
	UsePlainText      bool
}

var (
	// MailService ...
	MailService *Mailer
)

type Message struct {
	// Log info
	Info string
	*gomail.Message
	confirmChan chan struct{}
}

type loginAuth struct {
	username, password string
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte{}, nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		default:
			return nil, fmt.Errorf("unknwon fromServer: %s", string(fromServer))
		}
	}
	return nil, nil
}

// SMTP AUTH LOGIN Auth Handler
func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}

func NewMessage(to []string, from, subject, htmlBody string) *Message {
	msg := gomail.NewMessage()
	msg.SetHeader("From", from)
	msg.SetHeader("To", to...)
	msg.SetHeader("Subject", subject)
	msg.SetDateHeader("Data", time.Now())

	contentType := "text/html"
	body := htmlBody
	body, err := html2text.FromString(htmlBody)
	if err != nil {
		glog.V(2).Infof("html2context: %v", err)
	}

	msg.SetBody(contentType, body)
	return &Message{
		Message:     msg,
		confirmChan: make(chan struct{}),
	}
}

type Sender struct{}

func (sender *Sender) Send(from string, to []string, msg io.WriterTo) error {
	opts := MailService
	tlsConfig := &tls.Config{
		InsecureSkipVerify: opts.SkipVerify,
		ServerName:         opts.Host,
	}

	if opts.UseCertificate {
		cert, err := tls.LoadX509KeyPair(opts.CertFile, opts.KeyFile)
		if err != nil {
			return err
		}
		tlsConfig.Certificates = []tls.Certificate{cert}
	}

	// Server port check
	conn, err := net.Dial("tcp", net.JoinHostPort(opts.Host, opts.Port))
	if err != nil {
		return err
	}
	defer conn.Close()

	isSecureConn := false
	if strings.HasSuffix(opts.Port, "465") {
		conn = tls.Client(conn, tlsConfig)
		isSecureConn = true
	}

	client, err := smtp.NewClient(conn, opts.Host)
	if err != nil {
		return fmt.Errorf("SMTP NewClient: %v", err)
	}

	if !opts.DisableHelo {
		hostname := opts.HeloHostname
		if len(hostname) == 0 {
			hostname, err = os.Hostname()
			if err != nil {
				return err
			}
		}

		if err = client.Hello(hostname); err != nil {
			return fmt.Errorf("Hello: %v", err)
		}
	}

	// If not using SMTPS, alway use STARTTLS if available
	hasStartTLS, _ := client.Extension("STARTTLS")
	if !isSecureConn && hasStartTLS {
		if err = client.StartTLS(tlsConfig); err != nil {
			return fmt.Errorf("StartTLS: %v", err)
		}
	}

	canAuth, options := client.Extension("AUTH")
	if canAuth && len(opts.User) > 0 {
		var auth smtp.Auth

		if strings.Contains(options, "CRAM-MD5") {
			auth = smtp.CRAMMD5Auth(opts.User, opts.Passwd)
		} else if strings.Contains(options, "PLAIN") {
			auth = smtp.PlainAuth("", opts.User, opts.Passwd, opts.Host)
		} else if strings.Contains(options, "LOGIN") {
			// Patch for AUTH LOGIN
			auth = LoginAuth(opts.User, opts.Passwd)
		}

		if auth != nil {
			if err = client.Auth(auth); err != nil {
				return fmt.Errorf("Auth: %v", err)
			}
		}
	}

	if err = client.Mail(from); err != nil {
		return fmt.Errorf("Mail: %v", err)
	}

	for _, rec := range to {
		if err = client.Rcpt(rec); err != nil {
			return fmt.Errorf("Rcpt: %v", err)
		}
	}

	w, err := client.Data()
	if err != nil {
		return fmt.Errorf("Data: %v", err)
	} else if _, err = msg.WriteTo(w); err != nil {
		return fmt.Errorf("WriteTo: %v", err)
	} else if err = w.Close(); err != nil {
		return fmt.Errorf("Close: %v", err)
	}

	return client.Quit()

}

var mailQ chan *Message

func processMailQ() {
	sender := &Sender{}
	for {
		select {
		case msg := <-mailQ:
			glog.V(0).Infof("New email sending request to %s: %s", msg.GetHeader("To"), msg.Info)
			if err := gomail.Send(sender, msg.Message); err != nil {
				glog.V(2).Infof("Fail to send email to %s: %s: %v", msg.GetHeader("To"), msg.Info, err)
			} else {
				glog.V(0).Infof("email to %s has been sended: %s", msg.GetHeader("To"), msg.Info)
			}
			msg.confirmChan <- struct{}{}
		}
	}
}

func NewMailer() {
	if mailQ != nil {
		return
	}
	mailQ = make(chan *Message, MailService.QueueLength)
	go processMailQ()
}

func Send(msg *Message) {
	mailQ <- msg
	go func() {
		<-msg.confirmChan
	}()
}
