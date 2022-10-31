package domain

type Group struct {
	Name   string
	Color  string
	ChatId int64
	Users  []User
}
