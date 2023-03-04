package models

import (
	"time"
)

type User struct {
	UserID       uint32     `json:"user_id" gorm:"size:11;primary_key:true;not null"`
	GoogleID     string     `json:"google_id"`
	Email        string     `json:"email"`
	RefreshToken string     `json:"refresh_token"`
	CreatedAt    *time.Time `json:"created_at" gorm:"<-:create"`
	UpdatedAt    *time.Time
}

func init() {

}

func (m *User) BeforeCreate() {

}

func (m User) TableName() string {
	return "tb_user"
}
