package models

type Book struct {
	ID     int64  `gorm:"primaryKey" json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}
