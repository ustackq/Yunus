package models

type Category struct {
	ID       int64
	Title    string
	Type     string
	Icon     string
	ParentID int64
	Sort     int
	URLToken string
}

type Draft struct {
	ID             int64
	UID            int64
	Type           string
	ItemID         int64
	Data           string
	ADDTimeUnix    int64
	UpdateTimeUnix int64
}

type EducationExperience struct {
	ID             int64
	UID            int64
	EduYears       int
	EduSchoolName  string
	EduSchoolType  string
	EduDepartments string
	AddTimeUnix    int64
	UpdateTimeUnix int64
}

type Feature struct {
	ID          int64
	Title       string
	Description string
	Icon        string
	TopicCount  int
	SEOTitle    string
	Enabled     bool
}

type FeatureTopic struct {
	ID        int64
	FeatureID int64
	TopicID   int64
}

type FavoriteTag struct {
	ID     int64
	UID    int64
	Title  string
	ItemID int64
	Type   string
}

type Invitation struct {
	ID               int64
	UID              int64
	InvitationCode   string
	InvitationEmail  string
	AddTimeUnix      int64
	AddIP            string
	ActivateExpire   bool
	ActivateTimeUnix int64
	ActivateIP       string
	ActivateStatus   int
	ActivateUID      int64
}

type Job struct {
	ID      int64
	JobName string
}

type IntegralLog struct {
	ID       int64
	UID      int64
	Action   string
	Integral int
	Note     string
	Balance  int
	ItemID   int64
	TimeUnix int64
}

type NavMenu struct {
	ID          int64
	Title       string
	Description string
	Type        string
	TypeID      int64
	Icon        string
	Sort        int
}

type Inbox struct {
	ID              int64
	UID             int64
	DialogID        int64
	Message         string
	AddTimeUnix     int64
	SenderRemove    bool
	RecipientRemove bool
	Receipt         int64
}

type InboxDialog struct {
	ID              int64
	SenderUID       int64
	SenderUnread    int64
	RecipientUID    int64
	CreatedTimeUnix int64
	UpdatedTimeUnix int64
	RecipientCount  int64
}

type Notification struct {
	ID              int64
	SenderUID       int64
	RecipientUID    int64
	ActionType      string
	ModelType       string
	SourceID        int64
	CreatedTimeUnix int64
	ReadFlag        bool
}

type NotificationData struct {
	NotificationID int64
	Data           string
}

type Focus struct {
	FocusID         int64
	ItemType        string
	ItemID          int64
	UID             int64
	CreatedTimeUnix int64
}

type Invite struct {
	ID               int64
	ItemType         string
	ItemID           int64
	SenderUID        int64
	RecipientsUID    int64
	Email            string
	CreatedTimeUnix  int64
	AvailabeTimeUnix int64
}

type School struct {
	ID       int
	Type     string
	Code     string
	Name     string
	AreaCode int
}

type Sessions struct {
	ID           int64
	Modified     int64
	Data         string
	LifeTimeUnix int64
}

type SearchCache struct {
	ID   int64
	Hash string
	Data string
	Time int64
}

type SystemSetting struct {
	ID      int64
	VarName string
	Value   string
}

type WeixinLogin struct {
	ID        int64
	Token     int
	UID       int64
	SessionID int
	Expire    int
}

type WeixinMessage struct {
	ID       int64
	WeixinID int64
	Content  string
	Action   string
	Time     int64
}

type Redirect struct {
	ID       int64
	ItemID   int64
	TargetID int64
	Time     int64
	UID      int64
}

type WorkExperience struct {
	ID              int64
	UID             int64
	StartYear       int
	EndYear         int
	CompanyName     string
	JobID           int
	CreatedTimeUnix int64
}

type ReputationCategory struct {
	ID              int64
	UID             int64
	CategotyID      int64
	UpdatedTimeUnix int64
	Reputation      int
	GratitudeCount  int
	AgressCount     int
	QuestionCount   int
}

type EdmTask struct {
	ID       int64
	Title    string
	Message  string
	Subject  string
	FromName string
	Time     int64
}

type EdmTaksData struct {
	ID       int64
	TaksID   int64
	Email    string
	SentTime int64
	ViewTime int64
}

type EdmUserData struct {
	ID        int64
	UserGroup int64
	Email     string
}

type EdmUserGroup struct {
	ID    int64
	Title string
	Time  int64
}

type EdmUnSubscription struct {
	ID    int64
	Email string
	Time  int64
}

type VerifyApply struct {
	ID     int64
	UID    int64
	Reason string
	Attach string
	Time   int64
	Name   string
	Data   string
	Status bool
	Type   string
}
