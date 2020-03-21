package models

import (
	"time"
)

// User is model to show users in app
type User struct {
	ID       int64
	Login    string
	Email    string
	Password string
	Phone    string
	Role     string
}

// Article is model to show articles in app
type Article struct {
	ID          int64
	Name        string
	Description string
	Text        string `sql:"type:text"`
	CreateAt    time.Time
	Author      User `gorm:"foreignkey:UserRefer"`
	UserRefer   uint
}

//Comment is model to show comments in app
type Comment struct {
	ID           int64
	Text         string `sql:"type:text"`
	CreateAt     time.Time
	Author       User `gorm:"foreignkey:UserRefer"`
	UserRefer    uint
	Article      Article `gorm:"foreignkey:ArticleRefer"`
	ArticleRefer uint
}
