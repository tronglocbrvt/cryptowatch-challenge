package models

import (
	"time"
)

type Price struct {
	PriceID   uint32     `json:"price_id" gorm:"size:11;primary_key:true;not null"`
	MarketID  uint32     `json:"market_id"`
	Price     float64    `json:"price"`
	CreatedAt *time.Time `json:"created_at" gorm:"<-:create"`
	UpdatedAt *time.Time
}

func init() {

}

func (m *Price) BeforeCreate() {

}

func (m Price) TableName() string {
	return "tb_price"
}
