package models

type Post struct {
	ID     uint   `gorm:"primariKey" json:"id"`
	Title  string `json:"title"`
	UserID uint   `json:"user_id"`
}
