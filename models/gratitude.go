package models

const (
	ArticleItem = iota
	AnswerItem
	CommentItem
)

type Gratitude struct {
	ID       int64
	UID      int64
	User     *User
	ItemID   int64
	ItemType int
	TimeUnix int64
}
