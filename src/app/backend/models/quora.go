package models

import (
	"github.com/ustack/Yunus/src/app/backend/errors"
)

var (
	ErrMissingIssueNum = errors.New("No question number specified")
)

type Approval struct {
	ID       int64
	Type     string
	Data     string
	UID      int64
	TimeUnix int64
}

// Question represents an question
type Answer struct {
	ID           int64
	PosterID     int64
	Poster       *User  `xorm:"-"`
	Title        string `xorm:"name"`
	Abstract     string `xorm:"abstract"`
	Content      string `xorm:"TEXT"`
	ViewCount    int
	FocusCount   int
	ComentCount  int
	AgressCount  int
	ThanksCount  int
	BestAnswer   int64
	LastAnswer   int64
	PopularValue float64
	IP           string
	Topics       []*Topic `xorm:"-"`
	Priority     int
	IsClosed     bool
	IsLocked     bool
	IsRead       bool `xorm:"-"`

	DeadLineUnix int64         `xorm:"-"`
	CreatedUnix  int64         `xorm:"-"`
	UpdatedUnix  int64         `xorm:"-"`
	Attachments  []*Attachment `xorm:"-"`
	Comments     []*Comment    `xorm:"-"`
}

type Question struct {
	ID                    int64
	Content               string
	Detail                string
	CreatedTimeUnix       int64
	UpdatedTimeUnix       int64
	PublishedUID          int64
	AnswerCount           int
	AnswerUser            int64
	ViewCount             int
	FocusCount            int64
	CommentCount          int
	ActionHistoryID       int64
	CategoryID            int
	AgressCount           int
	AgainstCount          int
	BestAnswer            int64
	HasAttach             bool
	UnverifiedModify      string
	UnverifiedModifyCount int
	IP                    string
	LastAnswer            int64
	PopularValue          int
	PopularValueUpdate    int
	Lock                  bool
	Anonymous             bool
	GratitudeCount        int
	QuestionContent       string
	IsRecommend           bool
	WeiboMsgID            int64
	ReceivedEmailID       int64
	ChapterID             int64
	Sort                  int
}

type QuestionFoucus struct {
	ID          int64
	QuestionID  int64
	UID         int64
	AddTimeUnix int64
}

type QuestionInvite struct {
	ID               int64
	QuestionID       int64
	SenderID         int64
	RecipientID      int64
	Email            string
	AddTimeUnix      int64
	AvaiableTimeUnix int64
}

type Uninterested struct {
	ID              int64
	ItemType        string
	ItemID          int64
	UID             int64
	User            *User `xorm:"-"`
	CreatedTimeUnix int64
}

type Report struct {
	ID              int64
	UID             int64
	Type            string
	TargetID        int64
	Reason          string
	CreatedTimeUnix int64
}

type ReputationTopic struct {
	ID              int64
	UID             int64
	TopicID         int64
	TopicCount      int64
	UpdatedTimeUnix int64
	AgreeCount      int
	GratitudeCount  int
	Reputation      int64
}

type RelatedTopic struct {
	ID        int64
	TopicID   int64
	RelatedID int64
}

type RelatedLinks struct {
	ID              int64
	UID             int64
	ItemType        string
	ItemID          int64
	Link            string
	CreatedTimeUnix int64
}
