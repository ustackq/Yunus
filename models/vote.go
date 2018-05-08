package models

type Vote struct {
	ID               int64
	UID              int64
	Type             int
	ItemID           int64
	Rating           int
	TimeUnix         int64
	ReputationFactor int
	ItemUID          int64
	VoteValue        int
}
