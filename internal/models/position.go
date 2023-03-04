package models

import (
	"time"
)

type Position struct {
	PositionID uint32     `json:"position_id" gorm:"size:11;primary_key:true;not null"`
	UserID     uint32     `json:"user_id" gorm:"size:11;not null"`
	Asset      string     `json:"price"`
	Side       string     `json:"side"`
	Size       uint32     `json:"size"`
	EntryPrice float64    `json:"entry_price"`
	Leverage   uint32     `json:"leverage"`
	Status     uint32     `json:"status"`
	MarketID   uint32     `json:"market_id"`
	ClosedAt   *time.Time `json:"closed_at"`
	UserInfo   *User      `gorm:"foreignKey:user_id;references:user_id"`
	CreatedAt  *time.Time `json:"created_at" gorm:"<-:create"`
	UpdatedAt  *time.Time
}

func init() {

}

func (m *Position) BeforeCreate() {

}

func (m Position) TableName() string {
	return "tb_position"
}
