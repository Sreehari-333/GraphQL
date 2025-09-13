package model

import ()

type UserDB struct {
	ID    string `gorm:"primaryKey"`
	Name  string
	Email string
	Posts []PostDB `gorm:"foreignKey:UserID"`
}

type PostDB struct {
	ID      string `gorm:"primaryKey"`
	Title   string
	Content string
	UserID  string
}
