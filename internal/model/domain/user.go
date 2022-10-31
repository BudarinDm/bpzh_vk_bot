package domain

type User struct {
	FirstName string             `firestore:"firstname"`
	LastName  string             `firestore:"lastname"`
	TgId      string             `firestore:"tgid"`
	VkId      int64              `firestore:"vkid"`
	Roles     map[string][]int64 `firestore:"roles"`
}
