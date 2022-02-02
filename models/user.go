package models

import "time"

type User struct {
	Id             int    `gorm:"not null;uniqueIndex;primary_key"`
	Name           string `gorm:"size:255;not null;"`
	Occupation     string `gorm:"size:255;not null;"`
	Email          string `gorm:"size:255;not null;"`
	PasswordHash   string `gorm:"size:255;not null;"`
	AvatarFilename string `gorm:"size:255;not null;"`
	Role           string `gorm:"size:255;not null;"`
	Token          string `gorm:"size:255;not null;"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Campaign       []Campaign
	Transaction    []Transaction
}
