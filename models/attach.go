package models

import (
	"time"
)

// Attachment represent a attachment of issue/comment/release.
type Attachment struct {
	ID           int64
	UUID         string `xorm:"uuid UNIQUE"`
	FileName     string
	FileLocation string
	AccessKey    int64
	ItemType     int
	ItemID       int64
	IssueID      int64 `xorm:"INDEX"`
	CommentID    int64
	Created      time.Time `xorm:"-"`
	CreatedUnix  int64
	WaitApproval bool
}
