package models

// Label represents a label of repository for issues.
type Label struct {
	ID              int64
	RepoID          int64 `xorm:"INDEX"`
	Name            string
	Color           string `xorm:"VARCHAR(7)"`
	NumIssues       int
	NumClosedIssues int
	NumOpenIssues   int  `xorm:"-"`
	IsChecked       bool `xorm:"-"`
}

type Topic struct {
	ID                    int64
	Title                 string
	CreatedTimeUnix       int64
	DiscussCount          int
	Description           string
	TopicAvatar           string
	TopicLock             bool
	FocusCount            int64
	UrlToken              string
	MergedID              int64
	SEOTitle              string
	ParentID              int64
	IsParent              bool
	DiscussCountLastWeek  int
	DiscussCountLastMonth int
	DiscussCountUpdate    int
}

type TopicFocus struct {
	ID              int64
	TopicID         int64
	UID             int64
	CreatedTimeUnix int64
}

type TopicMerge struct {
	ID       int64
	SourceID int64
	TargetID int64
	UID      int64
	Time     int64
}

type TopicRelation struct {
	ID              int64
	TopicID         int64
	ItemID          int64
	CreatedTimeUnix string
	UID             int64
	Type            string
}

type GEOLocation struct {
	ID              int64
	ItemID          int64
	ItemType        string
	Latitude        float64
	LongLatitude    float64
	CreatedTimeUnix int64
}
