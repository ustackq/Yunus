package models

type Article struct {
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
	VoteID       int64
	DeadLineUnix int64         `xorm:"-"`
	CreatedUnix  int64         `xorm:"-"`
	UpdatedUnix  int64         `xorm:"-"`
	Attachments  []*Attachment `xorm:"-"`
	Comments     []*Comment    `xorm:"-"`
}
