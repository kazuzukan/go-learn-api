package models

import "time"

type CampaignImage struct {
	Id         int `gorm:"not null;uniqueIndex;primary_key;"`
	CampaignId int `gorm:"index"`
	FileName   int `gorm:"size:255;not null"`
	IsPrimary  int `gorm:"size:4;not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
