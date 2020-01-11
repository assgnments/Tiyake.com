package entity

import "time"

type Role struct {
	ID   uint
	Name string `gorm:"type:varchar(255)"`
}

type User struct {
	ID       uint
	FullName string `gorm:"type:varchar(255);not null"`
	Email    string `gorm:"type:varchar(255);not null;unique"`
	Password string `gorm:"type:varchar(255)"`
	RoleID   uint
}

type Session struct {
	ID         uint
	UUID       uint
	Expires    int64  `gorm:"type:varchar(255);not null"`
	SigningKey []byte `gorm:"type:varchar(255);not null"`
}

type Question struct {
	ID          uint
	Title       string `gorm:"type:varchar(255);not null"`
	Description string `gorm:"type:varchar(1000);not null"`
	Image       string `gorm:"type:varchar(255)"`
	UserID      string `gorm:"type:varchar(255);not null"`
	CreatedAt   time.Time
	Answers     []Answer `gorm:"one2many:answer"`
}

type Answer struct {
	ID         uint
	UserID     string `gorm:"type:varchar(1000);not null"`
	Message    string `gorm:"type:varchar(1000);not null"`
	QuestionID uint
	UpVotersID []uint
	CreatedAt  time.Time
}
