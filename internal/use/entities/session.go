package entities

type Session struct {
	Timed

	UserID int
	Token  string
}
