package models

// Email definition
type Email struct {
  ID int64
  UID int64
  Email string
  IsActivation bool
}

// GetEmailByID ...
func GetEmailByID(uid int64)(*Email, error){
  email:= &Email{}

	 rows, err := x.Query("select * from user", uid)
	 if err != nil {
		 return nil, err
	 }
	 defer rows.Close()
	 for rows.Next() {
		 err = rows.Scan(email)
		 if err != nil {
			 return nil, err
		 }
	 }
   return email, err
}

// GetEmail ...
func GetEmail(email string)(*Email, error){
  email = &Email{}

	 rows, err := x.Query("select * from user", email)
	 if err != nil {
		 return false, err
	 }
	 defer rows.Close()
	 for rows.Next() {
		 err = rows.Scan(email)
		 if err != nil {
			 return false, err
		 }
	 }
   return email, err
}

// IsEmailUsed ...
func IsEmailUsed(email string) (bool, error) {
  if len(email) == 0 {
    return true, nil
  }
  _, err := GetEmailByID(uid)
  if err != nil {
    return false, err
  }
    return true, nil
}
