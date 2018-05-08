package models

// CommentType defines whether a comment is just a simple comment, an action (like close) or a reference.
type CommentType int

const (
	// Plain comment, can be associated with a commit (CommitID > 0) and a line (LineNum > 0)
	AnswerComment CommentType = iota
	ArticleComment
)

// Comment represents a comment in commit and issue page.
type Comment struct {
	ID                int64
	Type              CommentType
	AnswerID          int64
	Answer            *User   `xorm:"-"`
	QuestionID        int64   `xorm:"INDEX"`
	Question          *Answer `xorm:"-"`
	Content           string  `xorm:"TEXT"`
	RenderedContent   string  `xorm:"-"`
	CommentCount      int64
	UninterestedCount int64
	AgainstCount      int64
	AggreeCount       int64
	ThanksCount       int64
	CreatedUnix       int64
	UpdatedUnix       int64
	IP                string
	Attachments       []*Attachment `xorm:"-"`
	AtUID             int64
	Votes             *Vote `xorm:"-"`
}
