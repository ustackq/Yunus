package models

type Page struct {
	ID          int64
	Title       string
	Keywords    string
	Description string
	Content     string
	Enabled     bool
}

type PostIndex struct {
	ID              int64
	PostID          int64
	PostType        string
	CreatedTimeUnix int64
	UpdatedTimeUnix int64
	CategoryID      int64
	IsRecommend     bool
	ViewCount       int64
	Anonymous       bool
	PopularValue    int
	UID             int64
	Lock            bool
	AgressCount     int64
	AnswerCount     int64
}
