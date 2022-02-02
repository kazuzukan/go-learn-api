package models

import "time"

type Transaction struct {
	Id         int `gorm:"not null;uniqueIndex;primary_key;"`
	CampaignId int `gorm:"index"`
	UserId     int `gorm:"index"`
	Amount     int
	Status     string `gorm:"size:255"`
	Code       string `gorm:"size:255"`
	PaymentUrl int    `gorm:"size:255"`
	CreatedAt  time.Time
	Update     time.Time
}
