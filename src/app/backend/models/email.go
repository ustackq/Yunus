package models

// Email definition
type Email struct {
  ID int64
  UID int64
  Email string
  IsActivation bool
}

// EmailQueue ...
type EmailQueue struct {
  ID int64
  Recipient string
  Subject string
  Message string
  IsError bool
  ErrorMessage string
}

// ReceivingEmailConfig ...
type ReceivingEmailConfig struct {
  ID int64
  Protocol string
  Server string
  SSL bool
  Port int
  UserName string
  Password string
  UID int64
  AccessKey string
  HasAttach bool
}

type ReceivedEmail struct {
  ID int64
  UID int64
  ConfigID int64
  MessageID int64
  Time int64
  From string
  Subject string
  Content string
  QuestionID int64
  TicketID int64
}

type WeixinThirdPartyAPI struct {
  ID int64
  AccountID int64
  URL string
  Token string
  Enabled bool
  Rank int
}

type HelpChapter struct {
  ID int64
  Title string
  Description string
  URLToken string
  Sort int
}








//Activate mark email as active
func (email *Email) Activate() (err error) {
  user, err := GetUserByID(email.UID)
  if err != nil {
    return err
  }

  session := engine.NewSession()
  defer session.Close()

  if err = session.Begin(); err != nil {
    return err
  }

  email.IsActivation = true

  if _, err := session.ID(email.ID).AllCols().Update(email); err != nil {
    return err
  } else if err = updateUser(session, user); err != nil {
    return err
  }

  return session.Commit()
}
