package campaign

import "time"

type Campaign struct {
	ID             int
	UserId         int
	Name           string
	ShortDesc      string
	Description    string
	Perks          string
	BackerCount    int
	GoalAmount     int
	CurrentAmount  int
	Slug           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	CampaignImages []CampaignImage
}

type CampaignImage struct {
	Id         int
	CampaignId int
	FileName   string
	IsPrimary  int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
