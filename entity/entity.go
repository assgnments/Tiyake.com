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
	UUID       string `gorm:"type:varchar(255);not null"`
	Expires    int64  `gorm:"type:varchar(255);not null"`
	SigningKey []byte `gorm:"type:varchar(255);not null"`
}


type Question struct {
	ID          uint
	Title       string
	Description string
	Image       string
	UserID      string
	CreatedAt   time.Time
	Comments    [] Comment
}

type Comment struct {
	ID        uint
	UserID    string
	Message   string
	CreatedAt time.Time
}

