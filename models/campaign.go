package models

import "time"

type Campaign struct {
	Id               int    `gorm:"not null;uniqueIndex;primary_key;"`
	UserId           int    `gorm:"index"`
	Name             string `gorm:"size:255;"`
	ShortDescription string `gorm:"size:255;"`
	GoalAmount       int
	CurrentAmount    int
	Description      string `gorm:"size:255;"`
	Slug             string `gorm:"size:255;"`
	Perks            string `gorm:"size:255;"`
	BackerCount      int
	CreatedAt        time.Time
	UpdatedAt        time.Time
	CampaignImage    []CampaignImage
	Transaction      []Transaction
}
