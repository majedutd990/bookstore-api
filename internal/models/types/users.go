package types

type (
	UserStatus uint

	Gender uint
)

const (
	_ UserStatus = iota
	UserActive
	UserInactive
	UserBan
)

const (
	_ Gender = iota
	Male
	Female
	More
)
