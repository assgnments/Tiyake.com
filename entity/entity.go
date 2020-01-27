package entity

import (
	"github.com/jinzhu/gorm"
)

type Role struct {
	gorm.Model
	Name string `gorm:"type:varchar(255)"`
}

type User struct {
	gorm.Model
	FullName string `gorm:"type:varchar(255);not null"`
	Email    string `gorm:"type:varchar(255);not null;unique"`
	Password string `gorm:"type:varchar(255)"`
	RoleID   uint
}

type Session struct {
	gorm.Model
	UUID       uint
	SessionId  string `gorm:"type:varchar(255);not null"`
	Expires    int64  `gorm:"type:varchar(255);not null"`
	SigningKey []byte `gorm:"type:varchar(255);not null"`
}

type Question struct {
	gorm.Model
	Title       string `gorm:"type:varchar(255);not null"`
	Description string `gorm:"type:varchar(1000);not null"`
	Image       string `gorm:"type:varchar(255)"`
	UserID      uint
	CategoryID  uint
	User        User `gorm:"many2many:user;"`
	Answers     []Answer
}

type Category struct {
	gorm.Model
	Name string `gorm:"type:varchar(255);not null"`
}

type Answer struct {
	gorm.Model
	UserID     uint
	User       User   `gorm:"many2many:user;"`
	Message    string `gorm:"type:varchar(1000);not null"`
	QuestionID uint
}

type UpVote struct {
	gorm.Model
	UserID   uint
	User     User `gorm:"many2many:user;"`
	AnswerID uint
	Answer   Answer `gorm:"many2many:answer;"`
}
