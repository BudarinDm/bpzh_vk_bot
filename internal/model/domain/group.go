package domain

type Group struct {
	Name   string  `firestore:"name"`
	Color  string  `firestore:"color"`
	ChatId int64   `firestore:"chatid"`
	Users  []int64 `firestore:"users"`
}
