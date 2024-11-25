package models

type Book struct {
	ID     int     `gorm:"primaryKey"`
	Isbn   string  `gorm:"unique;type:varchar(13)"`
	Title  string  `gorm:"type:varchar(255)"`
	Author *Author `gorm:"embedded"`
}

type Author struct {
	Firstname string `gorm:"type:varchar(100)"`
	Lastname  string `gorm:"type:varchar(100)"`
}
